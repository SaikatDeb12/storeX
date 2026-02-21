package models

import "time"

type Asset struct {
	ID            string    `json:"id" db:"id"`
	UserID        string    `json:"userID" db:"user_id"`
	Name          string    `json:"name" db:"name"`
	Brand         string    `json:"brand" db:"brand"`
	Model         string    `json:"model" db:"model"`
	SerialNo      string    `json:"serialNo" db:"serial_no"`
	Type          string    `json:"type" db:"type"`
	Status        string    `json:"status" db:"status"`
	Owner         string    `json:"owner" db:"owner"`
	AssignedByID  string    `json:"assignedByID" db:"assigned_by_id"`
	AssignedTo    time.Time `json:"assignedTo" db:"assigned_to"`
	WarrantyStart time.Time `json:"warrantyStart" db:"warranty_start"`
	ServiceStart  time.Time `json:"serviceStart" db:"service_start"`
	ReturnedOn    time.Time `json:"returnedOn" db:"returned_on"`
	CreatedAt     time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt     time.Time `json:"updatedAt" db:"updated_at"`
	ArchivedAt    time.Time `json:"archivedAt" db:"archived_at"`
	ArchivedBy    time.Time `json:"archivedBy" db:"archived_by"`
}

type Laptop struct {
	ID              string `json:"id" db:"id"`
	AssetID         string `json:"assetID" db:"asset_id"`
	Processor       string `json:"processor" db:"processor"`
	RAM             string `json:"ram" db:"ram"`
	Storage         string `json:"storage" db:"storage"`
	OperatingSystem string `json:"operatingSystem" db:"operating_system"`
	Charger         string `json:"charger" db:"charger"`
	DevicePassword  string `json:"devicePassword" db:"device_password"`
}

type Keyboard struct {
	ID           string `json:"id" db:"id"`
	AssetID      string `json:"assetID" db:"asset_id"`
	Layout       string `json:"layout" db:"layout"`
	Connectivity string `json:"connectivity" db:"connectivity"`
}

type Mouse struct {
	ID           string `json:"id" db:"id"`
	AssetID      string `json:"assetID" db:"asset_id"`
	DPI          string `json:"dpi" db:"dpi"`
	Connectivity string `json:"connectivity" db:"connectivity"`
}

type Mobile struct {
	ID              string `json:"id" db:"id"`
	AssetID         string `json:"assetID" db:"asset_id"`
	OperatingSystem string `json:"operatingSystem" db:"operating_system"`
	RAM             string `json:"connectivity" db:"connectivity"`
	Storage         string `json:"storage" db:"storage"`
	Charger         string `json:"charger" db:"charger"`
	DevicePassword  string `json:"devicePassword" db:"device_password"`
}
