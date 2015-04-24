package game

// Игровые координаты
type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// Сравнение текущей координаты с переданной
func (position *Position) Equal(otherPosition *Position) bool {
	if position.X == otherPosition.X && position.Y == otherPosition.Y {
		return true
	}
	return false
}

// Двухмерный вектор
type Vector2d struct {
	X int
	Y int
}
