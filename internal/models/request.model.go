package models

type RegisterRequest struct {
	Name        string `json:"name" validate:"required, min=3, max=100"`
	Email       string `json:"email" validate:"required, email"`
	PhoneNumber string `json:"phoneNumber" validate:"required,len=10"`
	UserRole    string `json:"userRole" validate:"required, oneof='admin employee project-manager asset-manager employee-manager'"`
	UserType    string `json:"userType" validate:"required, oneof='client remotestate'"`
	Password    string `json:"password" validate:"required, min=8, max=20"`
}

type LoginRequest struct {
	Email    string `json:"email" db:"email" validate:"required, email"`
	Password string `json:"password" db:"password" validate:"required, min=8, max=20"`
}
