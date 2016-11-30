package game


// Количество пропускаемых тиков
const SKIP_TICK = 0


// Структура сущности
type Entity struct {
    // Имя
    Name string                             `json:"name"`
    // Текущие координаты
    Position *Position                      `json:"position"`
    // Скорость перемещения (за тик)
    Speed int                               `json:"-"`
    // Тик предыдущего обновления
    prevTick int64                          `json:"-"`
    // Текущее поведение
    CurrentBehavior *Behavior               `json:"-"`
    // Дополнительные параметры сущности
    Storage *EntityParamStorage             `json:"storage"`
}

// Создание экземпляра структуры
func NewEntity(name string, position *Position) *Entity {
    entity := &Entity{}
    entity.Name = name
    entity.Position = position
    entity.Speed = 1
    entity.Storage = &EntityParamStorage{}

    // TODO: Тестовая логика
    b := &Behavior{}
    b.Name = "Chest Run"
    b.TickLogic = ChestRun
    entity.AddBehavior(b)

    return entity
}

// Добавление нового поведения сущности
func (entity *Entity) AddBehavior(behavior *Behavior) {
    entity.CurrentBehavior = behavior
}


// Хранилище дополнительных параметров сущности
type EntityParamStorage struct {
    Params []*EntityParam
}

// Получение из хранилища параметра с указанным ключем
func (storage *EntityParamStorage) Get(key string, def interface{}) interface{} {
    for _, param := range storage.Params {
        if param.Key == key {
            return param.Value
        }
    }

    return def
}

// Добавление / замена параметра в хранилище
// Возвращает true, если был создан новый параметр, false при замене существующего параметра
func (storage *EntityParamStorage) Set(key string, value interface{}) bool {
    for _, param := range storage.Params {
        if param.Key == key {
            param.Value = value
            return false
        }
    }

    newParam := &EntityParam{Key: key, Value: value}
    storage.Params = append(storage.Params, newParam)

    return true
}


// Параметр сущности
type EntityParam struct {
    Key string
    Value interface{}
}