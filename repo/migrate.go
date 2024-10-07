package repo

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Sakenzhassulan/it-analytics-test-task/config"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func runMigrations(db *sql.DB, config *config.Config) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}

	migrator, err := migrate.NewWithDatabaseInstance(fmt.Sprintf("file://%s", "migrations"), config.DBName, driver)
	if err != nil {
		return err
	}

	// if err := migrator.Steps(-3); err != nil {
	// 	return err
	// }

	err = migrator.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}
