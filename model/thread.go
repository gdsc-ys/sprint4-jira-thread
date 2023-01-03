package model

type Thread struct {
	gorm.model
	Title string `json:"title"`
	Content string `json:"content"`
	UserID string `json:"userid"`
	Comments []Comment  `json:"comment"`
}