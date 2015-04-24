package game

type Unit struct {
	*_Entity
	// Speed of unit at per tick
	Speed int
}

func NewUnit(name string, position *Position) *Unit {
	unit := &Unit{Speed: 1}
	unit._Entity = NewEntity(name, position)
	unit.Type = "unit"

	b := &Behavior{}
	b.Name = "Chest Run"
	b.TickLogic = ChestRun

	unit.SetCurrentBehavior(b)

	return unit
}
