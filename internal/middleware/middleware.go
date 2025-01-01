package middleware

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	tele "gopkg.in/telebot.v4"
	"hotel-management/internal/domain"
	"hotel-management/internal/domain/usecase"
	"hotel-management/internal/repository"
	"log"
)

type EmployeeRepository interface {
	IsEmployeeWithPositions(ctx context.Context, username string, positions []domain.Position) (bool, error)
}

type Middleware struct {
	employeeRepository EmployeeRepository
}

func NewMiddleware(conn *pgx.Conn) *Middleware {
	return &Middleware{employeeRepository: repository.NewEmployeeRepository(conn)}
}

func (m *Middleware) Logger() tele.MiddlewareFunc {
	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(c tele.Context) error {
			log.Println(fmt.Sprintf("Получено сообщение от @%s: '%s'", c.Sender().Username, c.Text()))
			return next(c)
		}
	}
}

func (m *Middleware) PermissionCheck(ctx context.Context) tele.MiddlewareFunc {
	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(c tele.Context) error {
			ok, err := m.employeeRepository.IsEmployeeWithPositions(ctx, c.Sender().Username, domain.GetAllPositions())
			if err != nil {
				return c.Send(usecase.ErrorMessage(err))
			}

			if !ok {
				return c.Send(domain.UnknownUserMsg)
			}
			return next(c)
		}
	}
}
