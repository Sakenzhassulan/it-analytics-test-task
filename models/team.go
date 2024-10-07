package models

type TeamsInput struct {
	Teams []string `json:"teams"`
}

type Team struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	DivisionName string `json:"division_name"`
	Played       int    `json:"played"`
	Won          int    `json:"won"`
	Drawn        int    `json:"drawn"`
	Lost         int    `json:"lost"`
	GoalsFor     int    `json:"goals_for"`
	Against      int    `json:"against"`
	GoalDiff     int    `json:"goal_difference"`
	Points       int    `json:"points"`
}
