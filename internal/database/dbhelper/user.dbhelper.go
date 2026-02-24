package dbhelper

import (
	"fmt"

	"github.com/SaikatDeb12/storeX/internal/database"
	"github.com/SaikatDeb12/storeX/internal/models"
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

func CreateUser(name, email, phoneNumber, role, employment, hashedPassword string) (string, error) {
	SQL := `
		INSERT INTO users(name, email, phone_number, role, employment, password)
		VALUES($1, TRIM(LOWER($2)), $3, $4, $5, $6)
		RETURNING id
	`
	var userID string
	err := database.DB.Get(&userID, SQL, name, email, phoneNumber, role, employment, hashedPassword)
	if err != nil {
		return "", err
	}
	return string(userID), nil
}

func CreateSession(userID string) (string, error) {
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
		SELECT id, email, password, archived_at 
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

// should show:
// user details
// .
// .
// .
// asset count, assignedStatus,
// func GetAssetCount(userID string) (int, error) {
// 	SQL := `
// 		SELECT COUNT(*) FROM assets
// 		WHERE assigned_to_id=$2
// 	`
// 	var count int
// 	err := database.DB.Get(&count, SQL, userID)
// 	return count, err
// }

func GetUserInfo() ([]models.UserInfoRequest, error) {
	SQL := `
		SELECT id, name, email, phone_number, role, employment, created_at
		FROM users
	`
	users := make([]models.UserInfoRequest, 0)
	err := database.DB.Select(&users, SQL)
	if err != nil {
		return users, err
	}

	fmt.Println(users)
	return users, err
}
