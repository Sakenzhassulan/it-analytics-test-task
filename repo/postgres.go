package repo

import (
	"database/sql"
	"fmt"

	"github.com/Sakenzhassulan/it-analytics-test-task/config"
	_ "github.com/lib/pq"
)

func connectDB(config *config.Config) *sql.DB {
	conn := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable",
		config.DBHost,
		config.DBPort,
		config.DBName,
		config.DBUser,
		config.DBPassword,
	)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err)
	}

	err = runMigrations(db, config)
	if err != nil {
		panic(err)
	}
	return db
}
