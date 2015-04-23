package game


// Структура, хранящая поведение сущности
type Behavior struct {
    // Наименование поведения
    Name string
    // Логика поведения
    TickLogic func(*World, *Entity)
}