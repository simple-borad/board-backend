package models

import "time"

type Board struct {
	ID	uint `gorm:"primaryKey"`
	Title	string `gorm:"size:255"`
	Content	string `gorm:"size:255"`
	Author	string `gorm:"size:255"`
	CreatedAt	time.Time	`gorm:"autoCreateTime"`
	UpdateAt	time.Time	`gorm:"autoUpdateTime"`
}