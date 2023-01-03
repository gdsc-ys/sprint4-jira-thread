package model

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Content  string `json:"content"`
	UserID   int32  `json:"userid"`
	ThreadID uint   `json:"threadid"`
}
