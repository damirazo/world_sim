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
func ChestRun(world *World, e Entity) {
	u := e.(*unit)
	ch := u.Get("chest").(*chest)
	pos := u.GetPosition()
	speed := u.speed

	if ch == nil {
		return
	} else if pos.Equal(ch.position) {
		ch.SetPosition(world.RandomPosition())
		return
	}

	target := ch.GetPosition()

	if pos.X > target.X {
		pos.X -= speed
		if pos.X <= target.X {
			pos.X = target.X
		}
	}

	if pos.X < target.X {
		pos.X += speed
		if pos.X >= target.X {
			pos.X = target.X
		}
	}

	if pos.Y > target.Y {
		pos.Y -= speed
		if pos.Y <= target.Y {
			pos.Y = target.Y
		}
	}

	if pos.Y < target.Y {
		pos.Y += speed
		if pos.Y >= target.Y {
			pos.Y = target.Y
		}
	}
}
