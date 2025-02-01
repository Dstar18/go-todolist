package main

import (
	"go-todolist/controllers"
	"go-todolist/database"
	"go-todolist/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// connect database
	database.Connect()

	// migration schema
	database.DB.AutoMigrate(&models.User{})
	database.DB.AutoMigrate(&models.Notes{})
	database.DB.AutoMigrate(&models.Items{})

	// initialize Gin
	r := gin.Default()

	// Routes public
	r.POST("register", controllers.StoreUser)

	// Routes user
	r.GET("users", controllers.GetUsers)

	// Routes notes
	r.GET("notes", controllers.GetNotes)
	r.POST("notes", controllers.StoreNotes)
	r.PUT("notes/:id", controllers.UpdateNotes)
	r.DELETE("notes/:id", controllers.DestroyNotes)

	// Routes items
	r.GET("item/:id", controllers.ShowItems)
	r.DELETE("item/:id", controllers.DestroyItems)

	// run server
	r.Run(":3000")
}
