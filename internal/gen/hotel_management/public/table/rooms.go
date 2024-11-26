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

var Rooms = newRoomsTable("public", "rooms", "")

type roomsTable struct {
	postgres.Table

	// Columns
	ID    postgres.ColumnInteger
	InUse postgres.ColumnBool
	Type  postgres.ColumnInteger

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type RoomsTable struct {
	roomsTable

	EXCLUDED roomsTable
}

// AS creates new RoomsTable with assigned alias
func (a RoomsTable) AS(alias string) *RoomsTable {
	return newRoomsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new RoomsTable with assigned schema name
func (a RoomsTable) FromSchema(schemaName string) *RoomsTable {
	return newRoomsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new RoomsTable with assigned table prefix
func (a RoomsTable) WithPrefix(prefix string) *RoomsTable {
	return newRoomsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new RoomsTable with assigned table suffix
func (a RoomsTable) WithSuffix(suffix string) *RoomsTable {
	return newRoomsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newRoomsTable(schemaName, tableName, alias string) *RoomsTable {
	return &RoomsTable{
		roomsTable: newRoomsTableImpl(schemaName, tableName, alias),
		EXCLUDED:   newRoomsTableImpl("", "excluded", ""),
	}
}

func newRoomsTableImpl(schemaName, tableName, alias string) roomsTable {
	var (
		IDColumn       = postgres.IntegerColumn("id")
		InUseColumn    = postgres.BoolColumn("in_use")
		TypeColumn     = postgres.IntegerColumn("type")
		allColumns     = postgres.ColumnList{IDColumn, InUseColumn, TypeColumn}
		mutableColumns = postgres.ColumnList{InUseColumn, TypeColumn}
	)

	return roomsTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:    IDColumn,
		InUse: InUseColumn,
		Type:  TypeColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
