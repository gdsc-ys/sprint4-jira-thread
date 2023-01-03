package controller

import (
	"net/http"

	"github.com/gdsc-ys/sprint4-jira-thread/database"
	"github.com/gdsc-ys/sprint4-jira-thread/model"
	"github.com/gin-gonic/gin"
)

func ReadThread(c *gin.Context) {
	thread := model.Thread{}

	if err := database.DB.First(&thread, c.Param("threadID")).Error; err != nil {
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

	thread.UserID = c.MustGet("userid").(int32)

	if err := database.DB.Create(&thread).Error; err != nil {
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

	if err := database.DB.First(&oldThread, c.Param("threadID")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := database.DB.Model(&oldThread).Updates(thread).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": thread,
	})
}

func DeleteThread(c *gin.Context) {
	thread := model.Thread{}

	if err := database.DB.Delete(&thread, c.Param("threadID")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": thread,
	})
}
