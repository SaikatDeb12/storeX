package models

import (
	"time"
)

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
	Role      string `json:"role"`
}

type UserInfoRequest struct {
	ID           string             `json:"id" db:"id"`
	Name         string             `json:"name" db:"name" validate:"required,min=3,max=50"`
	Email        string             `json:"email" db:"email" validate:"required,email"`
	PhoneNumber  string             `json:"phoneNumber" db:"phone_number" validate:"required,len=10"`
	Role         string             `json:"role" db:"role" validate:"required"`
	Employment   string             `json:"employment" db:"employment" validate:"required"`
	CreatedAt    string             `json:"createdAt" db:"created_at" validate:"required"`
	AssetDetails []AssetInfoRequest `json:"assetDetails"`
}

type AssetInfoRequest struct {
	ID     string `json:"id" db:"id"`
	Brand  string `json:"brand" db:"brand"`
	Model  string `json:"model" db:"model"`
	Status string `json:"status" db:"status"`
	Type   string `json:"assetType" db:"asset_type"`
}

type CreateAssetRequest struct {
	Brand         string    `json:"brand" db:"brand" validate:"required"`
	Model         string    `json:"model" db:"model" validate:"required"`
	SerialNumber  string    `json:"serialNumber" db:"serial_number" validate:"required"`
	Type          string    `json:"assetType" db:"asset_type" validate:"required,oneof=laptop keyboard mouse mobile"`
	Status        string    `json:"status" db:"status" validate:"required,oneof=available assigned in_service under_repair damaged"`
	Owner         string    `json:"owner" db:"owner_type" validate:"required,oneof=client remotestate"`
	WarrantyStart time.Time `json:"warrantyStart" db:"warranty_start" validate:"required"`
	WarrantyEnd   time.Time `json:"warrantyEnd" db:"warranty_end" validate:"required"`

	Laptop   *LaptopRequest   `json:"laptop"`
	Keyboard *KeyboardRequest `json:"keyboard"`
	Mouse    *MouseRequest    `json:"mouse"`
	Mobile   *MobileRequest   `json:"mobile"`
}

type LaptopRequest struct {
	Processor       string  `json:"processor" db:"processor"`
	RAM             string  `json:"ram" db:"ram"`
	Storage         string  `json:"storage" db:"storage"`
	OperatingSystem string  `json:"operatingSystem" db:"operating_system"`
	Charger         *string `json:"charger" db:"charger"`
	DevicePassword  string  `json:"devicePassword" db:"device_password"`
}

type KeyboardRequest struct {
	Layout       *string `json:"layout" db:"layout"`
	Connectivity string  `json:"connectivity" db:"connectivity"`
}

type MouseRequest struct {
	DPI          *int   `json:"dpi" db:"dpi"`
	Connectivity string `json:"connectivity" db:"connectivity"`
}

type MobileRequest struct {
	OperatingSystem string  `json:"operatingSystem" db:"operating_system"`
	RAM             string  `json:"ram" db:"ram"`
	Storage         string  `json:"storage" db:"storage"`
	Charger         *string `json:"charger" db:"charger"`
	DevicePassword  string  `json:"devicePassword" db:"device_password"`
}
