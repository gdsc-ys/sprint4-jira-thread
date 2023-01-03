package model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/gdsc-ys/sprint4-jira-thread/model"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("jira-thread.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(&model.Thread{}, &model.Comment{})
}
