package models

import "time"

type Book struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Title     string    `json:"title"`
	Isbn      string    `json:"isbn"`
	Writter   string    `json:"writter"`
}
