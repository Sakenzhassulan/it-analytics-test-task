package models

type Result struct {
	Id              int    `json:"id"`
	FirstTeamId     int    `json:"first_team_id"`
	SecondTeamId    int    `json:"second_team_id"`
	DivisionName    string `json:"division_name"`
	FirstTeamScore  int    `json:"first_team_score"`
	SecondTeamScore int    `json:"second_team_score"`
	Stage           string `json:"stage"`
}
