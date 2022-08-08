package controllers

import (
	"net/http"
	"routine_todo/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Get all tasks
func FindTasks(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var tasks []models.Tasks
	db.Find(&tasks)

	c.JSON(http.StatusOK, gin.H{"data": tasks})
}
