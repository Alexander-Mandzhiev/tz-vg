package entity

import "time"

type Todo struct {
	ID          int       `json:"id,omitempty" db:"id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	DueDate     time.Time `json:"dueDate" db:"due_date"`
	CreatedAt   time.Time `json:"createdAt,omitempty" db:"created_at"`
	UpdatedAt   time.Time `json:"updatedAt,omitempty" db:"updated_at"`
}
