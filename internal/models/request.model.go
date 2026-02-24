package models

type RegisterRequest struct {
	Name        string `json:"name" validate:"required,min=3,max=50"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required,min=8,max=20"`
	PhoneNumber string `json:"phoneNumber" validate:"required,len=10"`
	Role        string `json:"role" validate:"required,oneof=admin employee project_manager asset_manager employee_manager"`
	Employment  string `json:"employment" validate:"required,oneof=full_time intern freelancer"`
}

type LoginRequest struct {
	Email    string `json:"email" db:"email" validate:"required,email"`
	Password string `json:"password" db:"password" validate:"required,min=8,max=20"`
}

type RequestContext struct {
	UserID    string `json:"user_id"`
	SessionID string `json:"session_id"`
}

type UserInfoRequest struct {
	ID          string `json:"id" db:"id"`
	Name        string `json:"name" db:"name" validate:"required,min=3,max=50"`
	Email       string `json:"email" db:"email" validate:"required,email"`
	PhoneNumber string `json:"phoneNumber" db:"phone_number" validate:"required,len=10"`
	Role        string `json:"role" db:"role" validate:"required,oneof=admin employee project_manager asset_manager employee_manager"`
	Employment  string `json:"employment" db:"employment" validate:"required,oneof=full_time intern freelancer"`
	CreatedAt   string `json:"createdAt" db:"created_at"`
}
