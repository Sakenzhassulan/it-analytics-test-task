package repo

import (
	"log"

	"github.com/Sakenzhassulan/it-analytics-test-task/models"
	"github.com/gin-gonic/gin"
)

func (r *Repo) CreateTeams(ctx *gin.Context, divA, divB []string) ([]models.Team, error) {
	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	query := `insert into teams(name, division_name) values ($1,$2)`
	for i := 0; i < len(divA); i++ {
		if _, err := tx.Exec(query, divA[i], "A"); err != nil {
			return nil, err
		}
	}

	for i := 0; i < len(divB); i++ {
		if _, err := tx.Exec(query, divB[i], "B"); err != nil {
			return nil, err
		}
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	query = `select id, name, division_name from teams`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var teams []models.Team
	for rows.Next() {
		var team models.Team
		err := rows.Scan(&team.Id, &team.Name, &team.DivisionName)
		if err != nil {
			return nil, err
		}
		teams = append(teams, team)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return teams, nil
}

func (r *Repo) GetTeamsByDivisionName(ctx *gin.Context, divisionName string) ([]models.Team, error) {
	query := `select id, name, division_name from teams where division_name = $1`
	rows, err := r.DB.Query(query, divisionName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var teams []models.Team
	for rows.Next() {
		var team models.Team
		err := rows.Scan(&team.Id, &team.Name, &team.DivisionName)
		if err != nil {
			return nil, err
		}

		teams = append(teams, team)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return teams, nil
}

func (r Repo) SaveTeams(ctx *gin.Context, divisionName string, teams map[int]models.Team) ([]models.Team, error) {
	query := `
	update teams set 
		played = $1,
		won = $2,
		drawn = $3,
		lost = $4,
		goals_for = $5,
		against = $6,
		goal_difference = $7,
		points = $8
	where id = $9
	`

	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	for k, team := range teams {
		log.Println(team.Id, team.Played)
		if _, err := tx.Exec(query,
			team.Played,
			team.Won,
			team.Drawn,
			team.Lost,
			team.GoalsFor,
			team.Against,
			team.GoalDiff,
			team.Points,
			k,
		); err != nil {
			return nil, err
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	sortedTeams, err := r.GetSortedTeams(divisionName, 8)
	if err != nil {
		return nil, err
	}
	return sortedTeams, nil
}

func (r Repo) GetSortedTeams(divisionName string, limit int) ([]models.Team, error) {
	query := `
	select 
		id, 
		name, 
		division_name, 
		played, 
		won, 
		drawn, 
		lost, 
		goals_for, 
		against,
		goal_difference, 
		points 
	from teams 
	where 
		division_name = $1 
	order by 
		points desc, 
		goal_difference desc, 
		goals_for desc, 
		name desc
	limit $2
	`

	rows, err := r.DB.Query(query, divisionName, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var teams []models.Team
	for rows.Next() {
		var team models.Team
		if err := rows.Scan(
			&team.Id,
			&team.Name,
			&team.DivisionName,
			&team.Played,
			&team.Won,
			&team.Drawn,
			&team.Lost,
			&team.GoalsFor,
			&team.Against,
			&team.GoalDiff,
			&team.Points); err != nil {
			return nil, err
		}

		teams = append(teams, team)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return teams, nil
}
