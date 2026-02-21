package models

import (
	"time"
)

type User struct {
	ID          string    `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Email       string    `json:"email" db:"email"`
	PhoneNumber string    `json:"phoneNumber" db:"phone_number"`
	Role        string    `json:"role" db:"role"`
	Type        string    `json:"type" db:"type"`
	Password    string    `json:"password" db:"password"`
	CreatedAt   time.Time `json:"createdAt" db:"create_at"`
	ArchivedAt  time.Time `json:"ArchivedAt" db:"archived_at"`
}

type Session struct {
	ID         string    `db:"id"`
	UserID     string    `db:"user_id"`
	CreatedAt  time.Time `db:"created_at"`
	ArchivedAt time.Time `db:"archived_at"`
}
