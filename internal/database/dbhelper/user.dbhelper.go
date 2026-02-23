package dbhelper

import (
	"github.com/SaikatDeb12/storeX/internal/database"
	"github.com/SaikatDeb12/storeX/internal/models"
)

func CheckUserExistsByEmail(email string) error {
	SQL := `
		SELECT COUNT(*)
		FROM users
		WHERE email=TRIM(LOWER($1)) AND archived_at IS NULL
	`

	err := database.DB.Get(SQL, email)
	return err
}

func CreateUser(name, email, phoneNumber, role, employment, hashedPassword string) (string, error) {
	SQL := `
		INSERT INTO users(name, email, phone_number, role, type, password)
		VALUES($1, TRIM(LOWER($2)), $3, $4, $5, $6, $7 )
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
		INSERT INTO session(user_id)
		VALUES($1)
		RETURNING id
	`
	var sessionID string
	err := database.DB.Get(&sessionID, SQL)
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
	err := database.DB.Get(user, SQL)
	if err != nil {
		return user, err
	}
	return user, nil
}
