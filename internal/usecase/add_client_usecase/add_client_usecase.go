package add_client_usecase

import (
	"context"
	"github.com/jackc/pgx/v5"
	tele "gopkg.in/telebot.v4"
	"hotel-management/internal/domain"
	"hotel-management/internal/repository"
	"hotel-management/internal/usecase"
)

type ClientRepository interface {
	AddClient(ctx context.Context, client domain.Client) error
}

type AddClientUseCase struct {
	clientRepo ClientRepository
}

func NewAddClientUseCase(conn *pgx.Conn) *AddClientUseCase {
	clientRepo := repository.NewClientRepository(conn)
	return &AddClientUseCase{clientRepo: clientRepo}
}

func (uc *AddClientUseCase) AddClient(c tele.Context) error {
	args := c.Args()
	if len(args) != 3 {
		return c.Send("Должно быть 3 аргумента: Имя, Фамилия, Номер паспорта")
	}

	// Имя
	name := args[0]

	// Фамилия
	surname := args[1]

	// Паспорт
	passport := args[2]

	// Сохранение
	client := domain.Client{
		Name:     name,
		Surname:  surname,
		Passport: passport,
	}

	err := uc.clientRepo.AddClient(context.Background(), client)
	if err != nil {
		return c.Send(usecase.ErrorMessage(err))
	}
	return c.Send("Клиент успешно добавлен!")
}
