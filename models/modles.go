package models

import (
    "time"

    "github.com/google/uuid"
)

// User represents the user table in the database
type User struct {
    UserID    uuid.UUID `json:"user_id"`
    UserName  string    `json:"user_name"`
    Password  string    `json:"password"`
    CreatedAt time.Time `json:"created_at"`
}

// Todo represents the todo table in the database
type Todo struct {
    TodoID      uuid.UUID `json:"todo_id"`
    Title       string    `json:"title"`
    Description string    `json:"description"`
    IsCompleted bool      `json:"is_completed"`
    CreatedAt   time.Time `json:"created_at"`
    UserID      uuid.UUID `json:"user_id"`
}
