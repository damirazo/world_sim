package main

import (
	"encoding/json"
	"game"
	"github.com/gorilla/websocket"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// Порт для запуска сервера
const (
	PORT           = ":7777"
	COUNT_ENTITIES = 1
)

// Создание игрового мира
var world = game.World{Width: 40, Height: 40}

func main() {
	log.Print("Server started at 127.0.0.1" + PORT)
	log.Println("Started entity creation...")
	// Создание и добавления в игровой мир сущностей

	chest := game.NewChest("Chest", world.RandomPosition())
	world.AddEntity(chest)

	for x := 0; x < COUNT_ENTITIES; x++ {
		unit := game.NewUnit("Name #"+strconv.Itoa(x), world.RandomPosition())
		unit.Storage.Set("Chest", chest)
		world.AddEntity(unit)
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
	if err := http.ListenAndServe(PORT, nil); err != nil {
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
			sendWorldState(ws)
			break
		}
	}
}

// Отправка указанного сообщения в указанный сокет
func sendMessage(ws *websocket.Conn, message string) {
	go func() {
		err := ws.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			ws.Close()
		}
	}()
}

// Сериализация и отправка координат всех сущностей подключенным клиентам
func sendWorldState(ws *websocket.Conn) {
	res, _ := json.Marshal(world.Population)
	go sendMessage(ws, string(res))
}
