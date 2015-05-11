package game

import (
	"time"
)

// Промежуток по времени между расчетами симуляции
const TICK = 100 * time.Millisecond

// Структура игрового мира
type World struct {
	// Высота
	Height int
	// Ширина
	Width int
	// Население
	Population []Entity
	// Текущий тик
	Tick int64
}

// Добавление новой сущности
func (world *World) AddEntity(entity Entity) {
	world.Population = append(world.Population, entity)
}

// Запуск игровой симуляции
func (world *World) Run() {
	for {
		// Ждем начало следующего тика
		time.Sleep(TICK)
		// Пробегаемся по всем сущностям и отдаем им управление
		for _, entity := range world.Population {
			if behavior := entity.GetCurrentBehavior(); behavior != nil {
				behavior.TickLogic(world, entity)
			}
		}

		world.Tick += 1
	}
}

// Генерация случайных координат, расположенных в игровом мире
func (world *World) RandomPosition() *Position {
	position := &Position{}
	position.X = RandomNumber(0, world.Width+1)
	position.Y = RandomNumber(0, world.Height+1)
	return position
}
