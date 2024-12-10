package add_room_usecase

import (
	"context"
	tele "gopkg.in/telebot.v4"
	"hotel-management/internal/domain"
	"hotel-management/internal/domain/usecase"
	"strconv"
	"strings"
)

type RoomRepository interface {
	AddRoom(ctx context.Context, room domain.Room) error
}

type AddRoomUseCase struct {
	roomRepo RoomRepository
}

func NewAddRoomUseCase(roomRepo RoomRepository) *AddRoomUseCase {
	return &AddRoomUseCase{roomRepo: roomRepo}
}

func (uc *AddRoomUseCase) AddRoom(c tele.Context) error {
	args := c.Args()
	if len(args) != 3 {
		return c.Send("Должно быть 3 аргумента: Номер, Категория, Цена за сутки")
	}

	// Номер
	number := args[0]

	// Категория
	var roomType domain.RoomCategory
	roomTypeName := strings.ToLower(args[1])
	switch roomTypeName {
	case string(domain.RoomCategoryNameStandard):
		roomType = domain.RoomCategoryStandard
	case string(domain.RoomCategoryNameComfort):
		roomType = domain.RoomCategoryComfort
	case string(domain.RoomTypeNameLuxe):
		roomType = domain.RoomCategoryLuxe

	default:
		return c.Send("Неизвестная категория номера")
	}

	// Цена
	priceStr := args[2]
	price, err := strconv.Atoi(priceStr)
	if err != nil {
		return c.Send(usecase.ErrorMessage(err))
	}

	// Сохранение
	room := domain.Room{
		Number: number,
		Type:   roomType,
		Price:  price,
	}

	err = uc.roomRepo.AddRoom(context.Background(), room)
	if err != nil {
		return c.Send(usecase.ErrorMessage(err))
	}
	return c.Send(domain.PrefixSuccess + "Номер успешно добавлен!")
}
