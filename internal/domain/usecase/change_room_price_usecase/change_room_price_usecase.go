package change_room_price_usecase

import (
	"context"
	"fmt"
	tele "gopkg.in/telebot.v4"
	"hotel-management/internal/domain"
	"hotel-management/internal/domain/usecase"
	"strconv"
)

type RoomRepository interface {
	ChangeRoomPrice(ctx context.Context, number string, price int) error
}

type ChangeRoomPriceUseCase struct {
	roomRepo RoomRepository
}

func NewChangeRoomPriceUseCase(roomRepo RoomRepository) *ChangeRoomPriceUseCase {
	return &ChangeRoomPriceUseCase{roomRepo: roomRepo}
}

func (uc *ChangeRoomPriceUseCase) ChangeRoomPrice(c tele.Context) error {
	args := c.Args()
	if len(args) != 2 {
		return c.Send("Должно быть 2 аргумента: Номер, Новая цена за сутки")
	}

	// Номер
	number := args[0]

	// Цена
	priceStr := args[1]
	price, err := strconv.Atoi(priceStr)
	if err != nil {
		return c.Send(usecase.ErrorMessage(err))
	}

	err = uc.roomRepo.ChangeRoomPrice(context.Background(), number, price)
	if err != nil {
		return c.Send(usecase.ErrorMessage(err))
	}
	return c.Send(domain.PrefixSuccess + fmt.Sprintf("Цена за сутки номера '%s' успешно обновлена на %d₽!", number, price))
}
