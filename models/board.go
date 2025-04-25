package models

import "time"

type Board struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	IsDefault bool      `json:"is_default"`
	Tasks     []*Task   `json:"sections"`
	CreatedAt time.Time `json:"created_at"`
}
