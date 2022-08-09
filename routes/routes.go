package routes

import (
	"html/template"
	"net/http"
	"routine_todo/controllers"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	r.SetFuncMap(template.FuncMap{
		"upper": strings.ToUpper,
	})
	r.Static("/public/assets", "./public/assets")
	r.LoadHTMLGlob("public/pages/*.html")

	// ====== VIEW
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Index",
		})
	})

	// ====== API
	r.GET("/api/routine", controllers.FindTasks)
	r.POST("/api/routine", controllers.CreateTask)
	r.GET("/api/routine/:id", controllers.FindTask)
	r.PATCH("/api/routine/:id", controllers.UpdateTask)
	r.DELETE("/api/routine/:id", controllers.DeleteTask)
	r.PATCH("/api/routine/done/:id", controllers.UpdateStatus)
	return r
}
