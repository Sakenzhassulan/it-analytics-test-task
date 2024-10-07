package service

import "github.com/gin-gonic/gin"

func (s *Service) DeleteTournament(ctx *gin.Context) (bool, error) {
	ok, err := s.Repo.DeleteTournament(ctx)
	if err != nil {
		return false, err
	}
	return ok, nil
}
