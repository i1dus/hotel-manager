package add_room_occupancy_usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	tele "gopkg.in/telebot.v4"
	"hotel-management/internal/domain"
	"hotel-management/internal/repository"
	"hotel-management/internal/usecase"
	"strconv"
	"strings"
	"time"
)

type RoomOccupancyRepository interface {
	AddRoomOccupancy(ctx context.Context, occupancy domain.RoomOccupancy) error
}

type RoomRepository interface {
	IsRoomExist(ctx context.Context, number string) (bool, error)
}

type ClientRepository interface {
	IsClientExist(ctx context.Context, clientID int64) (bool, error)
}

type AddRoomOccupancyUseCase struct {
	roomOccupancyRepo RoomOccupancyRepository
	roomRepo          RoomRepository
	clientRepo        ClientRepository
}

func NewAddRoomOccupancyUseCase(conn *pgx.Conn) *AddRoomOccupancyUseCase {
	roomOccupancyRepo := repository.NewRoomOccupancyRepository(conn)
	roomRepo := repository.NewRoomRepository(conn)
	clientRepo := repository.NewClientRepository(conn)
	return &AddRoomOccupancyUseCase{roomOccupancyRepo: roomOccupancyRepo, roomRepo: roomRepo, clientRepo: clientRepo}
}

func (uc *AddRoomOccupancyUseCase) AddRoomOccupancy(c tele.Context) error {
	ctx := context.Background()
	args := c.Args()
	if len(args) < 3 {
		return c.Send("Должно быть минимум 3 аргумента: Номер комнаты, ID клиента, Конец занятости (DD-MM-YYYY), опционально: описание")
	}

	// ТГ-логин
	roomNumber := args[0]
	exist, err := uc.roomRepo.IsRoomExist(ctx, roomNumber)
	if err != nil {
		return c.Send(usecase.ErrorMessage(err))
	}
	if !exist {
		return c.Send("Комната не найдена")
	}

	// ID-клиента
	clientIDStr := args[1]
	clientID, err := strconv.Atoi(clientIDStr)
	if err != nil {
		return c.Send(usecase.ErrorMessage(err))
	}
	exist, err = uc.clientRepo.IsClientExist(ctx, int64(clientID))
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
		ClientID:    clientID,
		StartAt:     time.Now(),
		EndAt:       *occupancyEndAt,
		Description: description,
	}

	err = uc.roomOccupancyRepo.AddRoomOccupancy(context.Background(), occupancy)
	if err != nil {
		return c.Send(usecase.ErrorMessage(err))
	}
	return c.Send("Занятость номера успешно добавлена!")
}

var ErrDateFormat = errors.New("неверный формат даты (должен быть: \"DD-MM-YYYY\")")

func parseDate(dateStr string) (*time.Time, error) {
	if len(dateStr) != 10 { // "DD-MM-YYYY"
		return nil, ErrDateFormat
	}

	dayStr, monthStr, yearStr := dateStr[:2], dateStr[3:5], dateStr[6:]
	fmt.Println(dayStr, monthStr, yearStr)
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
	date := time.Date(year, time.Month(month), day, 12, 0, 0, 0, mskLocation)
	return &date, nil
}
