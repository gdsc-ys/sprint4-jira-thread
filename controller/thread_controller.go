package controller

import (
	"net/http"

	"github.com/gdsc-ys/sprint4-jira-thread/model"
	"github.com/gdsc-ys/sprint4-jira-thread/database"
	"github.com/gin-gonic/gin"
)

func ReadThread(c *gin.Context) {
	thread := model.Thread{}

	if err := database.db.First(&thread, c.param("titleID")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": thread,
	})
}

func CreateThread(c *gin.Context) {
	thread := model.Thread{}

	if err := c.ShouldBind(&thread); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := database.db.Create(&thread).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": thread,
	})
}

func UpdateThread(c *gin.Context) {
	thread := model.Thread{}
	oldThread := model.Thread{}
	
	if err := c.ShouldBind(&thread); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := database.db.First(&oldThread, c.param("titleID")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := database.db.Model(&oldThread).Updates(thread).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
}