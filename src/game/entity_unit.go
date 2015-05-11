package game

type unit struct {
	*entity
	speed int
}

func NewUnit(name string, position *Position) *unit {
	u := &unit{speed: 1}
	u.entity = newEntity(name, position)
	u._type = "unit"

	b := &Behavior{}
	b.Name = "Chest Run"
	b.TickLogic = ChestRun

	u.SetCurrentBehavior(b)

	return u
}
