package room_handler

import tele "gopkg.in/telebot.v4"

type AddRoomUseCase interface {
	AddRoom(c tele.Context) error
}

type ListRoomsUseCase interface {
	ListRooms(c tele.Context) error
}

type ChangeRoomPriceUseCase interface {
	ChangeRoomPrice(c tele.Context) error
}

type CleanRoomUseCase interface {
	CleanRoom(c tele.Context) error
}

type RoomCleanedUseCase interface {
	RoomCleaned(c tele.Context) error
}

// RoomHandler has room managing commands
type RoomHandler struct {
	bot                    *tele.Bot
	addRoomUseCase         AddRoomUseCase
	listRoomsUseCase       ListRoomsUseCase
	changeRoomPriceUseCase ChangeRoomPriceUseCase
	cleanRoomUseCase       CleanRoomUseCase
	roomCleanedUseCase     RoomCleanedUseCase
}

func NewRoomHandler(bot *tele.Bot, addRoomUseCase AddRoomUseCase, listRoomsUseCase ListRoomsUseCase, changeRoomPriceUseCase ChangeRoomPriceUseCase, cleanRoomUseCase CleanRoomUseCase, roomCleanedUseCase RoomCleanedUseCase) *RoomHandler {
	return &RoomHandler{bot: bot, addRoomUseCase: addRoomUseCase, listRoomsUseCase: listRoomsUseCase, changeRoomPriceUseCase: changeRoomPriceUseCase, cleanRoomUseCase: cleanRoomUseCase, roomCleanedUseCase: roomCleanedUseCase}
}

func (h *RoomHandler) RegisterHandlers() {
	h.bot.Handle("/add_room", h.addRoomUseCase.AddRoom)
	h.bot.Handle("/rooms", h.listRoomsUseCase.ListRooms)
	h.bot.Handle("/change_room_price", h.changeRoomPriceUseCase.ChangeRoomPrice)
	h.bot.Handle("/clean_room", h.cleanRoomUseCase.CleanRoom)
	h.bot.Handle("/room_cleaned", h.roomCleanedUseCase.RoomCleaned)
}
