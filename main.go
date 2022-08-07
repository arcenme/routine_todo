package main

import "routine_todo/models"

func main() {
	db := models.SetupDB()
	db.AutoMigrate(&models.Tasks{})
}
