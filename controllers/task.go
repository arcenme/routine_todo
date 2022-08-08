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
