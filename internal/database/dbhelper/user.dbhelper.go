package dbhelper

import "github.com/SaikatDeb12/storeX/internal/database"

func CheckUserExistsByEmail(email string) error {
	SQL := `
		SELECT COUNT(*)
		FROM users
		WHERE email=TRIM(LOWER($1)) AND archived_at IS NULL
	`

	err := database.DB.Get(SQL, email)
	return err
}

func CreateUser(name, email, phoneNumber, userRole, userType, hashedPassword string) (string, error) {
	SQL := `
		INSERT INTO users(name, email, phone_number, role, type, password)
		VALUES($1, $2, $3, $4, $5, $6, $7 )
		RETURNING id
	`
	var userID string
	err := database.DB.Get(&userID, SQL, name, email, phoneNumber, userRole, userType, hashedPassword)
	if err != nil {
		return "", err
	}
	return string(userID), nil
}
