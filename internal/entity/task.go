package entity

import "time"

type Task struct {
	Id          int       `json:"id"`
	Header      string    `json:"header"`
	Description string    `json:"description"`
	Date        time.Time `json:"created_at"`
}
