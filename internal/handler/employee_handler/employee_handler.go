package employee_handler

import tele "gopkg.in/telebot.v4"

type AddEmployeeUseCase interface {
	AddEmployee(c tele.Context) error
}

type RemoveEmployeeUseCase interface {
	RemoveEmployee(c tele.Context) error
}

type ListEmployeesUseCase interface {
	ListEmployees(c tele.Context) error
}

// EmployeeHandler has employees managing commands
type EmployeeHandler struct {
	bot                   *tele.Bot
	addEmployeeUseCase    AddEmployeeUseCase
	removeEmployeeUseCase RemoveEmployeeUseCase
	listEmployeesUseCase  ListEmployeesUseCase
}

func NewEmployeeHandler(bot *tele.Bot, addEmployeeUseCase AddEmployeeUseCase, removeEmployeeUseCase RemoveEmployeeUseCase, listEmployeesUseCase ListEmployeesUseCase) *EmployeeHandler {
	return &EmployeeHandler{bot: bot, addEmployeeUseCase: addEmployeeUseCase, removeEmployeeUseCase: removeEmployeeUseCase, listEmployeesUseCase: listEmployeesUseCase}
}

func (h *EmployeeHandler) RegisterHandlers() {
	h.bot.Handle("/add_employee", h.addEmployeeUseCase.AddEmployee)
	h.bot.Handle("/remove_employee", h.removeEmployeeUseCase.RemoveEmployee)
	h.bot.Handle("/employees", h.listEmployeesUseCase.ListEmployees)
}
