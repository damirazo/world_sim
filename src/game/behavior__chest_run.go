package game

// Поведение сущности
// -------------------
// Выбирается случайное место на карте, в котором создается сундук,
// персонаж начинает перемещаться к данному сундуку, до тех пор, пока не доберется до него.
// После этого создается новый сундук и погоня за ним продолжается
// -------------------
// Используемые параметры:
// 1. HasTarget - выбрал ли персонаж себе цель
// 2. Target    - координаты выбранной цели
// -------------------
func ChestRun(world *World, _entity Entity) {
	entity := _entity.(*Unit)
	chest := entity.Storage.Get("Chest", nil).(*Chest)

	if chest == nil {
		return
	} else if entity.Position.Equal(chest.Position) {
		chest.SetPosition(world.RandomPosition())
		return
	}

	target := chest.GetPosition()

	if entity.Position.X > target.X {
		entity.Position.X -= entity.Speed
		if entity.Position.X <= target.X {
			entity.Position.X = target.X
		}
	}

	if entity.Position.X < target.X {
		entity.Position.X += entity.Speed
		if entity.Position.X >= target.X {
			entity.Position.X = target.X
		}
	}

	if entity.Position.Y > target.Y {
		entity.Position.Y -= entity.Speed
		if entity.Position.Y <= target.Y {
			entity.Position.Y = target.Y
		}
	}

	if entity.Position.Y < target.Y {
		entity.Position.Y += entity.Speed
		if entity.Position.Y >= target.Y {
			entity.Position.Y = target.Y
		}
	}
}
