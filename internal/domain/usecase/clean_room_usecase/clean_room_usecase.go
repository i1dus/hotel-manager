package clean_room_usecase

import (
	"context"
	tele "gopkg.in/telebot.v4"
	"hotel-management/internal/domain"
	"hotel-management/internal/domain/usecase"
)

type RoomRepository interface {
	ChangeRoomCleaned(ctx context.Context, number string, cleaned bool) error
}

type CleanRoomUseCase struct {
	roomRepo RoomRepository
}

func NewCleanRoomUseCase(roomRepo RoomRepository) *CleanRoomUseCase {
	return &CleanRoomUseCase{roomRepo: roomRepo}
}

func (uc *CleanRoomUseCase) CleanRoom(c tele.Context) error {
	ctx := context.Background()
	args := c.Args()
	if len(args) != 1 {
		return c.Send("Должен быть 1 аргумент: Номер комнаты")
	}

	// Номер
	number := args[0]

	err := uc.roomRepo.ChangeRoomCleaned(ctx, number, false)
	if err != nil {
		return c.Send(usecase.ErrorMessage(err))
	}
	return c.Send(domain.PrefixSuccess + "Уборка в комнату выставлена")
}
