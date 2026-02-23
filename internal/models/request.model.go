package models

type RegisterRequest struct {
	Name        string `json:"name" validate:"required, min=3, max=100"`
	Email       string `json:"email" validate:"required, email"`
	Password    string `json:"password" validate:"required, min=8, max=20"`
	PhoneNumber string `json:"phoneNumber" validate:"required,len=10"`
	Role        string `json:"role" validate:"required, oneof='admin employee project_manager asset_manager employee_manager'"`
	Employment  string `json:"employment" validate:"required, oneof='full_time intern freelancer'"`
}

type LoginRequest struct {
	Email    string `json:"email" db:"email" validate:"required, email"`
	Password string `json:"password" db:"password" validate:"required, min=8, max=20"`
}
