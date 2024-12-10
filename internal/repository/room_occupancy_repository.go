package repository

import (
	"context"
	"github.com/go-jet/jet/v2/postgres"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"hotel-management/internal/domain"
	"hotel-management/internal/gen/hotel_management/public/model"
	"hotel-management/internal/gen/hotel_management/public/table"
	"time"
)

var (
	ErrRoomOccupanciesNotFound = errors.New("занятость номеров не найдены")
	ErrRoomOccupancyNotFound   = errors.New("занятость номера не найдена")
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
		Passport:    occupancy.Passport,
		StartAt:     occupancy.StartAt,
		EndAt:       occupancy.EndAt,
		Description: occupancy.Description,
	}

	stmt, args := table.RoomOccupancies.
		INSERT(table.RoomOccupancies.AllColumns.Except(table.RoomOccupancies.ID)).
		MODEL(modelOccupancy).Sql()

	_, err := r.conn.Exec(ctx, stmt, args...)
	return err
}

func (r *RoomOccupancyRepository) ListRoomOccupancy(ctx context.Context) ([]domain.RoomOccupancy, error) {
	stmt, args := postgres.SELECT(
		table.RoomOccupancies.AllColumns).
		FROM(table.RoomOccupancies).Sql()

	var modelRoomOccupancies []model.RoomOccupancies

	rows, err := r.conn.Query(ctx, stmt, args...)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrRoomOccupanciesNotFound
		}
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var roomOccupancy model.RoomOccupancies
		err := rows.Scan(&roomOccupancy.ID, &roomOccupancy.RoomNumber, &roomOccupancy.Passport, &roomOccupancy.StartAt, &roomOccupancy.EndAt, &roomOccupancy.Description)
		if err != nil {
			return nil, err
		}
		modelRoomOccupancies = append(modelRoomOccupancies, roomOccupancy)
	}

	rooms := lo.Map(modelRoomOccupancies, func(modelRoomOccupancy model.RoomOccupancies, index int) domain.RoomOccupancy {
		return domain.RoomOccupancy{
			ID:          int(modelRoomOccupancy.ID),
			RoomNumber:  modelRoomOccupancy.RoomNumber,
			Passport:    modelRoomOccupancy.Passport,
			StartAt:     modelRoomOccupancy.StartAt,
			EndAt:       modelRoomOccupancy.EndAt,
			Description: modelRoomOccupancy.Description,
		}
	})

	return rooms, nil
}

func (r *RoomOccupancyRepository) IsRoomOccupancyPossible(ctx context.Context, occupancy domain.RoomOccupancy) (bool, error) {
	stmt, args := postgres.SELECT(
		postgres.COUNT(postgres.STAR)).
		FROM(table.RoomOccupancies).
		WHERE(table.RoomOccupancies.RoomNumber.EQ(postgres.String(occupancy.RoomNumber)).
			AND(table.RoomOccupancies.StartAt.LT(postgres.TimestampzT(*occupancy.EndAt)).
				AND(table.RoomOccupancies.EndAt.GT(postgres.TimestampzT(occupancy.StartAt)))),
		).
		Sql()

	var count int64
	err := r.conn.QueryRow(ctx, stmt, args...).Scan(&count)
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

func (r *RoomOccupancyRepository) ChangeRoomOccupancyEndAt(ctx context.Context, occupancyID int64, endAt time.Time) error {
	stmt, args := table.RoomOccupancies.
		UPDATE(table.RoomOccupancies.EndAt).
		SET(postgres.TimestampzT(endAt)).
		WHERE(table.RoomOccupancies.ID.EQ(postgres.Int(occupancyID))).Sql()

	res, err := r.conn.Exec(ctx, stmt, args...)
	if err != nil {
		return err
	}
	if res.RowsAffected() == 0 {
		return ErrRoomOccupancyNotFound
	}
	return nil
}

func (r *RoomOccupancyRepository) ListOccupiedRooms(ctx context.Context, now time.Time) ([]domain.RoomOccupancy, error) {
	stmt, args := postgres.SELECT(
		table.RoomOccupancies.AllColumns).
		FROM(table.RoomOccupancies).
		WHERE(table.RoomOccupancies.StartAt.LT(postgres.TimestampzT(now)).
			AND(table.RoomOccupancies.EndAt.GT(postgres.TimestampzT(now))),
		).Sql()

	var modelRoomOccupancies []model.RoomOccupancies

	rows, err := r.conn.Query(ctx, stmt, args...)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrRoomOccupanciesNotFound
		}
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var roomOccupancy model.RoomOccupancies
		err := rows.Scan(&roomOccupancy.ID, &roomOccupancy.RoomNumber, &roomOccupancy.Passport, &roomOccupancy.StartAt, &roomOccupancy.EndAt, &roomOccupancy.Description)
		if err != nil {
			return nil, err
		}
		modelRoomOccupancies = append(modelRoomOccupancies, roomOccupancy)
	}

	rooms := lo.Map(modelRoomOccupancies, func(modelRoomOccupancy model.RoomOccupancies, index int) domain.RoomOccupancy {
		return domain.RoomOccupancy{
			ID:          int(modelRoomOccupancy.ID),
			RoomNumber:  modelRoomOccupancy.RoomNumber,
			Passport:    modelRoomOccupancy.Passport,
			StartAt:     modelRoomOccupancy.StartAt,
			EndAt:       modelRoomOccupancy.EndAt,
			Description: modelRoomOccupancy.Description,
		}
	})

	return rooms, nil
}
