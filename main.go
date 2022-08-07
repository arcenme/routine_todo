package main

import (
	"html/template"
	"net/http"
	"routine_todo/models"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	// migrate database
	db := models.SetupDB()
	db.AutoMigrate(&models.Tasks{})

	// route
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	r.SetFuncMap(template.FuncMap{
		"upper": strings.ToUpper,
	})
	r.Static("/public/assets", "./public/assets")
	r.LoadHTMLGlob("public/pages/*.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Index",
		})
	})

	// i'm use port 4121
	r.Run(":4121")
}
