package service

import "github.com/Sakenzhassulan/it-analytics-test-task/repo"

type Service struct {
	Repo *repo.Repo
}

func New(repo *repo.Repo) *Service {
	return &Service{
		Repo: repo,
	}
}
