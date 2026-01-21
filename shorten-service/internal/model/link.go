package model

import "time"

type Link struct {
    Code        string    `gorm:"primaryKey;size:16"`
    OriginalURL string    `gorm:"not null"`
    CreatedAt   time.Time `gorm:"autoCreateTime"`
}