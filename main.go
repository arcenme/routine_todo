package main

import (
	"routine_todo/models"
	"routine_todo/routes"
)

func main() {
	// migrate database
	db := models.SetupDB()
	db.AutoMigrate(&models.Tasks{})

	// route
	r := routes.SetupRoutes(db)

	// i'm use port 4121
	r.Run(":4121")
}
