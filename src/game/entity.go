package game

import "encoding/json"

// Количество пропускаемых тиков
const SKIP_TICK = 0

type Entity interface {
    GetPosition() *Position
    SetPosition(*Position)

    GetCurrentBehavior() *Behavior
    SetCurrentBehavior(*Behavior)

    Get(string) interface{}
    GetWithDefault(string, interface{}) interface{}
    Set(string, interface{})
}

type entity struct {
    _type    string
    state    string
    name     string
    position *Position
    // Тик предыдущего обновления
    prevTick        int64
    currentBehavior *Behavior
    context         map[string]interface{}
}

func newEntity(name string, position *Position) *entity {
    e := &entity{}
    e.name = name
    e.position = position
    e.context = map[string]interface{}{}
    return e
}

// Position
// ---------------------------------------------------------------------------

func (e *entity) GetPosition() *Position {
    return e.position
}

func (e *entity) SetPosition(position *Position) {
    e.position = position
}

// CurrentBehavior
// ---------------------------------------------------------------------------

func (e *entity) SetCurrentBehavior(behavior *Behavior) {
    e.currentBehavior = behavior
}

func (e *entity) GetCurrentBehavior() *Behavior {
    return e.currentBehavior
}

// Context
// ---------------------------------------------------------------------------

func (e *entity) Get(param string) interface{} {
    return e.context[param]
}

func (e *entity) GetWithDefault(param string, _default interface{}) interface{} {
    if val := e.Get(param); val != nil {
        return val
    }
    return _default
}

func (e *entity) Set(param string, value interface{}) {
    e.context[param] = value
}

// Interface Marshaller
// ---------------------------------------------------------------------------

func (e *entity) MarshalJSON() ([]byte, error) {
    return json.Marshal(struct {
        Type     string    `json:"type"`
        Name     string    `json:"name"`
        Position *Position `json:"position"`
    }{
        e._type,
        e.name,
        e.position,
    })
}
