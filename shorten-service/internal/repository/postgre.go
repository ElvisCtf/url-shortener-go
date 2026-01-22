package repository

import (
	"errors"
	"fmt"
	"log"

	"shorten-service/internal/model"
	"shorten-service/internal/util"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgreRepo struct {
    db *gorm.DB
}

func NewPostgreRepo() *PostgreRepo {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		util.Env("DB_HOST", "localhost"),
		util.Env("DB_USER", "postgres16"),
		util.Env("DB_PASSWORD", "postgres16"),
		util.Env("DB_NAME", "urlshortener"),
		util.Env("DB_PORT", "5432"),
		util.Env("DB_SSLMODE", "disable"),
		util.Env("TZ", "Asia/Hong_Kong"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil { 
		log.Fatal(err)
		return nil
	}

	if migrateErr := db.AutoMigrate(&model.Link{}); migrateErr != nil { 
		log.Fatal(err)
	}

	return &PostgreRepo{db: db}
}


func (repo *PostgreRepo) Save(originalURL string) (string, error) {
	link := &model.Link{
		OriginalURL: originalURL,
	}
	if err := repo.db.Create(link).Error; err != nil {
		log.Fatal(err)
		return "", saveErr
	}
	return link.Code, nil
}


func (repo *PostgreRepo) FindByCode(code string) (*model.Link, error) {
    var link model.Link
    err := repo.db.
        Where("code = ?", code).
        First(&link).Error

    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, nil
        }
        return nil, err
    }
    return &link, nil
}