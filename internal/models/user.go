package models

type User struct {
	ID        string `json:"id" db:"id"`
	Email     string `json:"email" db:"email"`
	Password  string `json:"password" db:"password"`
	CreatedAt string `json:"created_at" db:"created_at"`
	UpdatedAt string `json:"updated_at" db:"updated_at"`
}

// json:"email"  → used when sending response in API  ----> use email , not Email
// db:"email"    → used when reading from database	  ----> use email , not Email
