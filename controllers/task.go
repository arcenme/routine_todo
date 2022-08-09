package controllers

import (
	"errors"
	"net/http"
	"routine_todo/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
)

type CreateTaskInput struct {
	Task     string `json:"task_name" binding:"required"`
	Assignee string `json:"assignee" binding:"required"`
	Deadline string `json:"deadline" binding:"required"`
}

type UpdateTaskInput struct {
	Task     string `json:"task_name" binding:"required"`
	Assignee string `json:"assignee" binding:"required"`
	Deadline string `json:"deadline" binding:"required"`
}

type UpdateTaskStatus struct {
	Status string `json:"status" binding:"required"`
}

// Validation
type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "lte":
		return "Should be less than " + fe.Param()
	case "gte":
		return "Should be greater than " + fe.Param()
	}
	return "Unknown error"
}

// Get all tasks
func FindTasks(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var tasks []models.Tasks
	db.Find(&tasks)

	c.JSON(http.StatusOK, gin.H{"data": tasks})
}

// Find a task
func FindTask(c *gin.Context) {
	// Get model if exist
	var task models.Tasks
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("task_id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": task})
}

// Create new task
func CreateTask(c *gin.Context) {
	// Validate input
	var input CreateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = ErrorMsg{fe.Field(), getErrorMsg(fe)}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
		}
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

// Update a task
func UpdateTask(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var task models.Tasks
	if err := db.Where("task_id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = ErrorMsg{fe.Field(), getErrorMsg(fe)}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
		}
		return
	}

	date := "2006-01-02"
	deadline, _ := time.Parse(date, input.Deadline)

	var updatedInput models.Tasks
	updatedInput.Deadline = deadline
	updatedInput.Assignee = input.Assignee
	updatedInput.Task_Name = input.Task

	db.Model(&task).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": task})
}

// Delete a task
func DeleteTask(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var book models.Tasks
	if err := db.Where("task_id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&book)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

// Update a task
func UpdateStatus(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var task models.Tasks
	if err := db.Where("task_id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateTaskStatus
	if err := c.ShouldBindJSON(&input); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = ErrorMsg{fe.Field(), getErrorMsg(fe)}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
		}
		return
	}

	var updatedInput models.Tasks
	updatedInput.IsDone = input.Status

	db.Model(&task).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": task})
}
