package service

import (
	"shorten-service/internal/model"
	"shorten-service/internal/repository"
)

type Shorten struct {
	baseURL string
    repo repository.Repository
}

func NewShorten(baseURL string, repo repository.Repository) *Shorten {
    return &Shorten{baseURL: baseURL, repo: repo}
}

func (s *Shorten) Create(originalURL string) model.ShortenResponse {
	code := s.repo.Save(originalURL)
	resp := model.ShortenResponse {
		OriginalURL: originalURL,
		ShortenURL: s.baseURL + "/" + code,
	}
	return resp
}
