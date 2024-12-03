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
)

var (
	ErrRoomsNotFound = errors.New("номера не найдены")
	ErrRoomNotFound  = errors.New("номер не найден")
)

type RoomRepository struct {
	conn *pgx.Conn
}

func NewRoomRepository(conn *pgx.Conn) *RoomRepository {
	return &RoomRepository{conn: conn}
}

func (r *RoomRepository) AddRoom(ctx context.Context, room domain.Room) error {
	modelRoom := model.Rooms{
		Number: room.Number,
		Type:   int32(room.Type),
		Price:  int32(room.Price),
	}

	stmt, args := table.Rooms.
		INSERT(table.Rooms.AllColumns.Except(table.Rooms.ID)).
		MODEL(modelRoom).Sql()

	_, err := r.conn.Exec(ctx, stmt, args...)
	return err
}

func (r *RoomRepository) ListRooms(ctx context.Context) ([]domain.Room, error) {
	stmt, args := postgres.SELECT(
		table.Rooms.AllColumns).
		FROM(table.Rooms).Sql()

	var modelRooms []model.Rooms

	rows, err := r.conn.Query(ctx, stmt, args...)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrRoomsNotFound
		}
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var room model.Rooms
		if err := rows.Scan(&room.ID, &room.Number, &room.Type, &room.Price); err != nil {
			return nil, err
		}
		modelRooms = append(modelRooms, room)
	}

	rooms := lo.Map(modelRooms, func(modelRoom model.Rooms, index int) domain.Room {
		return domain.Room{
			Number: modelRoom.Number,
			Type:   domain.RoomCategory(modelRoom.Type),
			Price:  int(modelRoom.Price),
		}
	})

	return rooms, nil
}

func (r *RoomRepository) ChangeRoomPrice(ctx context.Context, number string, price int) error {
	stmt, args := table.Rooms.
		UPDATE(table.Rooms.Price).
		SET(postgres.Int(int64(price))).
		WHERE(table.Rooms.Number.EQ(postgres.String(number))).Sql()

	res, err := r.conn.Exec(ctx, stmt, args...)
	if err != nil {
		return err
	}
	if res.RowsAffected() == 0 {
		return ErrRoomNotFound
	}
	return nil
}
