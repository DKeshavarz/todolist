package models

import "time"

type Task struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	IsDone      bool      `json:"is_done"`
	Label       string    `json:"label"`
	CreatedAt   time.Time `json:"created_at"`
}
