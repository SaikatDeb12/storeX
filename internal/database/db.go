package database

import (
	"errors"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

var DB *sqlx.DB

type SSLMODETYPE string

const sslmode SSLMODETYPE = "disable"

func migrateUp(db *sqlx.DB) error {
	driver, driverErr := postgres.WithInstance(db.DB, &postgres.Config{})
	if driverErr != nil {
		return driverErr
	}

	migrateIns, migrateErr := migrate.NewWithDatabaseInstance("file://internal/database/migrations", "postgres", driver)
	if migrateErr != nil {
		return migrateErr
	}

	if err := migrateIns.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}

func Connect() error {
}
