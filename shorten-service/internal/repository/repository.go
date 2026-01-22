package repository

import (
    "errors"

    "shorten-service/internal/model"
)

var saveErr = errors.New("Cannot save")
var notFoundErr = errors.New("not found")

type Repository interface {
    Save(originalURL string) (string, error)
    FindByCode(code string) (*model.Link, error)
}

func NewRepo(storage string) Repository {
    if storage == "postgre" {
        return NewPostgreRepo()
    } else {
        return NewMemoryRepo()
    }
}