package main


import (
    "log"
    "game"
    "net/http"
    "strconv"
    "html/template"
    "encoding/json"
    "github.com/gorilla/websocket"
)


// Порт для запуска сервера
const ADDR string = ":7777"

// Создание игрового мира
var world = game.World{Width: 40, Height: 40}


func main() {
    log.Print("Server started at 127.0.0.1" + ADDR)

    log.Println("Started entity creation...")
    // Создание и добавления в игровой мир сущностей
    x := 0
    for x < 1 {
        entity := game.NewEntity("Name #" + strconv.Itoa(x), world.RandomPostion())
        world.AddEntity(entity)
        x += 1
    }
    log.Println("All entities created!")

    // Запуск симуляции
    go world.Run()
    log.Println("Started world simulation...")

    // Инициализация http обработчиков
    http.HandleFunc("/", mainHandler)
    http.HandleFunc("/game", webSocketHandler)
    http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, r.URL.Path[1:])
    })

    log.Println("All handlers started...")

    // Запуск сервера
    if err := http.ListenAndServe(":7777", nil); err != nil {
        log.Fatal("ListenAndServe:", err)
    }
}


// Отображение главной страницы сайта
func mainHandler(w http.ResponseWriter, r *http.Request) {
    var homeTempl = template.Must(template.ParseFiles("templates/home.html"))
    homeTempl.Execute(w, nil)
}

// Обработчик соединений с websocket
func webSocketHandler(w http.ResponseWriter, r *http.Request) {
    ws, err := websocket.Upgrade(w, r, nil, 1024, 1024)
    defer ws.Close()

    if _, ok := err.(websocket.HandshakeError); ok {
        http.Error(w, "Not a websocket handshake", 400)
        return
    } else if err != nil {
        return
    }

    // Слушаем websocket в цикле
    for {
        _, message, err := ws.ReadMessage()
        if err != nil {
            break
        }

        // Смотрим, какая из полученных команд нам подходит
        switch string(message) {
            case "state":
            sendWorldState(ws);
            break;
        }
    }
}


// Отправка указанного сообщения в указанный сокет
func sendMessage(ws *websocket.Conn, message string) {
    go func() {
        err := ws.WriteMessage(websocket.TextMessage, []byte(message))
        if (err != nil) {
            ws.Close()
        }
    }()
}

// Сериализация и отправка координат всех сущностей подключенным клиентам
func sendWorldState(ws *websocket.Conn) {
    res, _ := json.Marshal(world.Population)
    go sendMessage(ws, string(res))
}