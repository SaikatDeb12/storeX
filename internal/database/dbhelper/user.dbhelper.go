package dbhelper

import (
	"errors"

	"github.com/SaikatDeb12/storeX/internal/database"
	"github.com/SaikatDeb12/storeX/internal/models"
	"github.com/jmoiron/sqlx"
)

func CheckUserExistsByEmail(email string) (bool, error) {
	SQL := `
		SELECT COUNT(*)
		FROM users
		WHERE email=TRIM(LOWER($1)) AND archived_at IS NULL
	`
	var count int
	err := database.DB.Get(&count, SQL, email)
	if err != nil {
		return false, err
	}

	return count > 0, err
}

func CreateUser(tx *sqlx.Tx, name, email, phoneNumber, employment, hashedPassword string) (string, error) {
	SQL := `
		INSERT INTO users(name, email, phone_number, employment, password)
		VALUES($1, TRIM(LOWER($2)), $3, $4, $5)
		RETURNING id
	`
	var userID string
	err := tx.Get(&userID, SQL, name, email, phoneNumber, employment, hashedPassword)
	if err != nil {
		return "", err
	}
	return string(userID), nil
}

func CreateSessionOnRegister(tx *sqlx.Tx, userID string) (string, error) {
	SQL := `
		INSERT INTO user_sessions(user_id)
		VALUES($1)
		RETURNING id
	`
	var sessionID string
	err := tx.Get(&sessionID, SQL, userID)
	if err != nil {
		return "", err
	}
	return string(sessionID), nil
}

func CreateSessionOnLogin(userID string) (string, error) {
	SQL := `
		INSERT INTO user_sessions(user_id)
		VALUES($1)
		RETURNING id
	`
	var sessionID string
	err := database.DB.Get(&sessionID, SQL, userID)
	if err != nil {
		return "", err
	}
	return string(sessionID), nil
}

func GetUserAuthByEmail(email string) (models.User, error) {
	SQL := `
		SELECT id, email, password, role
		FROM users 
		WHERE email=TRIM(LOWER($1)) AND archived_at IS NULL
	`
	var user models.User
	err := database.DB.Get(&user, SQL, email)
	if err != nil {
		return user, err
	}
	return user, nil
}

func FetchUserRole(tx *sqlx.Tx, userID string) (string, error) {
	SQL := `
		SELECT role 
		FROM users
		WHERE id=$1 AND archived_at IS NULL
	`

	var role string
	err := tx.Get(&role, SQL, userID)
	return role, err
}

func AssignUserRole(userID, role string) error {
	SQL := `
		UPDATE users
		SET role=$2
		WHERE id=$1 AND archived_at IS NULL
	`
	result, err := database.DB.Exec(SQL, userID, role)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("user not found")
	}
	return nil
}

func FetchAssetInfo(userID string) ([]models.AssetInfoRequest, error) {
	SQL := `
	SELECT id, brand, model, status, asset_type
	FROM assets
	WHERE assigned_to_id = $1
	AND archived_at IS NULL
	`
	assets := []models.AssetInfoRequest{}
	err := database.DB.Select(&assets, SQL, userID)
	return assets, err
}

func FetchUsers(name, role, employment, assetStatus string, limit, offset int) ([]models.UserInfoRequest, error) {
	SQL := `
	SELECT 
		u.id,
		u.name,
		u.email,
		u.phone_number,
		u.role,
		u.employment,
		u.created_at,
		a.id AS asset_id,
		a.brand,
		a.model,
		a.status,
		a.asset_type
	FROM users u
	LEFT JOIN assets a
		ON a.assigned_to_id = u.id
		AND a.archived_at IS NULL
	WHERE 
		($1 = '' OR u.name ILIKE '%' || $1 || '%')
		AND ($2 = '' OR u.role::TEXT = $2)
		AND ($3 = '' OR u.employment::TEXT = $3)
		AND ($4 = '' OR a.status::TEXT = $4)
		AND u.archived_at IS NULL
	`
	userAssetRows := make([]models.UserAssetRow, 0)
	err := database.DB.Select(&userAssetRows, SQL, name, role, employment, assetStatus)
	if err != nil {
		return nil, err
	}

	userMap := make(map[string]*models.UserInfoRequest)

	for _, userAssetInfo := range userAssetRows {
		if _, ok := userMap[userAssetInfo.ID]; !ok {
			userMap[userAssetInfo.ID] = &models.UserInfoRequest{
				ID:           userAssetInfo.ID,
				Name:         userAssetInfo.Name,
				Email:        userAssetInfo.Email,
				PhoneNumber:  userAssetInfo.PhoneNumber,
				Role:         userAssetInfo.Role,
				Employment:   userAssetInfo.Employment,
				CreatedAt:    userAssetInfo.CreatedAt,
				AssetDetails: []models.AssetInfoRequest{},
			}
		}
		if userAssetInfo.AssetID != nil {
			asset := models.AssetInfoRequest{
				ID:     *userAssetInfo.AssetID,
				Brand:  *userAssetInfo.Brand,
				Model:  *userAssetInfo.Model,
				Status: *userAssetInfo.Status,
				Type:   *userAssetInfo.AssetType,
			}
			userMap[userAssetInfo.ID].AssetDetails = append(userMap[userAssetInfo.ID].AssetDetails, asset)
		}
	}

	users := make([]models.UserInfoRequest, 0, len(userMap))
	for _, user := range userMap {
		users = append(users, *user)
	}
	return users, nil
}

// to make change on the original slice:
// to either use join
// or to use userID slice and assetID slice and then traverse and then put in the map
// postgres any operator

// change in the copy not the original
// for _, user := range users {
// 	userDetails, err := GetAssetInfo(user.ID)
// 	if err != nil {
// 		return users, err
// 	}
// 	fmt.Println(userDetails)
// 	user.AssetDetails = userDetails
// }
// }

func FetchUserByID(userID string) (models.UserInfoRequest, error) {
	SQL := `
		SELECT id, name, email, phone_number, role, employment, created_at
		FROM users
		WHERE archived_at IS NULL AND id=$1
	`
	var user models.UserInfoRequest
	err := database.DB.Get(&user, SQL, userID)
	if err != nil {
		return user, err
	}

	assets, err := FetchAssetInfo(userID)
	if err != nil {
		return user, err
	}
	if len(assets) == 0 {
		return user, nil
	}
	user.AssetDetails = assets
	return user, err
}

func ValidateUserSession(sessionID string) (bool, error) {
	SQL := `
		SELECT COUNT(*) 
		FROM user_sessions
		WHERE id=$1 AND archived_at IS NULL
	`
	var count int
	err := database.DB.Get(&count, SQL, sessionID)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func UpdateUserSession(sessionID string) error {
	SQL := `
		UPDATE user_sessions
		SET archived_at=NOW()
		WHERE id=$1 AND archived_at IS NULL
	`
	result, err := database.DB.Exec(SQL, sessionID)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("session not found")
	}
	return nil
}

func DeleteUser(tx *sqlx.Tx, userID string) error {
	SQL := `
		UPDATE users
		SET archived_at=NOW()
		WHERE id=$1 AND archived_at IS NULL
	`
	result, err := tx.Exec(SQL, userID)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("session not found")
	}
	return nil
}

func DeleteUserSession(tx *sqlx.Tx, userID string) error {
	SQL := `
		UPDATE user_sessions
		SET archived_at=NOW()
		WHERE user_id=$1 AND  archived_at IS NULL
	`

	result, err := tx.Exec(SQL, userID)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("session not found")
	}
	return nil
}
