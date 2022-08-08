package controllers

import (
	"net/http"
	"routine_todo/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CreateTaskInput struct {
	Task     string `json:"task_name"`
	Assignee string `json:"assignee"`
	Deadline string `json:"deadline"`
}

// Get all tasks
func FindTasks(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var tasks []models.Tasks
	db.Find(&tasks)

	c.JSON(http.StatusOK, gin.H{"data": tasks})
}

// Create new task
func CreateTask(c *gin.Context) {
	// Validate input
	var input CreateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		// TODO: validation input null
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!!"})
		return
	}

	date := "2006-01-02"
	deadline, _ := time.Parse(date, input.Deadline)

	// Create task
	task := models.Tasks{Assignee: input.Assignee, Task_Name: input.Task, Deadline: deadline}

	db := c.MustGet("db").(*gorm.DB)
	db.Create(&task)

	c.JSON(http.StatusOK, gin.H{"data": task})
}
