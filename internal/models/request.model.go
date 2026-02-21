package models

type RegisterRequest struct {
	Name        string `json:"name" db:"name"`
	Email       string `json:"email" db:"email"`
	PhoneNumber string `json:"phoneNumber" db:"phone_number"`
	Role        string `json:"role" db:"role"`
	Type        string `json:"type" db:"type"`
	Password    string `json:"password" db:"password"`
}

type LoginRequest struct {
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}
