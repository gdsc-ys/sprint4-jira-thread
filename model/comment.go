package model

type Comment struct {
	gorm.model
	Content string `json:"content"`
	UserID  string `json:"userid"`
}
