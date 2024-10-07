package repo

import (
	"database/sql"

	"github.com/Sakenzhassulan/it-analytics-test-task/config"
)

type Repo struct {
	DB *sql.DB
}

func New(config *config.Config) *Repo {
	db := connectDB(config)
	return &Repo{
		DB: db,
	}
}
