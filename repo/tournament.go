package repo

import "github.com/gin-gonic/gin"

func (r *Repo) DeleteTournament(ctx *gin.Context) (bool, error) {
	query := `delete from results`
	query2 := `delete from teams`
	query3 := `delete from divisions`

	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return false, err
	}
	defer tx.Rollback()

	if _, err := tx.Exec(query); err != nil {
		return false, err
	}

	if _, err := tx.Exec(query2); err != nil {
		return false, err
	}

	if _, err := tx.Exec(query3); err != nil {
		return false, err
	}

	if err := tx.Commit(); err != nil {
		return false, err
	}
	return true, nil
}
