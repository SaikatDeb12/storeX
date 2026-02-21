package models

type RegisterRequest struct {
	Name        string `json:"name" db:"name" validate:"required, min=3, max=100"`
	Email       string `json:"email" db:"email" validate:"required, email"`
	PhoneNumber string `json:"phoneNumber" db:"phone_number" validate:"required,len=10"`
	Role        string `json:"role" db:"role" validate:"required, oneof='admin employee project-manager asset-manager employee-manager'"`
	Type        string `json:"type" db:"type" validate:"required, oneof='client remotestate'"`
	Password    string `json:"password" db:"password" validate:"required, min=8, max=20"`
}

type LoginRequest struct {
	Email    string `json:"email" db:"email" validate:"required, email"`
	Password string `json:"password" db:"password" validate:"required, min=8, max=20"`
}
