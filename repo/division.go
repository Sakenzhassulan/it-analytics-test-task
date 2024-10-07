package repo

import (
	"github.com/gin-gonic/gin"
)

func (r *Repo) CreateDivisions(ctx *gin.Context) error {
	divisions := []string{"A", "B"}

	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query := `insert into divisions(name) values ($1)`
	for i := 0; i < 2; i++ {
		if _, err := tx.Exec(query, divisions[i]); err != nil {
			return err
		}
	}

	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}
