package controller

import (
	"net/http"

	"github.com/gdsc-ys/sprint4-jira-thread/model"
	"github.com/gin-gonic/gin"
)

func ReadAllThreadComment(c *gin.Context) {
	comments := []model.Comment{}

	if err := database.db.Find(&comments, c.param("threadID")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": comments,
	})
}

func CreateThreadComment(c *gin.Context) {
	comment := model.Comment{}

	if err := c.ShouldBind(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := database.db.Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": comment,
	})
}

func UpdateThreadComment(c *gin.Context) {
	comment := model.Comment{}
	oldComment := model.Comment{}

	if err := c.ShouldBind(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := database.db.First(&oldComment, c.param("commentID")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := database.db.Model(&oldComment).Updates(comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": comment,
	})
}

func DeleteThreadComment(c *gin.Context) {
	comment := model.Comment{}

	if err := database.db.Delete(&comment, c.param("commentID")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": comment,
	})
}
