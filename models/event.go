package models

type Event struct {
	EventId   int64  `json:"eventId"`
	EventType string `json:"eventType"`
	UserId    int64  `json:"userID"`
	EventTime string `json:"eventTime"`
	PayLoad   string `json:"payload"`
}
