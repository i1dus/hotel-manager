package add_room_occupancy_usecase

import (
	"context"
	"errors"
	tele "gopkg.in/telebot.v4"
	"hotel-management/internal/domain"
	"hotel-management/internal/domain/usecase"
	"strconv"
	"strings"
	"time"
)

const (
	HourEndAt = 12
)

type RoomOccupancyRepository interface {
	AddRoomOccupancy(ctx context.Context, occupancy domain.RoomOccupancy) error
	IsRoomOccupancyPossible(ctx context.Context, occupancy domain.RoomOccupancy) (bool, error)
}

type RoomRepository interface {
	IsRoomExist(ctx context.Context, number string) (bool, error)
}

type ClientRepository interface {
	IsClientExist(ctx context.Context, passport string) (bool, error)
}

type AddRoomOccupancyUseCase struct {
	roomOccupancyRepo RoomOccupancyRepository
	roomRepo          RoomRepository
	clientRepo        ClientRepository
}

func NewAddRoomOccupancyUseCase(roomOccupancyRepo RoomOccupancyRepository, roomRepo RoomRepository, clientRepo ClientRepository) *AddRoomOccupancyUseCase {
	return &AddRoomOccupancyUseCase{roomOccupancyRepo: roomOccupancyRepo, roomRepo: roomRepo, clientRepo: clientRepo}
}

func (uc *AddRoomOccupancyUseCase) AddRoomOccupancy(c tele.Context) error {
	ctx := context.Background()
	args := c.Args()
	if len(args) < 3 {
		return c.Send("Должно быть минимум 3 аргумента: Номер комнаты, Паспорт клиента, Конец занятости (DD-MM-YYYY), опционально: описание")
	}

	// Номер комнаты
	roomNumber := args[0]
	exist, err := uc.roomRepo.IsRoomExist(ctx, roomNumber)
	if err != nil {
		return c.Send(usecase.ErrorMessage(err))
	}
	if !exist {
		return c.Send("Комната не найдена")
	}

	// ID-клиента
	passport := args[1]
	exist, err = uc.clientRepo.IsClientExist(ctx, passport)
	if err != nil {
		return c.Send(usecase.ErrorMessage(err))
	}
	if !exist {
		return c.Send("Клиент не найден")
	}

	// Конец занятости
	occupancyEndAtStr := args[2]
	occupancyEndAt, err := parseDate(occupancyEndAtStr)
	if err != nil {
		return c.Send(usecase.ErrorMessage(err))
	}

	// Описание
	description := strings.Join(args[3:], " ")

	// Сохранение
	occupancy := domain.RoomOccupancy{
		RoomNumber:  roomNumber,
		Passport:    passport,
		StartAt:     time.Now(),
		EndAt:       occupancyEndAt,
		Description: description,
	}

	// Возможно ли добавить занятость (нет ли пересечения с другими занятостями)
	ok, err := uc.roomOccupancyRepo.IsRoomOccupancyPossible(ctx, occupancy)
	if err != nil {
		return c.Send(usecase.ErrorMessage(err))
	}
	if !ok {
		return c.Send("Занятость невозможно добавить ввиду пересечания с другой занятостью")
	}

	err = uc.roomOccupancyRepo.AddRoomOccupancy(context.Background(), occupancy)
	if err != nil {
		return c.Send(usecase.ErrorMessage(err))
	}
	return c.Send(domain.PrefixSuccess + "Занятость номера успешно добавлена!")
}

var ErrDateFormat = errors.New("неверный формат даты (должен быть: \"DD-MM-YYYY\")")

func parseDate(dateStr string) (*time.Time, error) {
	if len(dateStr) != 10 { // "DD-MM-YYYY"
		return nil, ErrDateFormat
	}

	dayStr, monthStr, yearStr := dateStr[:2], dateStr[3:5], dateStr[6:]

	day, err := strconv.Atoi(dayStr)
	if err != nil {
		return nil, err
	}
	month, err := strconv.Atoi(monthStr)
	if err != nil {
		return nil, err
	}
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return nil, err
	}
	mskLocation, _ := time.LoadLocation("Europe/Moscow")
	date := time.Date(year, time.Month(month), day, HourEndAt, 0, 0, 0, mskLocation)
	return &date, nil
}
