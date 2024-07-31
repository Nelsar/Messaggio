package db

import (
	"log"

	"messaggio.com/models"
)

func AddEvent(event models.Event) error {
	connection, err := DBConnect()
	if err != nil {
		log.Println(err.Error())
	}

	_, err = connection.Exec("INSERT INTO events (eventId, eventType, userId, payload) VALUES ($1, $2, $3, $4)", event.EventId, event.EventType, event.UserId, event.PayLoad)
	if err != nil {
		log.Println(err.Error())
	}

	defer connection.Close()

	return err
}

func UpdateEvent(event models.Event) error {
	connection, err := DBConnect()
	if err != nil {
		log.Println(err)
	}

	_, err = connection.Exec("UPDATE events SET eventType = $1, userId = $2, payload = $3 WHERE eventId = $4", event.EventType, event.UserId, event.PayLoad, event.EventId)
	if err != nil {
		log.Println(err)
	}

	defer connection.Close()

	return err
}

func GetEvents() ([]models.Event, error) {
	connection, err := DBConnect()
	if err != nil {
		log.Println(err)
	}

	rows, err := connection.Query("SELECT eventId, eventType, userId, eventTime, payload FROM events")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	events := []models.Event{}

	for rows.Next() {
		event := models.Event{}

		err := rows.Scan(&event.EventId, &event.EventType, &event.UserId, &event.EventTime, &event.PayLoad)
		if err != nil {
			log.Println(err)
		}

		events = append(events, event)
	}

	return events, err
}
