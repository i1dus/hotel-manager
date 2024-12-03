package handler

import (
	"context"
	"github.com/jackc/pgx/v5"
	tele "gopkg.in/telebot.v4"
	"hotel-management/internal/usecase/add_client_usecase"
	"hotel-management/internal/usecase/add_employee_usecase"
	"hotel-management/internal/usecase/add_room_usecase"
	"hotel-management/internal/usecase/change_room_price_usecase"
	list_employee_usecase "hotel-management/internal/usecase/list_employees_usecase"
	"hotel-management/internal/usecase/list_rooms_usecase"
	"hotel-management/internal/usecase/remove_employee_usecase"
)

type AddEmployeeUseCase interface {
	AddEmployee(c tele.Context) error
}

type RemoveEmployeeUseCase interface {
	RemoveEmployee(c tele.Context) error
}

type ListEmployeesUseCase interface {
	ListEmployees(c tele.Context) error
}

type AddClientUseCase interface {
	AddClient(c tele.Context) error
}

type AddRoomUseCase interface {
	AddRoom(c tele.Context) error
}

type ListRoomsUseCase interface {
	ListRooms(c tele.Context) error
}

type ChangeRoomPriceUseCase interface {
	ChangeRoomPrice(c tele.Context) error
}

type HandlerController struct {
	bot                    *tele.Bot
	addEmployeeUseCase     AddEmployeeUseCase
	removeEmployeeUseCase  RemoveEmployeeUseCase
	listEmployeesUseCase   ListEmployeesUseCase
	addClientUseCase       AddClientUseCase
	addRoomUseCase         AddRoomUseCase
	listRoomsUseCase       ListRoomsUseCase
	changeRoomPriceUseCase ChangeRoomPriceUseCase
}

func NewHandlerController(bot *tele.Bot, conn *pgx.Conn) *HandlerController {
	addEmployeeUseCase := add_employee_usecase.NewAddEmployeeUseCase(conn)
	removeEmployeeUseCase := remove_employee_usecase.NewRemoveEmployeeUseCase(conn)
	listEmployeesUseCase := list_employee_usecase.NewListEmployeesUseCase(conn)
	addClientUseCase := add_client_usecase.NewAddClientUseCase(conn)
	addRoomUseCase := add_room_usecase.NewAddRoomUseCase(conn)
	listRoomsUseCase := list_rooms_usecase.NewListRoomsUseCase(conn)
	changeRoomPriceUseCase := change_room_price_usecase.NewChangeRoomPriceUseCase(conn)
	return &HandlerController{
		bot:                    bot,
		addEmployeeUseCase:     addEmployeeUseCase,
		removeEmployeeUseCase:  removeEmployeeUseCase,
		listEmployeesUseCase:   listEmployeesUseCase,
		addClientUseCase:       addClientUseCase,
		addRoomUseCase:         addRoomUseCase,
		listRoomsUseCase:       listRoomsUseCase,
		changeRoomPriceUseCase: changeRoomPriceUseCase,
	}
}

func (ctrl *HandlerController) RegisterHandlers(ctx context.Context) {
	ctrl.bot.Handle(tele.OnText, func(c tele.Context) error {
		return c.Send("я работаю!")
	})

	// employees managing commands
	ctrl.bot.Handle("/add_employee", ctrl.addEmployeeUseCase.AddEmployee)
	ctrl.bot.Handle("/remove_employee", ctrl.removeEmployeeUseCase.RemoveEmployee)
	ctrl.bot.Handle("/employees", ctrl.listEmployeesUseCase.ListEmployees)

	// clients managing commands
	ctrl.bot.Handle("/add_client", ctrl.addClientUseCase.AddClient)

	// rooms managing commads
	ctrl.bot.Handle("/add_room", ctrl.addRoomUseCase.AddRoom)
	ctrl.bot.Handle("/list_rooms", ctrl.listRoomsUseCase.ListRooms)
	ctrl.bot.Handle("/change_room_price", ctrl.changeRoomPriceUseCase.ChangeRoomPrice)
}
