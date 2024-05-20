package models

import "time"

type Todo struct {
	ID                   string    `json:"id"`
	UserID               string    `json:"userId"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
	DeletedAt            *time.Time `json:"deleted_at,omitempty"`
	CompletionPercentage int       `json:"completion_percentage"`
	Messages             []Message `json:"messages"`
}
