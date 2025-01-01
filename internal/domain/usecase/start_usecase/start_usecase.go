package start_usecase

import (
	"context"
	tele "gopkg.in/telebot.v4"
	"hotel-management/internal/domain"
	"hotel-management/internal/domain/usecase"
)

type EmployeeRepository interface {
	UpsertEmployeeUserID(ctx context.Context, username string, userID int) error
}

type StartUseCase struct {
	employeeRepo EmployeeRepository
	menu         *tele.ReplyMarkup
}

func NewStartUseCase(employeeRepo EmployeeRepository, menu *tele.ReplyMarkup) *StartUseCase {
	return &StartUseCase{employeeRepo: employeeRepo, menu: menu}
}

func (uc *StartUseCase) Start(c tele.Context) error {
	ctx := context.Background()
	employee := c.Sender()

	err := uc.employeeRepo.UpsertEmployeeUserID(ctx, employee.Username, int(employee.ID))
	if err != nil {
		return c.Send(usecase.ErrorMessage(err))
	}

	return c.Send(domain.WelcomeMsg, uc.menu)
}
