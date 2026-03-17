package dbhelper

import (
	"errors"
	"time"

	"github.com/SaikatDeb12/storeX/internal/database"
	"github.com/SaikatDeb12/storeX/internal/models"
	"github.com/jmoiron/sqlx"
)

func InsertKeyboardDetails(tx *sqlx.Tx, assetID string, req models.KeyboardRequest) error {
	SQL := `
		INSERT INTO keyboards(asset_id, connectivity, layout)
		VALUES($1, $2, $3)
	`

	args := []interface{}{
		assetID,
		req.Connectivity,
		req.Layout,
	}

	_, err := tx.Exec(SQL, args...)
	return err
}

func InsertLaptopDetails(tx *sqlx.Tx, assetID string, req models.LaptopRequest) error {
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

	_, err := tx.Exec(SQL, args...)
	return err
}

func InsertMouseDetails(tx *sqlx.Tx, assetID string, req models.MouseRequest) error {
	SQL := `
		INSERT INTO mice(asset_id, dpi, connectivity)
		VALUES($1, $2, $3)
	`

	args := []interface{}{
		assetID,
		req.DPI,
		req.Connectivity,
	}

	_, err := tx.Exec(SQL, args...)
	return err
}

func InsertMobileDetails(tx *sqlx.Tx, assetID string, req models.MobileRequest) error {
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

	_, err := tx.Exec(SQL, args...)
	return err
}

func CreateAsset(tx *sqlx.Tx, model models.CreateAssetRequest) (string, error) {
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
	err := tx.Get(&assetID, SQL, args...)
	if err != nil {
		return assetID, err
	}
	return assetID, nil
}

func FetchAssets(brand, model, assetType, serial_number, status, owner string, limit, offset int) ([]models.AllAssetsInfoRequest, error) {
	SQL := `SELECT id, brand, model, asset_type, serial_number, status, owner_type, assigned_by_id, assigned_to_id, assigned_at, warranty_start, warranty_end, service_start, service_end, returned_at, created_at, updated_at 
          FROM assets
          WHERE archived_at IS NULL 
          AND (
              $1= '' or brand ILIKE '%'||$1||'%'
          )
          AND(
              $2 ='' or model ILIKE '%'||$2||'%'
          )
          AND (
              $3='' or asset_type::text ILIKE '%'||$3||'%'
          )
          AND(
              $4 ='' or serial_number ILIKE '%'||$4||'%'
          )
          AND(
              $5='' or status::text ILIKE '%'||$5||'%'
          )
          AND(
              $6=''or owner_type::text ILIKE '%'||$6||'%'
          )
          ORDER BY created_at
		  LIMIT $7 OFFSET $8
    `
	var result []models.AllAssetsInfoRequest
	err := database.DB.Select(&result, SQL, brand, model, assetType, serial_number, status, owner, limit, offset)
	return result, err
}

func FetchAssetsInfo(userID, assetStatus string) ([]models.AssetInfoRequest, error) {
	SQL := `
		SELECT id, brand, model, status, asset_type
		FROM assets
		WHERE assigned_to_id=$1
		AND ($2 = '' OR status::TEXT=$2)
	`
	assetDetails := make([]models.AssetInfoRequest, 0)
	err := database.DB.Select(&assetDetails, SQL, userID, assetStatus)
	return assetDetails, err
}

func GettingAssetsCount() (models.DashboardSummaryRequest, error) {
	SQL := `
		SELECT 
		COUNT(*) AS total,
		COUNT(*) FILTER (WHERE status='available' ) AS available,
		COUNT(*) FILTER (WHERE status='assigned' ) AS assigned,
		COUNT(*) FILTER (WHERE status='under_repair' ) AS waitingForRepair,
		COUNT(*) FILTER (WHERE status='in_service' ) AS inService,
		COUNT(*) FILTER (WHERE status='damaged' ) AS damaged
		FROM assets
		WHERE archived_at IS NULL
	`
	var result models.DashboardSummaryRequest
	err := database.DB.Get(&result, SQL)
	return result, err
}

func AssignAssets(id, assignedById, assignedTo string) error {
	SQL := `UPDATE assets
          SET assigned_to_id=$3,
			  assigned_by_id=$2,
              assigned_at=NOW(),
              status='assigned',
              updated_at=NOW()
          WHERE id=$1 AND status='available'
          AND archived_at IS NULL 
	`
	// _, err := database.DB.Exec(SQL, assignedById, assignedTo, id)
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

func UpdateAsset(
	tx *sqlx.Tx,
	assetID string,
	brand, model, serialNo, assetType, status, owner *string,
	warrantyStart, warrantyEnd *time.Time,
) error {
	SQL := `
	UPDATE assets
	SET
		brand = COALESCE($2, brand),
		model = COALESCE($3, model),
		serial_number = COALESCE($4, serial_number),
		asset_type = COALESCE($5, asset_type),
		status = COALESCE($6, status),
		owner_type = COALESCE($7, owner_type),
		warranty_start = COALESCE($8, warranty_start),
		warranty_end = COALESCE($9, warranty_end),
		updated_at = now()
	WHERE id = $1
	AND archived_at IS NULL
	`

	result, err := tx.Exec(SQL,
		assetID,
		brand,
		model,
		serialNo,
		assetType,
		status,
		owner,
		warrantyStart,
		warrantyEnd,
	)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("asset not found or archived")
	}

	return nil
}

func UpdateLaptop(tx *sqlx.Tx, assetID string, laptop *models.LaptopRequest) error {
	SQL := `
		UPDATE laptops
		SET
			processor = $2,
			ram = $3,
			storage = $4,
			operating_system = $5,
			charger = $6,
			device_password = $7
		WHERE asset_id = $1
    `

	_, err := tx.Exec(
		SQL,
		assetID,
		laptop.Processor,
		laptop.RAM,
		laptop.Storage,
		laptop.OperatingSystem,
		laptop.Charger,
		laptop.DevicePassword,
	)

	return err
}

func UpdateMouse(tx *sqlx.Tx, assetID string, mouse *models.MouseRequest) error {
	SQL := `
		UPDATE mice
		SET
			dpi = $2,
			connectivity = $3
		WHERE asset_id = $1
    `

	_, err := tx.Exec(SQL, assetID, mouse.DPI, mouse.Connectivity)
	return err
}

func UpdateKeyboard(tx *sqlx.Tx, assetID string, keyboard *models.KeyboardRequest) error {
	SQL := `
		UPDATE keyboards
		SET
			layout = $2,
			connectivity = $3
		WHERE asset_id = $1
    `

	_, err := tx.Exec(SQL, assetID, keyboard.Layout, keyboard.Connectivity)
	return err
}

func UpdateMobile(tx *sqlx.Tx, assetID string, mobile *models.MobileRequest) error {
	SQL := `
		UPDATE mobiles
		SET
			operating_system = $2,
			ram = $3,
			storage = $4,
			charger = $5,
			device_password = $6
		WHERE asset_id = $1
    `

	_, err := tx.Exec(
		SQL,
		assetID,
		mobile.OperatingSystem,
		mobile.RAM,
		mobile.Storage,
		mobile.Charger,
		mobile.DevicePassword,
	)

	return err
}

func SentToService(assetId string, serviceStart, serviceEnd time.Time) error {
	SQL := `UPDATE assets
			SET 
				status='in_service', 
				service_start=$2, 
				service_end=$3, 
				updated_at=now()
			WHERE id=$1 
			AND archived_at IS NULL 
			AND status='available'
	`
	result, err := database.DB.Exec(SQL, assetId, serviceStart, serviceEnd)
	if err != nil {
		return err
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("asset not found or already archived")
	}
	return nil
}

func UnassignAssets(tx *sqlx.Tx, userID string) error {
	SQL := `
		UPDATE assets
		SET 
			assigned_to_id = NULL,
			assigned_by_id = NULL,
			assigned_at = NULL,
			status = 'available',
			returned_at = now(),
			updated_at = NOW()
		WHERE assigned_to_id = $1
      	AND archived_at IS NULL
	`
	_, err := tx.Exec(SQL, userID)
	return err
}
