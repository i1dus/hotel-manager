package domain

type Employee struct {
	ID       int
	Username string
	Name     string
	Position Position
}

// Position - должность (позиция) работника.
type Position int

const (
	PositionUnknown      Position = 0
	PositionOwner        Position = 1
	PositionManager      Position = 2
	PositionReceptionist Position = 3
	PositionMaid         Position = 4
)

func (p Position) GetPositionName() PositionName {
	switch p {
	case PositionOwner:
		return PositionNameOwner
	case PositionManager:
		return PositionNameManager
	case PositionReceptionist:
		return PositionNameReceptionist
	case PositionMaid:
		return PositionNameMaid
	}
	return PositionNameUnknown
}

// PositionName - название должности (позиции) работника.
type PositionName string

const (
	PositionNameUnknown      PositionName = "неизвестный"
	PositionNameOwner        PositionName = "владелец"
	PositionNameManager      PositionName = "менеджер"
	PositionNameReceptionist PositionName = "ресепшионист"
	PositionNameMaid         PositionName = "горничный"
)
