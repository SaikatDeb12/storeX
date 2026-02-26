package dbhelper

import (
	"errors"
	"fmt"

	"github.com/SaikatDeb12/storeX/internal/database"
	"github.com/SaikatDeb12/storeX/internal/models"
)

func InsertKeyboardDetails(assetID string, req models.KeyboardRequest) error {
	SQL := `
		INSERT INTO keyboards(asset_id, connectivity, layout)
		VALUES($1, $2, $3)
	`

	args := []interface{}{
		assetID,
		req.Connectivity,
		req.Layout,
	}

	_, err := database.DB.Exec(SQL, args...)
	return err
}

func InsertLaptopDetails(assetID string, req models.LaptopRequest) error {
	SQL := `
		INSERT INTO laptops(asset_id, processor, ram, storage, operating_system, charger, device_password)
		VALUES($1, $2, $3, $4, $5, $6, $7)
	`

	args := []interface{}{
		assetID,
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
		INSERT INTO mice(asset_id, dpi, connectivity)
		VALUES($1, $2, $3)
	`

	args := []interface{}{
		assetID,
		req.DPI,
		req.Connectivity,
	}

	_, err := database.DB.Exec(SQL, args...)
	return err
}

func InsertMobileDetails(assetID string, req models.MobileRequest) error {
	SQL := `
		INSERT INTO mobiles(asset_id, operating_system, ram, storage, charger, device_password)
		VALUES($1, $2, $3, $4, $5, $6)
	`

	args := []interface{}{
		assetID,
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
		INSERT INTO assets(brand, model, serial_number, asset_type, owner_type, warranty_start, warranty_end)
		VALUES($1, $2, $3, $4, $5, $6, $7)
		RETURNING id
	`

	args := []interface{}{
		model.Brand,
		model.Model,
		model.SerialNumber,
		model.Type,
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

func ShowAssets(brand, model, assetType, serial_number, status, owner string) ([]models.AllAssetsInfoRequest, error) {
	SQL := `SELECT id, brand, model, asset_type, serial_number, status, owner_type, assigned_by_id, assigned_to_id, assigned_at, warranty_start, warranty_end, service_start, service_end, returned_at, created_at,updated_at 
          FROM assets
          WHERE archived_at IS NULL 
          AND (
              $1= '' or brand LIKE'%'||$1||'%'
          )
          AND(
              $2 ='' or model LIKE'%'||$2||'%'
          )
          AND (
              $3='' or asset_type::text LIKE'%'||$3||'%'
          )
          AND(
              $4 ='' or serial_number LIKE'%'||$4||'%'
          )
          AND(
              $5='' or status::text LIKE '%'||$5||'%'
          )
          AND(
              $6=''or owner_type::text LIKE '%'||$6||'%'
          )
          ORDER BY created_at
          `
	assets := make([]models.AllAssetsInfoRequest, 0)

	err := database.DB.Select(&assets, SQL, brand, model, assetType, serial_number, status, owner)
	if err != nil {
		return nil, err
	}
	return assets, nil
}

func AssignedAssets(id, assignedById, assignedTo string) error {
	SQL := `UPDATE assets
          SET assigned_to_id=$3,
			  assigned_by_id=$2,
              assigned_at=NOW(),
              status='assigned',
              updated_at=NOW()
          WHERE id=$1
          AND archived_at IS NULL 
              `
	// _, err := database.DB.Exec(SQL, assignedById, assignedTo, id)
	fmt.Println(id)
	fmt.Println(assignedById)
	fmt.Println(assignedTo)
	res, err := database.DB.Exec(SQL, id, assignedById, assignedTo)
	if err != nil {
		return err
	}
	rows, _ := res.RowsAffected()
	if rows == 0 {
		return errors.New("asset not found or already archived")
	}
	return nil
}
