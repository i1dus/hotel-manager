package add_room_occupancy_usecase

import (
	"context"
	"github.com/jackc/pgx/v5"
	tele "gopkg.in/telebot.v4"
	"hotel-management/internal/domain"
	"hotel-management/internal/repository"
	"hotel-management/internal/usecase"
	"strconv"
	"time"
)

type RoomOccupancyRepository interface {
	AddRoomOccupancy(ctx context.Context, occupancy domain.RoomOccupancy) error
}

type AddRoomOccupancyUseCase struct {
	roomOccupancyRepo RoomOccupancyRepository
}

func NewAddRoomOccupancyUseCase(conn *pgx.Conn) *AddRoomOccupancyUseCase {
	roomOccupancyRepo := repository.NewRoomOccupancyRepository(conn)
	return &AddRoomOccupancyUseCase{roomOccupancyRepo: roomOccupancyRepo}
}

func (uc *AddRoomOccupancyUseCase) AddRoomOccupancy(c tele.Context) error {
	args := c.Args()
	if len(args) != 3 {
		return c.Send("Должно быть 3 аргумента: Номер комнаты, ID клиента, Конец занятости (DD-MM-YYYY)")
	}

	// ТГ-логин
	roomNumber := args[0]
	// todo: добавить проверку на существование

	// ID-клиента
	clientIDStr := args[1]
	clientID, err := strconv.Atoi(clientIDStr)
	if err != nil {
		return c.Send(usecase.ErrorMessage(err))
	}
	// todo: добавить проверку на существование

	// Конец занятости
	//occupancyEndAt := args[2]
	// todo: распарсить

	// Сохранение
	occupancy := domain.RoomOccupancy{
		RoomNumber:  roomNumber,
		ClientID:    clientID,
		StartAt:     time.Now(),
		EndAt:       time.Time{},
		Description: "", // todo: добавить описание
	}

	err = uc.roomOccupancyRepo.AddRoomOccupancy(context.Background(), occupancy)
	if err != nil {
		return c.Send(usecase.ErrorMessage(err))
	}
	return c.Send("Занятость номера успешно добавлена!")
}
