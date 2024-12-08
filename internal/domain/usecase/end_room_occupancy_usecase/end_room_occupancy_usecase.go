package end_room_occupancy_usecase

import (
	"context"
	tele "gopkg.in/telebot.v4"
	"hotel-management/internal/domain"
	"hotel-management/internal/domain/usecase"
	"strconv"
	"time"
)

type RoomOccupancyRepository interface {
	ChangeRoomOccupancyEndAt(ctx context.Context, occupancyID int64, endAt time.Time) error
}

type EndRoomOccupancyUseCase struct {
	roomOccupancyRepo RoomOccupancyRepository
}

func NewEndRoomOccupancyUseCase(roomOccupancyRepo RoomOccupancyRepository) *EndRoomOccupancyUseCase {
	return &EndRoomOccupancyUseCase{roomOccupancyRepo: roomOccupancyRepo}
}

func (uc *EndRoomOccupancyUseCase) EndRoomOccupancy(c tele.Context) error {
	ctx := context.Background()
	args := c.Args()
	if len(args) != 1 {
		return c.Send("Должен быть 1 аргумент: ID занятости")
	}

	// ID занятости
	occupancyIDStr := args[0]
	occupancyID, err := strconv.Atoi(occupancyIDStr)
	if err != nil {
		return c.Send(usecase.ErrorMessage(err))
	}

	// Изменение конца занятости
	err = uc.roomOccupancyRepo.ChangeRoomOccupancyEndAt(ctx, int64(occupancyID), time.Now())
	if err != nil {
		return c.Send(usecase.ErrorMessage(err))
	}
	return c.Send(domain.PrefixSuccess + "Занятость успешно снята")
}
