package list_room_occupancies_usecase

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
	ListRoomOccupancy(ctx context.Context) ([]domain.RoomOccupancy, error)
}

type ListRoomOccupancyUseCase struct {
	roomOccupancyRepo RoomOccupancyRepository
}

func NewListRoomOccupancyUseCase(roomOccupancyRepo RoomOccupancyRepository) *ListRoomOccupancyUseCase {
	return &ListRoomOccupancyUseCase{roomOccupancyRepo: roomOccupancyRepo}
}

func (uc *ListRoomOccupancyUseCase) ListRoomOccupancy(c tele.Context) error {
	roomOccupancies, err := uc.roomOccupancyRepo.ListRoomOccupancy(context.Background())
	if err != nil {
		return c.Send(usecase.ErrorMessage(err))
	}

	message := strings.Builder{}
	message.WriteString("Занятость номеров:")

	if len(roomOccupancies) == 0 {
		message.WriteString("\nЗанятости номеров не найдены")
		return c.Send(message.String())
	}

	for _, roomOccupancy := range roomOccupancies {
		message.WriteString(fmt.Sprintf("\nID: %d, Номер: '%s', Паспорт клиента: '%s', Начало занятости: '%s', Конец занятости: '%s', Описание: '%s'",
			roomOccupancy.ID, roomOccupancy.RoomNumber, roomOccupancy.Passport,
			roomOccupancy.StartAt.Format(time.DateTime), roomOccupancy.EndAt.Format(time.DateTime), roomOccupancy.Description))
	}
	return c.Send(message.String())
}
