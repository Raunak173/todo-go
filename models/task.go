package models

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Heading     string `json:"heading"`
	Description string `json:"description"`
	IsCompleted bool   `json:"is_completed" gorm:"default:false"`
}
