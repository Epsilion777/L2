package main

import (
	"L2/develop/dev11/model"
	"L2/develop/dev11/storage"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

/*
=== HTTP server ===

Реализовать HTTP сервер для работы с календарем. В рамках задания необходимо работать строго со стандартной HTTP библиотекой.
В рамках задания необходимо:
 1. Реализовать вспомогательные функции для сериализации объектов доменной области в JSON.
 2. Реализовать вспомогательные функции для парсинга и валидации параметров методов /create_event и /update_event.
 3. Реализовать HTTP обработчики для каждого из методов API, используя вспомогательные функции и объекты доменной области.
 4. Реализовать middleware для логирования запросов

Методы API:

	POST /create_event
	POST /update_event
	POST /delete_event
	GET /events_for_day
	GET /events_for_week
	GET /events_for_month

Параметры передаются в виде www-url-form-encoded (т.е. обычные user_id=3&date=2019-09-09).
В GET методах параметры передаются через queryString, в POST через тело запроса.
В результате каждого запроса должен возвращаться JSON документ содержащий либо {"result": "..."} в случае успешного выполнения метода,
либо {"error": "..."} в случае ошибки бизнес-логики.

В рамках задачи необходимо:
 1. Реализовать все методы.
 2. Бизнес логика НЕ должна зависеть от кода HTTP сервера.
 3. В случае ошибки бизнес-логики сервер должен возвращать HTTP 503. В случае ошибки входных данных (невалидный int например) сервер должен возвращать HTTP 400. В случае остальных ошибок сервер должен возвращать HTTP 500. Web-сервер должен запускаться на порту указанном в конфиге и выводить в лог каждый обработанный запрос.
 4. Код должен проходить проверки go vet и golint.
*/

// Config - структура для считывания конфигурационного файла
type Config struct {
	Port int `json:"port"`
}

// Response - ответ для сериализации ошибки
type Response struct {
	Error string `json:"error"`
}

// Logger - middleware, который логирует запросы
type Logger struct {
	handler http.Handler
}

func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("method: %s | url: %s", r.Method, r.URL)
	l.handler.ServeHTTP(w, r)
}

// NewLogger - конструктор для Logger
func NewLogger(handler http.Handler) *Logger {
	return &Logger{handler: handler}
}

// CreateEvent - обработчик для создания события
func CreateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			SendError(w, r, err.Error(), http.StatusBadRequest)
			return
		}

		// Получаем данные для Event из формы
		userID := r.Form.Get("user_id")
		date := r.Form.Get("date")
		description := r.Form.Get("description")

		// Создаем структуру event
		e, err := NewEvent("", userID, date, description)
		if err != nil {
			SendError(w, r, err.Error(), http.StatusBadRequest)
			return
		}

		// Добавляем event в кеш
		event, err := storage.CreateEvent(e.UserID, e.EventDate, e.Description)
		if err != nil {
			SendError(w, r, err.Error(), http.StatusServiceUnavailable)
			return
		}

		// Формирует ответ
		response := map[string]model.Event{"result": event}
		w.Header().Set("Content-Type", "application/json")

		encoder := json.NewEncoder(w)
		if err := encoder.Encode(response); err != nil {
			SendError(w, r, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		SendError(w, r, "Method not allowerd", http.StatusMethodNotAllowed)
		return
	}
}

// UpdateEvent - обработчик для обновления события
func UpdateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			SendError(w, r, err.Error(), http.StatusBadRequest)
			return
		}

		// Получаем данные для Event из формы
		ID := r.Form.Get("id")
		userID := r.Form.Get("user_id")
		date := r.Form.Get("date")
		description := r.Form.Get("description")

		// Создаем структуру event
		e, err := NewEvent(ID, userID, date, description)
		if err != nil {
			SendError(w, r, err.Error(), http.StatusBadRequest)
			return
		}

		// Добавляем event в кеш
		event, err := storage.UpdateEvent(e.ID, e.UserID, e.EventDate, e.Description)
		if err != nil {
			SendError(w, r, err.Error(), http.StatusServiceUnavailable)
			return
		}

		// Формирует ответ
		response := map[string]model.Event{"result": event}
		w.Header().Set("Content-Type", "application/json")

		encoder := json.NewEncoder(w)
		if err := encoder.Encode(response); err != nil {
			SendError(w, r, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		SendError(w, r, "Method not allowerd", http.StatusMethodNotAllowed)
		return
	}
}

// DeleteEvent - обработчик для удаления события
func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			SendError(w, r, err.Error(), http.StatusBadRequest)
			return
		}

		// Получаем ID event и конвертируем в int
		ID := r.Form.Get("id")
		intID, err := strconv.Atoi(ID)
		if err != nil {
			SendError(w, r, err.Error(), http.StatusBadRequest)
			return
		}

		// Удаляем event из кеша
		err = storage.DeleteEvent(intID)
		if err != nil {
			SendError(w, r, err.Error(), http.StatusServiceUnavailable)
			return
		}

		// Формирует ответ
		response := map[string]string{"result": "success"}
		w.Header().Set("Content-Type", "application/json")

		encoder := json.NewEncoder(w)
		if err := encoder.Encode(response); err != nil {
			SendError(w, r, err.Error(), http.StatusInternalServerError)
			return
		}

	} else {
		SendError(w, r, "Method not allowerd", http.StatusMethodNotAllowed)
		return
	}
}

// GetEventsForDay - обработчик для получения событий за день
func GetEventsForDay(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Получаем id для Event из queryString и конвертируем в int
		userID := r.URL.Query().Get("user_id")
		intUserID, err := strconv.Atoi(userID)
		if err != nil {
			SendError(w, r, err.Error(), http.StatusBadRequest)
			return
		}

		// Создаем слайс userEvents, котоыре есть у пользователя
		userEvents := make([]model.Event, 0)
		for _, v := range storage.Cache {
			if v.UserID == intUserID {
				userEvents = append(userEvents, v)
			}
		}

		// Создаем слайс filteredEvents, котоыре удовлетворяют условию <= 24 часов
		filteredEvents := make([]model.Event, 0)
		now := time.Now()
		// Фильтруем события, оставляя только те, которые произошли в последние 24 часа
		for _, event := range userEvents {
			if now.Sub(event.EventDate) <= 24*time.Hour {
				filteredEvents = append(filteredEvents, event)
			}
		}

		// Формирует ответ
		response := map[string][]model.Event{"result": filteredEvents}
		w.Header().Set("Content-Type", "application/json")

		encoder := json.NewEncoder(w)
		if err := encoder.Encode(response); err != nil {
			SendError(w, r, err.Error(), http.StatusInternalServerError)
			return
		}

	} else {
		SendError(w, r, "Method not allowerd", http.StatusMethodNotAllowed)
	}
}

// GetEventsForWeek - обработчик для получения событий за неделю
func GetEventsForWeek(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Получаем id для Event из queryString и конвертируем в int
		userID := r.URL.Query().Get("user_id")
		intUserID, err := strconv.Atoi(userID)
		if err != nil {
			SendError(w, r, err.Error(), http.StatusBadRequest)
			return
		}

		// Создаем слайс userEvents, котоыре есть у пользователя
		userEvents := make([]model.Event, 0)
		for _, v := range storage.Cache {
			if v.UserID == intUserID {
				userEvents = append(userEvents, v)
			}
		}

		// Создаем слайс filteredEvents, котоыре удовлетворяют условию <= недели
		filteredEvents := make([]model.Event, 0)
		now := time.Now()
		// Фильтруем события, оставляя только те, которые произошли в последнюю неделю
		for _, event := range userEvents {
			if now.Sub(event.EventDate) <= 24*7*time.Hour {
				filteredEvents = append(filteredEvents, event)
			}
		}

		// Формирует ответ
		response := map[string][]model.Event{"result": filteredEvents}
		w.Header().Set("Content-Type", "application/json")

		encoder := json.NewEncoder(w)
		if err := encoder.Encode(response); err != nil {
			SendError(w, r, err.Error(), http.StatusInternalServerError)
			return
		}

	} else {
		SendError(w, r, "Method not allowerd", http.StatusMethodNotAllowed)
	}
}

// GetEventsForMonth - обработчик для получения событий за месяц
func GetEventsForMonth(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Получаем id для Event из queryString и конвертируем в int
		userID := r.URL.Query().Get("user_id")
		intUserID, err := strconv.Atoi(userID)
		if err != nil {
			SendError(w, r, err.Error(), http.StatusBadRequest)
			return
		}

		// Создаем слайс userEvents, котоыре есть у пользователя
		userEvents := make([]model.Event, 0)
		for _, v := range storage.Cache {
			if v.UserID == intUserID {
				userEvents = append(userEvents, v)
			}
		}

		// Создаем слайс filteredEvents, котоыре удовлетворяют условию <= месяца
		filteredEvents := make([]model.Event, 0)
		now := time.Now()
		// Фильтруем события, оставляя только те, которые произошли в последний месяц
		for _, event := range userEvents {
			if now.Sub(event.EventDate) <= 24*30*time.Hour {
				filteredEvents = append(filteredEvents, event)
			}
		}

		// Формирует ответ
		response := map[string][]model.Event{"result": filteredEvents}
		w.Header().Set("Content-Type", "application/json")

		encoder := json.NewEncoder(w)
		if err := encoder.Encode(response); err != nil {
			SendError(w, r, err.Error(), http.StatusInternalServerError)
			return
		}

	} else {
		SendError(w, r, "Method not allowerd", http.StatusMethodNotAllowed)
	}
}

// NewEvent - конструктор для Event
func NewEvent(id, userID, date, description string) (model.Event, error) {
	var intID int
	var err error
	// Если id передавали, то конверитруем в int
	if id != "" {
		intID, err = strconv.Atoi(id)
	}

	if err != nil {
		return model.Event{}, err
	}

	intUserID, err := strconv.Atoi(userID)
	if err != nil {
		return model.Event{}, err
	}

	validDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return model.Event{}, err
	}

	return model.Event{
		ID:          intID,
		UserID:      intUserID,
		EventDate:   validDate,
		Description: description}, nil

}

// SendError - функция для отправки ошибки в виде JSON
func SendError(w http.ResponseWriter, r *http.Request, err string, ststusCode int) {
	response := Response{
		Error: err,
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(ststusCode)

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(response); err != nil {
		http.Error(w, "Error encoding JSON", ststusCode)
		return
	}
}

func main() {
	// Открываем конфиг
	file, err := os.Open("develop/dev11/config.json")
	if err != nil {
		log.Fatalln("Error opening config file:", err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	config := Config{}
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatalln("Error decoding config file:", err)
	}

	// Инициализируем кеш
	storage.InitStorage()

	// Проиписываем роуты и обработчиков
	http.HandleFunc("/create_event", CreateEvent)
	http.HandleFunc("/update_event", UpdateEvent)
	http.HandleFunc("/delete_event", DeleteEvent)
	http.HandleFunc("/events_for_day", GetEventsForDay)
	http.HandleFunc("/events_for_week", GetEventsForWeek)
	http.HandleFunc("/events_for_month", GetEventsForMonth)

	// Запускаем прослушивание на порту из конфига, а так же устанавливаем middleware - Logger
	http.ListenAndServe(fmt.Sprintf(":%d", config.Port), NewLogger(http.DefaultServeMux))
}
