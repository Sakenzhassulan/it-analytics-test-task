package repo

import (
	"github.com/Sakenzhassulan/it-analytics-test-task/models"
	"github.com/gin-gonic/gin"
)

func (r *Repo) SaveResults(ctx *gin.Context, results map[int]models.Result) error {
	query := `
	insert into results(
		first_team_id, 
		second_team_id, 
		division_name, 
		first_team_score, 
		second_team_score, 
		stage
	) values($1 ,$2, $3, $4, $5, $6)`

	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	for _, result := range results {
		if _, err := tx.Exec(query,
			result.FirstTeamId,
			result.SecondTeamId,
			result.DivisionName,
			result.FirstTeamScore,
			result.SecondTeamScore,
			result.Stage,
		); err != nil {
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (r *Repo) GetPlayOffResults(ctx *gin.Context, divisionName string) ([]models.Result, error) {
	query := `
	select 
		id, 
		first_team_id, 
		second_team_id, 
		division_name, 
		first_team_score, 
		second_team_score, 
		stage
	from results
	where division_name = $1
	`

	rows, err := r.DB.Query(query, divisionName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []models.Result
	for rows.Next() {
		var result models.Result
		if err := rows.Scan(
			&result.Id,
			&result.FirstTeamId,
			&result.SecondTeamId,
			&result.DivisionName,
			&result.FirstTeamScore,
			&result.SecondTeamScore,
			&result.Stage,
		); err != nil {
			return nil, err
		}

		results = append(results, result)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}
