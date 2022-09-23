package models

import "time"

type Book struct {
	ID        uint       `gorm:"primary_key;"`
	UserId    int        `json:"user_id"`
	Title     string     `json:"title"`
	Isbn      string     `json:"isbn"`
	Writer    string     `json:"writer"`
	CreatedAt time.Time  `gorm:"autoCreateTime"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime"`
	DeletedAt *time.Time `gorm:"autoDeleteTime"`
}
