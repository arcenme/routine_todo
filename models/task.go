package models

import (
	"time"
)

type Tasks struct {
	TaskId    uint      `json:"id" gorm:"primary_key"`
	Task_Name string    `json:"task_name" gorm:"size:50;not null"`
	Assignee  string    `json:"assignee" gorm:"size:50;not null"`
	Deadline  time.Time `json:"deadline" gorm:"default:CURRENT_TIMESTAMP;not null"`
	IsDone    string    `json:"is_done" gorm:"default:0;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}
