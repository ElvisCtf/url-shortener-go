package service

import "redirect-service/internal/repository"

type Redirect struct {
}

func NewRedirect(repo repository.Repository) *Redirect {
	return &Redirect{}
}