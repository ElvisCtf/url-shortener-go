package repository

import (
    "shorten-service/internal/model"
)

type Repository interface {
    Save(originalURL string) string
    FindByCode(code string) (*model.Link, error)
}

func NewRepo(storage string) Repository {
    if storage == "postgre" {
        return NewPostgreRepo()
    } else {
        return NewMemoryRepo()
    }
}