package model

import (
	"gorm.io/gorm"
)

type Thread struct {
	gorm.Model
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	UserID   int32     `json:"userid"`
	Comments []Comment `json:"comments"`
}
