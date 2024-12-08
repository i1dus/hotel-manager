package statatistics_usecase

import (
	"context"
	"fmt"
	tele "gopkg.in/telebot.v4"
	"hotel-management/internal/domain"
	"hotel-management/internal/domain/usecase"
	"strings"
	"time"
)

type RoomOccupancyRepository interface {
	ListOccupiedRooms(ctx context.Context, now time.Time) ([]domain.RoomOccupancy, error)
}

type StatisticsUseCase struct {
	roomOccupancyRepo RoomOccupancyRepository
}

func NewStatisticsUseCase(roomOccupancyRepo RoomOccupancyRepository) *StatisticsUseCase {
	return &StatisticsUseCase{roomOccupancyRepo: roomOccupancyRepo}
}

func (uc *StatisticsUseCase) Statistics(c tele.Context) error {
	occupiedRooms, err := uc.roomOccupancyRepo.ListOccupiedRooms(context.Background(), time.Now())
	if err != nil {
		return c.Send(usecase.ErrorMessage(err))
	}

	message := strings.Builder{}
	message.WriteString("Занятые номера на текущий момент:")

	if len(occupiedRooms) == 0 {
		message.WriteString("\nЗанятые номера не найдены")
		return c.Send(message.String())
	}

	for _, occupiedRoom := range occupiedRooms {
		message.WriteString(fmt.Sprintf("\nID: %d, Номер: '%s', Паспорт клиента: '%s', Начало занятости: '%s', Конец занятости: '%s', Описание: '%s'",
			occupiedRoom.ID, occupiedRoom.RoomNumber, occupiedRoom.Passport,
			occupiedRoom.StartAt.Format(time.DateTime), occupiedRoom.EndAt.Format(time.DateTime), occupiedRoom.Description))
	}
	return c.Send(message.String())
}
