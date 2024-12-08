package list_employee_usecase

import (
	"context"
	"fmt"
	tele "gopkg.in/telebot.v4"
	"hotel-management/internal/domain"
	"hotel-management/internal/domain/usecase"
	"strings"
)

type EmployeeRepository interface {
	ListEmployees(ctx context.Context) ([]domain.Employee, error)
}

type ListEmployeesUseCase struct {
	employeeRepo EmployeeRepository
}

func NewListEmployeesUseCase(employeeRepo EmployeeRepository) *ListEmployeesUseCase {
	return &ListEmployeesUseCase{employeeRepo: employeeRepo}
}

func (uc *ListEmployeesUseCase) ListEmployees(c tele.Context) error {
	employees, err := uc.employeeRepo.ListEmployees(context.Background())
	if err != nil {
		return c.Send(usecase.ErrorMessage(err))
	}

	message := strings.Builder{}
	message.WriteString("Сотрудники:")

	if len(employees) == 0 {
		message.WriteString("\nСотрудники не найдены")
		return c.Send(message.String())
	}

	for _, employee := range employees {
		message.WriteString(fmt.Sprintf("\nUsername: @%s, Должность: '%s', Имя: '%s'",
			employee.Username, employee.Position.GetPositionName(), employee.Name))
	}
	return c.Send(message.String())
}
