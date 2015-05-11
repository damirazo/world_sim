package game

type chest struct {
	*entity
}

func NewChest(name string, position *Position) *chest {
    c := &chest{}
    c.entity = newEntity(name, position)
	c._type = "chest"
    return c
}
