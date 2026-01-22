package repository

import "errors"

var notFoundErr = errors.New("not found")

type Repository interface {
	
}

func NewRepo() Repository {
    return newPostgreRepo()
}