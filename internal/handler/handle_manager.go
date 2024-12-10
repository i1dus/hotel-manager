package handler

import (
	"github.com/jackc/pgx/v5"
	tele "gopkg.in/telebot.v4"
	"hotel-management/internal/domain/usecase/add_client_usecase"
	"hotel-management/internal/domain/usecase/add_employee_usecase"
	"hotel-management/internal/domain/usecase/add_room_occupancy_usecase"
	"hotel-management/internal/domain/usecase/add_room_usecase"
	"hotel-management/internal/domain/usecase/change_room_price_usecase"
	"hotel-management/internal/domain/usecase/clean_room_usecase"
	"hotel-management/internal/domain/usecase/end_room_occupancy_usecase"
	"hotel-management/internal/domain/usecase/help_usecase"
	"hotel-management/internal/domain/usecase/list_employees_usecase"
	"hotel-management/internal/domain/usecase/list_room_occupancies_usecase"
	"hotel-management/internal/domain/usecase/list_rooms_usecase"
	"hotel-management/internal/domain/usecase/remove_employee_usecase"
	"hotel-management/internal/domain/usecase/room_cleaned_usecase"
	"hotel-management/internal/domain/usecase/statatistics_usecase"
	"hotel-management/internal/repository"
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

type AddRoomOccupancyUseCase interface {
	AddRoomOccupancy(c tele.Context) error
}

type ListRoomOccupancyUseCase interface {
	ListRoomOccupancy(c tele.Context) error
}

type EndRoomOccupancyUseCase interface {
	EndRoomOccupancy(c tele.Context) error
}

type CleanRoomUseCase interface {
	CleanRoom(c tele.Context) error
}

type RoomCleanedUseCase interface {
	RoomCleaned(c tele.Context) error
}

type StatisticsUseCase interface {
	Statistics(c tele.Context) error
}

type HelpUseCase interface {
	Help(c tele.Context) error
}

type HandlerController struct {
	bot                      *tele.Bot
	addEmployeeUseCase       AddEmployeeUseCase
	removeEmployeeUseCase    RemoveEmployeeUseCase
	listEmployeesUseCase     ListEmployeesUseCase
	addClientUseCase         AddClientUseCase
	addRoomUseCase           AddRoomUseCase
	listRoomsUseCase         ListRoomsUseCase
	changeRoomPriceUseCase   ChangeRoomPriceUseCase
	addRoomOccupancyUseCase  AddRoomOccupancyUseCase
	listRoomOccupancyUseCase ListRoomOccupancyUseCase
	endRoomOccupancyUseCase  EndRoomOccupancyUseCase
	cleanRoomUseCase         CleanRoomUseCase
	roomCleanedUseCase       RoomCleanedUseCase
	statisticsUseCase        StatisticsUseCase
	helpUseCase              HelpUseCase
}

func NewHandlerController(bot *tele.Bot, conn *pgx.Conn) *HandlerController {
	clientRepository := repository.NewClientRepository(conn)
	employeeRepository := repository.NewEmployeeRepository(conn)
	roomRepository := repository.NewRoomRepository(conn)
	roomOccupancyRepository := repository.NewRoomOccupancyRepository(conn)

	return &HandlerController{
		bot:                      bot,
		helpUseCase:              help_usecase.NewHelpUseCase(),
		addEmployeeUseCase:       add_employee_usecase.NewAddEmployeeUseCase(employeeRepository),
		removeEmployeeUseCase:    remove_employee_usecase.NewRemoveEmployeeUseCase(employeeRepository),
		listEmployeesUseCase:     list_employee_usecase.NewListEmployeesUseCase(employeeRepository),
		addClientUseCase:         add_client_usecase.NewAddClientUseCase(clientRepository),
		addRoomUseCase:           add_room_usecase.NewAddRoomUseCase(roomRepository),
		listRoomsUseCase:         list_rooms_usecase.NewListRoomsUseCase(roomRepository),
		changeRoomPriceUseCase:   change_room_price_usecase.NewChangeRoomPriceUseCase(roomRepository),
		addRoomOccupancyUseCase:  add_room_occupancy_usecase.NewAddRoomOccupancyUseCase(roomOccupancyRepository, roomRepository, clientRepository),
		listRoomOccupancyUseCase: list_room_occupancies_usecase.NewListRoomOccupancyUseCase(roomOccupancyRepository),
		endRoomOccupancyUseCase:  end_room_occupancy_usecase.NewEndRoomOccupancyUseCase(roomOccupancyRepository),
		cleanRoomUseCase:         clean_room_usecase.NewCleanRoomUseCase(roomRepository),
		roomCleanedUseCase:       room_cleaned_usecase.NewRoomCleanedUseCase(roomRepository),
		statisticsUseCase:        statatistics_usecase.NewStatisticsUseCase(roomOccupancyRepository),
	}
}

func (ctrl *HandlerController) RegisterHandlers() {
	ctrl.bot.Handle(tele.OnText, func(c tele.Context) error {
		return c.Send("🚀 Я работаю!")
	})

	// employees managing commands
	ctrl.bot.Handle("/add_employee", ctrl.addEmployeeUseCase.AddEmployee)
	ctrl.bot.Handle("/remove_employee", ctrl.removeEmployeeUseCase.RemoveEmployee)
	ctrl.bot.Handle("/employees", ctrl.listEmployeesUseCase.ListEmployees)

	// clients managing commands
	ctrl.bot.Handle("/add_client", ctrl.addClientUseCase.AddClient)

	// rooms managing commands
	ctrl.bot.Handle("/add_room", ctrl.addRoomUseCase.AddRoom)
	ctrl.bot.Handle("/rooms", ctrl.listRoomsUseCase.ListRooms)
	ctrl.bot.Handle("/change_room_price", ctrl.changeRoomPriceUseCase.ChangeRoomPrice)
	ctrl.bot.Handle("/clean_room", ctrl.cleanRoomUseCase.CleanRoom)
	ctrl.bot.Handle("/room_cleaned", ctrl.roomCleanedUseCase.RoomCleaned)

	// room occupancy managing commands
	ctrl.bot.Handle("/add_occupancy", ctrl.addRoomOccupancyUseCase.AddRoomOccupancy)
	ctrl.bot.Handle("/occupancies", ctrl.listRoomOccupancyUseCase.ListRoomOccupancy)
	ctrl.bot.Handle("/end_occupancy", ctrl.endRoomOccupancyUseCase.EndRoomOccupancy)

	// other commands
	ctrl.bot.Handle("/help", ctrl.helpUseCase.Help)
	ctrl.bot.Handle("/stats", ctrl.statisticsUseCase.Statistics)
}
