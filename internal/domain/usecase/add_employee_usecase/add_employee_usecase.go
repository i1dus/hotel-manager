package add_employee_usecase

import (
	"context"
	tele "gopkg.in/telebot.v4"
	"hotel-management/internal/domain"
	"hotel-management/internal/domain/usecase"
	"strings"
)

type EmployeeRepository interface {
	AddEmployee(ctx context.Context, employee domain.Employee) error
}

type AddEmployeeUseCase struct {
	employeeRepo EmployeeRepository
}

func NewAddEmployeeUseCase(employeeRepo EmployeeRepository) *AddEmployeeUseCase {
	return &AddEmployeeUseCase{employeeRepo: employeeRepo}
}

func (uc *AddEmployeeUseCase) AddEmployee(c tele.Context) error {
	args := c.Args()
	if len(args) != 3 {
		return c.Send("Должно быть 3 аргумента: TG-логин, Позиция, Имя")
	}

	// ТГ-логин
	username := args[0]
	if len(username) > 1 && !strings.HasPrefix(username, "@") {
		return c.Send("Логин должен начинаться с '@'")
	}

	// Должность
	var position domain.Position
	positionName := strings.ToLower(args[1])
	switch positionName {
	case string(domain.PositionNameManager):
		position = domain.PositionManager
	case string(domain.PositionNameReceptionist):
		position = domain.PositionReceptionist
	case string(domain.PositionNameMaid):
		position = domain.PositionMaid
	default:
		return c.Send("Неизвестная должность")
	}

	// Имя
	name := args[2]

	// Сохранение
	employee := domain.Employee{
		Username: username[1:],
		Name:     name,
		Position: position,
	}

	err := uc.employeeRepo.AddEmployee(context.Background(), employee)
	if err != nil {
		return c.Send(usecase.ErrorMessage(err))
	}
	return c.Send(domain.PrefixSuccess + "Сотрудник успешно добавлен!")
}
