package repository

import (
	"context"
	"github.com/jackc/pgx/v5"
	"hotel-management/internal/domain"
	"hotel-management/internal/gen/hotel_management/public/model"
	"hotel-management/internal/gen/hotel_management/public/table"
)

type RoomOccupancyRepository struct {
	conn *pgx.Conn
}

func NewRoomOccupancyRepository(conn *pgx.Conn) *RoomOccupancyRepository {
	return &RoomOccupancyRepository{conn: conn}
}

func (r *RoomOccupancyRepository) AddRoomOccupancy(ctx context.Context, occupancy domain.RoomOccupancy) error {
	modelOccupancy := model.RoomOccupancies{
		RoomNumber:  occupancy.RoomNumber,
		ClientID:    int32(occupancy.ClientID),
		StartAt:     occupancy.StartAt,
		EndAt:       &occupancy.EndAt,
		Description: &occupancy.Description,
	}

	stmt, args := table.RoomOccupancies.
		INSERT(table.RoomOccupancies.AllColumns.Except(table.RoomOccupancies.ID)).
		MODEL(modelOccupancy).Sql()

	_, err := r.conn.Exec(ctx, stmt, args...)
	return err
}
