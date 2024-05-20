package models

import "time"

type Message struct {
	ID         string     `json:"id"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty"`
	Content    string     `json:"content"`
	IsCompleted bool      `json:"is_completed"`
}

