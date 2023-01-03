package controller

import (
	"net/http"

	"github.com/gdsc-ys/sprint4-jira-thread/database"
	"github.com/gdsc-ys/sprint4-jira-thread/model"
	"github.com/gin-gonic/gin"
)

func ReadAllThreadComment(c *gin.Context) {
	comments := []model.Comment{}

	if err := database.DB.Where("thread_id = ?", c.Param("threadID")).Find(&comments).Error; err != nil {
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

	comment.UserID = c.MustGet("userid").(int32)

	if err := database.DB.Create(&comment).Error; err != nil {
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

	if err := database.DB.First(&oldComment, c.Param("commentID")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := database.DB.Model(&oldComment).Updates(comment).Error; err != nil {
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

	if err := database.DB.Delete(&comment, c.Param("commentID")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": comment,
	})
}
