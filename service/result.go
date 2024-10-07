package service

import (
	"errors"
	"math/rand"

	"github.com/Sakenzhassulan/it-analytics-test-task/models"
	"github.com/gin-gonic/gin"
)

func (s *Service) GenerateResults(ctx *gin.Context, divisionName string) ([]models.Team, error) {
	if divisionName != "A" && divisionName != "B" {
		return nil, errors.New("incorrect division name")
	}

	divTeams, err := s.Repo.GetTeamsByDivisionName(ctx, divisionName)
	if err != nil {
		return nil, err
	}

	var id int
	results := make(map[int]models.Result)

	for i := 0; i < len(divTeams); i++ {
		for j := i + 1; j < len(divTeams); j++ {
			first := rand.Intn(10)
			second := rand.Intn(10)

			result := newResult(first, second, divTeams[i].Id, divTeams[j].Id, divisionName, "GR")
			results[id] = result

			id++
		}
	}

	if err := s.Repo.SaveResults(ctx, results); err != nil {
		return nil, err
	}

	teams := make(map[int]models.Team)
	for i := 0; i < len(results); i++ {
		result := results[i] // results хранит в себе данные всех игр

		team1 := teams[result.FirstTeamId]  // команда 1
		team2 := teams[result.SecondTeamId] // команда 2

		if result.FirstTeamScore > result.SecondTeamScore {
			// Если первая команда победила
			team1.Played += 1
			team1.Won += 1
			team1.GoalsFor += result.FirstTeamScore
			team1.Against += result.SecondTeamScore
			team1.GoalDiff = team1.GoalsFor - team1.Against
			team1.Points += 3

			team2.Played += 1
			team2.Lost += 1
			team2.GoalsFor += result.SecondTeamScore
			team2.Against += result.FirstTeamScore
			team2.GoalDiff = team2.GoalsFor - team2.Against

			teams[result.FirstTeamId] = team1
			teams[result.SecondTeamId] = team2
		} else if result.FirstTeamScore < result.SecondTeamScore {
			// победившая команда - 2
			team1.Played += 1
			team1.Lost += 1
			team1.GoalsFor += result.FirstTeamScore
			team1.Against += result.SecondTeamScore
			team1.GoalDiff = team1.GoalsFor - team1.Against

			team2.Played += 1
			team2.Won += 1
			team2.GoalsFor += result.SecondTeamScore
			team2.Against += result.FirstTeamScore
			team2.GoalDiff = team2.GoalsFor - team2.Against
			team2.Points += 3

			teams[result.FirstTeamId] = team1
			teams[result.SecondTeamId] = team2
		} else {
			// ничья
			team1.Played += 1
			team1.Drawn += 1
			team1.GoalsFor += result.FirstTeamScore
			team1.Against += result.SecondTeamScore
			team1.GoalDiff = team1.GoalsFor - team1.Against
			team1.Points += 1

			team2.Played += 1
			team2.Drawn += 1
			team2.GoalsFor += result.SecondTeamScore
			team2.Against += result.FirstTeamScore
			team2.GoalDiff = team2.GoalsFor - team2.Against
			team2.Points += 3

			teams[result.FirstTeamId] = team1
			teams[result.SecondTeamId] = team2
		}
	}

	sortedTeams, err := s.Repo.SaveTeams(ctx, divisionName, teams)
	if err != nil {
		return nil, err
	}
	return sortedTeams, nil
}

func (s *Service) GeneratePlayOffResults(ctx *gin.Context) ([]models.Result, error) {
	bestTeamsA, err := s.Repo.GetSortedTeams("A", 4)
	if err != nil {
		return nil, err
	}

	bestTeamsB, err := s.Repo.GetSortedTeams("B", 4)
	if err != nil {
		return nil, err
	}

	results := make(map[int]models.Result)
	var semiFinalistsIds []int
	var id int

	// QF (quarter final 1/4) results generated
	for i := 0; i < len(bestTeamsA); i++ {
		first, second := noDrawResults()

		if first > second {
			semiFinalistsIds = append(semiFinalistsIds, bestTeamsA[i].Id)
		} else {
			semiFinalistsIds = append(semiFinalistsIds, bestTeamsB[len(bestTeamsB)-i-1].Id)
		}

		result := newResult(first, second, bestTeamsA[i].Id, bestTeamsB[len(bestTeamsB)-i-1].Id, "PLAY-OFF", "QF")
		results[id] = result

		id++
	}

	var finalistsIds []int

	// SF (semi final 1/2) results generated
	for i := 0; i < 2; i++ {
		first, second := noDrawResults()
		if first > second {
			finalistsIds = append(finalistsIds, semiFinalistsIds[i])
		} else {
			finalistsIds = append(finalistsIds, semiFinalistsIds[i+2])
		}

		result := newResult(first, second, semiFinalistsIds[i], semiFinalistsIds[i+2], "PLAY-OFF", "SF")
		results[id] = result

		id++
	}

	// F (final)
	for i := 0; i < 1; i++ {
		first, second := noDrawResults()

		result := newResult(first, second, finalistsIds[i], finalistsIds[i+1], "PLAY-OFF", "F")
		results[id] = result
	}

	err = s.Repo.SaveResults(ctx, results)
	if err != nil {
		return nil, err
	}

	playOffResults, err := s.Repo.GetPlayOffResults(ctx, "PLAY-OFF")
	if err != nil {
		return nil, err
	}

	return playOffResults, nil
}

func newResult(first, second, firstTeamId, secondTeamId int, divisionName, stage string) models.Result {
	var result models.Result
	result.FirstTeamScore = first
	result.SecondTeamScore = second
	result.DivisionName = divisionName
	result.FirstTeamId = firstTeamId
	result.SecondTeamId = secondTeamId
	result.Stage = stage
	return result
}

func noDrawResults() (int, int) {
	first := rand.Intn(10)
	second := rand.Intn(10)

	if first == second {
		return noDrawResults()
	}
	return first, second
}
