package game

type Chest struct {
    *_Entity
}

func NewChest(name string, position *Position) *Chest {
    chest := &Chest{}
    chest._Entity = NewEntity(name, position)
	chest.Type = "chest"
    return chest
}
