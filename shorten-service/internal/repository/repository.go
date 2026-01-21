package repository

import (
    "time"
)

type Link struct {
    Code        string
    OriginalURL string
    CreatedAt 	time.Time
}

type Repository interface {
    Save(originalURL string) string
    FindByCode(code string) (*Link, error)
}