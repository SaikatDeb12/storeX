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
