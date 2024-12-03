package repository

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"hotel-management/internal/domain"
	"hotel-management/internal/gen/hotel_management/public/model"
	"hotel-management/internal/gen/hotel_management/public/table"
)

var (
	ClientNotFound  = errors.New("клиент не найден")
	ClientsNotFound = errors.New("клиенты не найдены")
)

type ClientRepository struct {
	conn *pgx.Conn
}

func NewClientRepository(conn *pgx.Conn) *ClientRepository {
	return &ClientRepository{conn: conn}
}

func (r *ClientRepository) AddClient(ctx context.Context, client domain.Client) error {
	modelClient := model.Clients{
		Name:     client.Name,
		Surname:  client.Surname,
		Passport: client.Passport,
	}

	stmt, args := table.Clients.
		INSERT(table.Clients.AllColumns.Except(table.Clients.ID)).
		MODEL(modelClient).Sql()

	_, err := r.conn.Exec(ctx, stmt, args...)
	return err
}
