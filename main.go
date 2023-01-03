package main

import (
	"github.com/gin-gonic/gin"

	"github.com/gdsc-ys/sprint4-jira-thread/controller"
	"github.com/gdsc-ys/sprint4-jira-thread/database"
)

func main() {
	router := gin.Default()

	database.ConnectDatabase()

	thread := router.Group("/thread")
	{
		thread.GET("/:threadID", controller.ReadThread)
		thread.POST("/", controller.CreateThread)
		thread.PUT("/:threadID", controller.UpdateThread)
		thread.DELETE("/:threadID", controller.DeleteThread)
	}

	comment := router.Group("/comment")
	{
		comment.GET("/:threadID", controller.ReadAllThreadComment)
		comment.POST("/:threadID", controller.CreateThreadComment)
		comment.PUT("/:threadID/:commentID", controller.UpdateThreadComment)
		comment.DELETE("/:threadID/:commentID", controller.DeleteThreadComment)
	}

	router.Run(":8080")
}
