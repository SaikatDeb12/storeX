package dbhelper

import (
	"github.com/SaikatDeb12/storeX/internal/database"
	"github.com/SaikatDeb12/storeX/internal/models"
)

func InsertKeyboardDetails(assetID string, req models.KeyboardRequest) error {
	SQL := `
		INSERT INTO keyboards(connectivity, layout)
		VALUES($1, $2)
	`

	args := []interface{}{
		req.Connectivity,
		req.Layout,
	}

	_, err := database.DB.Exec(SQL, args...)
	return err
}

func InsertLaptopDetails(assetID string, req models.LaptopRequest) error {
	SQL := `
		INSERT INTO laptops(processor, ram, storage, operating_sysem, charger, device_password)
		VALUES($1, $2, $3, $4, $5, $6)
	`

	args := []interface{}{
		req.Processor,
		req.RAM,
		req.Storage,
		req.OperatingSystem,
		req.Charger,
		req.DevicePassword,
	}

	_, err := database.DB.Exec(SQL, args...)
	return err
}

func InsertMouseDetails(assetID string, req models.MouseRequest) error {
	SQL := `
		INSERT INTO mice(dpi, connectivity)
		VALUES($1, $2)
	`

	args := []interface{}{
		req.DPI,
		req.Connectivity,
	}

	_, err := database.DB.Exec(SQL, args...)
	return err
}

func InsertMobileDetails(assetID string, req models.MobileRequest) error {
	SQL := `
		INSERT INTO mobiles(operating_system, ram, storage, charger, device_password)
		VALUES($1, $2, $3, $4, $5)
	`

	args := []interface{}{
		req.OperatingSystem,
		req.RAM,
		req.Storage,
		req.Charger,
		req.DevicePassword,
	}

	_, err := database.DB.Exec(SQL, args...)
	return err
}

func CreateAsset(model models.CreateAssetRequest) (string, error) {
	SQL := `
		INSERT INTO assets(brand, model, serial_number, asset_type, status, owner_type, warranty_start, warranty_end)
		VALUES($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id
	`

	args := []interface{}{
		model.Brand,
		model.Model,
		model.SerialNumber,
		model.Type,
		model.Status,
		model.Owner,
		model.WarrantyStart,
		model.WarrantyEnd,
	}

	var assetID string
	err := database.DB.Get(&assetID, SQL, args...)
	if err != nil {
		return assetID, err
	}
	return assetID, nil
}
