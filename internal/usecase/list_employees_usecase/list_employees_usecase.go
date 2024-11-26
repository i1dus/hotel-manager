package list_employee_usecase

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

type EmployeeRepository interface {
	ListEmployees(ctx context.Context) ([]domain.Employee, error)
}

type ListEmployeesUseCase struct {
	employeeRepo EmployeeRepository
}

func NewListEmployeesUseCase(conn *pgx.Conn) *ListEmployeesUseCase {
	employeeRepo := repository.NewEmployeeRepository(conn)
	return &ListEmployeesUseCase{employeeRepo: employeeRepo}
}

func (uc *ListEmployeesUseCase) ListEmployees(c tele.Context) error {
	employees, err := uc.employeeRepo.ListEmployees(context.Background())
	if err != nil {
		return c.Send(usecase.ErrorMessage(err))
	}

	message := strings.Builder{}
	message.WriteString("Сотрудники:")
	for _, employee := range employees {
		message.WriteString(fmt.Sprintf("\nUsername: @%s, Должность: '%s', Имя: '%s'",
			employee.Username, employee.Position.GetPositionName(), employee.Name))
	}
	return c.Send(message.String())
}
