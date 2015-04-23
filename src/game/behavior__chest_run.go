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
func ChestRun(world *World, entity *Entity) {
    hasTarget := entity.Storage.Get("HasTarget", false).(bool)
    randomPosition := world.RandomPostion()

    if !hasTarget {
        entity.Storage.Set("HasTarget", true)
        entity.Storage.Set("Target", randomPosition)
    } else if (entity.Position.Equal(entity.Storage.Get("Target", &Position{}).(*Position))) {
        entity.Storage.Set("HasTarget", false)
        return
    }

    target := entity.Storage.Get("Target", &Position{}).(*Position)

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