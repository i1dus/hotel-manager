package handlers

import (
	"context"
	"github.com/jackc/pgx/v5"
	tele "gopkg.in/telebot.v4"
	"hotel-management/internal/usecase/add_employee_usecase"
	list_employee_usecase "hotel-management/internal/usecase/list_employees_usecase"
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

type HandlerController struct {
	bot                   *tele.Bot
	addEmployeeUseCase    AddEmployeeUseCase
	removeEmployeeUseCase RemoveEmployeeUseCase
	listEmployeesUseCase  ListEmployeesUseCase
}

func NewHandlerController(bot *tele.Bot, conn *pgx.Conn) *HandlerController {
	addEmployeeUseCase := add_employee_usecase.NewAddEmployeeUseCase(conn)
	removeEmployeeUseCase := remove_employee_usecase.NewRemoveEmployeeUseCase(conn)
	listEmployeesUseCase := list_employee_usecase.NewListEmployeesUseCase(conn)
	return &HandlerController{
		bot:                   bot,
		addEmployeeUseCase:    addEmployeeUseCase,
		removeEmployeeUseCase: removeEmployeeUseCase,
		listEmployeesUseCase:  listEmployeesUseCase,
	}
}

func (ctrl *HandlerController) RegisterHandlers(ctx context.Context) error {
	ctrl.bot.Handle("/hello", func(c tele.Context) error {
		return c.Send("Hello!")
	})

	ctrl.bot.Handle("/add_employee", ctrl.addEmployeeUseCase.AddEmployee)
	ctrl.bot.Handle("/remove_employee", ctrl.removeEmployeeUseCase.RemoveEmployee)
	ctrl.bot.Handle("/employees", ctrl.listEmployeesUseCase.ListEmployees)
	return nil
}
