package add_client_usecase

import (
	"context"
	tele "gopkg.in/telebot.v4"
	"hotel-management/internal/domain"
	"hotel-management/internal/domain/usecase"
)

type ClientRepository interface {
	AddClient(ctx context.Context, client domain.Client) error
}

type AddClientUseCase struct {
	clientRepo ClientRepository
}

func NewAddClientUseCase(clientRepo ClientRepository) *AddClientUseCase {
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
	return c.Send(domain.PrefixSuccess + "Клиент успешно добавлен!")
}
