package list_rooms_usecase

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	tele "gopkg.in/telebot.v4"
	"hotel-management/internal/domain"
	"hotel-management/internal/repository"
	"hotel-management/internal/usecase"
	"strings"
)

type RoomRepository interface {
	ListRooms(ctx context.Context) ([]domain.Room, error)
}

type ListRoomsUseCase struct {
	roomRepo RoomRepository
}

func NewListRoomsUseCase(conn *pgx.Conn) *ListRoomsUseCase {
	roomRepo := repository.NewRoomRepository(conn)
	return &ListRoomsUseCase{roomRepo: roomRepo}
}

func (uc *ListRoomsUseCase) ListRooms(c tele.Context) error {
	rooms, err := uc.roomRepo.ListRooms(context.Background())
	if err != nil {
		return c.Send(usecase.ErrorMessage(err))
	}

	message := strings.Builder{}
	message.WriteString("Номера:")
	for _, room := range rooms {
		message.WriteString(fmt.Sprintf("\nНомер: '%s', Категория: '%s', Цена за сутки: %d₽",
			room.Number, room.Type.GetRoomTypeName(), room.Price))
	}
	return c.Send(message.String())
}
