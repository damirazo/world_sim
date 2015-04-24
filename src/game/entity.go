package game

// Количество пропускаемых тиков
const SKIP_TICK = 0

type Entity interface {
	GetPosition() *Position
	SetPosition(*Position)

	GetCurrentBehavior() *Behavior
	SetCurrentBehavior(*Behavior)
}

// Структура сущности
type _Entity struct {
	// Тип сущности
	Type string `json:"type"`
	// Имя
	Name string `json:"name"`
	// Текущие координаты
	Position *Position `json:"position"`
	// Тик предыдущего обновления
	prevTick int64 `json:"-"`
	// Текущее поведение
	CurrentBehavior *Behavior `json:"-"`
	// Дополнительные параметры сущности
	Storage *EntityParamStorage `json:"-"`
}

// Создание экземпляра структур
func NewEntity(name string, position *Position) *_Entity {
	entity := &_Entity{}
	entity.Name = name
	entity.Position = position
	entity.Storage = &EntityParamStorage{}
	return entity
}

// Position
// ---------------------------------------------------------------------------

func (entity *_Entity) GetPosition() *Position {
	return entity.Position
}

func (entity *_Entity) SetPosition(position *Position) {
	entity.Position = position
}

// CurrentBehavior
// ---------------------------------------------------------------------------

func (entity *_Entity) SetCurrentBehavior(behavior *Behavior) {
	entity.CurrentBehavior = behavior
}

func (entity *_Entity) GetCurrentBehavior() *Behavior {
	return entity.CurrentBehavior
}

// ---------------------------------------------------------------------------

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
	Key   string
	Value interface{}
}
