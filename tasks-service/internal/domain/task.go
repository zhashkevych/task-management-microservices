package domain

import "time"

type (
	Task struct {
		Id        int       `json:"id" db:"id"`
		Title     string    `json:"title" db:"title" binding:"required,min=3"`
		CreatedAt time.Time `json:"created_at" db:"created_at"`
		UserId    int       `json:"user_id" db:"user_id"`
	}
)
