package service

import (
	"errors"
	"math/rand"

	"github.com/Sakenzhassulan/it-analytics-test-task/models"
	"github.com/gin-gonic/gin"
)

func (s *Service) CreateTeams(ctx *gin.Context, list []string) ([]models.Team, error) {
	if len(list) != 16 {
		return nil, errors.New("length of teams must be 16")
	}

	if err := s.Repo.CreateDivisions(ctx); err != nil {
		return nil, err
	}

	rand.Shuffle(len(list), func(i, j int) {
		list[i], list[j] = list[j], list[i]
	})
	divisionATeams := list[:8]
	divisionBTeams := list[8:]

	teams, err := s.Repo.CreateTeams(ctx, divisionATeams, divisionBTeams)
	if err != nil {
		return nil, err
	}
	return teams, nil
}
