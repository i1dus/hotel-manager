//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/postgres"
)

var RoomOccupancies = newRoomOccupanciesTable("public", "room_occupancies", "")

type roomOccupanciesTable struct {
	postgres.Table

	// Columns
	ID          postgres.ColumnInteger
	RoomNumber  postgres.ColumnString
	ClientID    postgres.ColumnInteger
	StartAt     postgres.ColumnTimestampz
	EndAt       postgres.ColumnTimestampz
	Description postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type RoomOccupanciesTable struct {
	roomOccupanciesTable

	EXCLUDED roomOccupanciesTable
}

// AS creates new RoomOccupanciesTable with assigned alias
func (a RoomOccupanciesTable) AS(alias string) *RoomOccupanciesTable {
	return newRoomOccupanciesTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new RoomOccupanciesTable with assigned schema name
func (a RoomOccupanciesTable) FromSchema(schemaName string) *RoomOccupanciesTable {
	return newRoomOccupanciesTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new RoomOccupanciesTable with assigned table prefix
func (a RoomOccupanciesTable) WithPrefix(prefix string) *RoomOccupanciesTable {
	return newRoomOccupanciesTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new RoomOccupanciesTable with assigned table suffix
func (a RoomOccupanciesTable) WithSuffix(suffix string) *RoomOccupanciesTable {
	return newRoomOccupanciesTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newRoomOccupanciesTable(schemaName, tableName, alias string) *RoomOccupanciesTable {
	return &RoomOccupanciesTable{
		roomOccupanciesTable: newRoomOccupanciesTableImpl(schemaName, tableName, alias),
		EXCLUDED:             newRoomOccupanciesTableImpl("", "excluded", ""),
	}
}

func newRoomOccupanciesTableImpl(schemaName, tableName, alias string) roomOccupanciesTable {
	var (
		IDColumn          = postgres.IntegerColumn("id")
		RoomNumberColumn  = postgres.StringColumn("room_number")
		ClientIDColumn    = postgres.IntegerColumn("client_id")
		StartAtColumn     = postgres.TimestampzColumn("start_at")
		EndAtColumn       = postgres.TimestampzColumn("end_at")
		DescriptionColumn = postgres.StringColumn("description")
		allColumns        = postgres.ColumnList{IDColumn, RoomNumberColumn, ClientIDColumn, StartAtColumn, EndAtColumn, DescriptionColumn}
		mutableColumns    = postgres.ColumnList{RoomNumberColumn, ClientIDColumn, StartAtColumn, EndAtColumn, DescriptionColumn}
	)

	return roomOccupanciesTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:          IDColumn,
		RoomNumber:  RoomNumberColumn,
		ClientID:    ClientIDColumn,
		StartAt:     StartAtColumn,
		EndAt:       EndAtColumn,
		Description: DescriptionColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
