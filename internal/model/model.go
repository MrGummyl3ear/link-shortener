package model

import (
	"errors"
)

var (
	ErrInvalidURL     = errors.New("invalid url")
	ErrNotFound       = errors.New("url is not found")
	ErrInvalidStorage = errors.New("invalid storage")
)

type Shortening struct {
	OriginalURL string `json:"original_url" gorm:"not null"`
	ShortUrl    string `json:"short_url" gorm:"primaryKey unique not null"`
}
