package storage

import (
	"L2/develop/dev11/model"
	"fmt"
	"time"
)

// Cache - хранилище Event
var Cache map[int]model.Event

// ID - id текущей записи в кеше
var ID int

// InitStorage - функция для инициализации кеша
func InitStorage() {
	Cache = make(map[int]model.Event)
}

// CreateEvent - функция для создания события в кеше
func CreateEvent(userID int, eventDate time.Time, description string) (model.Event, error) {
	newEvent := model.Event{
		ID:          ID,
		UserID:      userID,
		EventDate:   eventDate,
		Description: description}
	Cache[ID] = newEvent
	ID++
	return newEvent, nil
}

// UpdateEvent - функция для обновления события в кеше
func UpdateEvent(id, userID int, eventDate time.Time, description string) (model.Event, error) {
	if _, ok := Cache[id]; !ok {
		return model.Event{}, fmt.Errorf("event with id %d was not found", id)
	}
	event := Cache[id]
	event.UserID = userID
	event.EventDate = eventDate
	event.Description = description
	Cache[id] = event
	return event, nil
}

// DeleteEvent - функция для удаления события из кеша
func DeleteEvent(id int) error {
	if _, ok := Cache[id]; !ok {
		return fmt.Errorf("event with id %d was not found", id)
	}
	delete(Cache, id)
	return nil
}
