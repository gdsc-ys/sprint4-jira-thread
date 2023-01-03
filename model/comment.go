package model

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Content  string `json:"content"`
	UserID   string `json:"userid"`
	ThreadID uint   `json:"threadid"`
}
