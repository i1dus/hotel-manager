package remove_employee_usecase

import (
	"context"
	"github.com/jackc/pgx/v5"
	tele "gopkg.in/telebot.v4"
	"hotel-management/internal/repository"
	"hotel-management/internal/usecase"
	"strings"
)

type EmployeeRepository interface {
	RemoveEmployee(ctx context.Context, username string) error
}

type RemoveEmployeeUseCase struct {
	employeeRepo EmployeeRepository
}

func NewRemoveEmployeeUseCase(conn *pgx.Conn) *RemoveEmployeeUseCase {
	employeeRepo := repository.NewEmployeeRepository(conn)
	return &RemoveEmployeeUseCase{employeeRepo: employeeRepo}
}

func (uc *RemoveEmployeeUseCase) RemoveEmployee(c tele.Context) error {
	args := c.Args()
	if len(args) != 1 {
		return c.Send("Должно быть 1 аргумента: TG-логин")
	}

	// ТГ-логин
	username := args[0]
	if len(username) > 1 && !strings.HasPrefix(username, "@") {
		return c.Send("Логин должен начинаться с '@'")
	}

	err := uc.employeeRepo.RemoveEmployee(context.Background(), username[1:])
	if err != nil {
		return c.Send(usecase.ErrorMessage(err))
	}
	return c.Send("Сотрудник успешно удален!")
}
