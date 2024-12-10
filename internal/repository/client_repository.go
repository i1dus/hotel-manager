package repository

import (
	"context"
	"github.com/go-jet/jet/v2/postgres"
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

func (r *ClientRepository) IsClientExist(ctx context.Context, passport string) (bool, error) {
	stmt, args := postgres.
		SELECT(postgres.COUNT(postgres.STAR)).
		FROM(table.Clients).
		WHERE(table.Clients.Passport.EQ(postgres.String(passport))).Sql()

	var count int64
	err := r.conn.QueryRow(ctx, stmt, args...).Scan(&count)
	if err != nil {
		return false, err
	}
	if count < 1 {
		return false, nil
	}
	return true, nil
}
