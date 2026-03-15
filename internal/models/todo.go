package models

import "time"

type Todo struct {
	ID        int       `json:"id" db:"id"`
	Title     string    `json:"title" db:"title"`
	Completed bool      `json:"completed" db:"completed"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	UserID    string    `json:"user_id" db:"user_id"`
}

// json:"email"  → used when sending response in API  ----> use email , not Email
// db:"email"    → used when reading from database	  ----> use email , not Email
