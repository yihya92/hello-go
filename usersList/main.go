package main

import (
	"log"
	"usersList/controllers"
	"usersList/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(gin.Logger())

	models.ConnectDatabase()

	users := r.Group("/users")
	{
		users.POST("/", controllers.CreateUser)
		users.GET("/", controllers.ListAllUsers)
		users.PUT("/:id", controllers.UpdateUser)
		users.GET("/one/:name", controllers.FindUserByName)
		users.GET("/all/:name", controllers.FindAllUsersByName)
		users.DELETE("/:id", controllers.DeleteUser)
		users.DELETE("/", controllers.DeleteAllUsers)
		users.POST("/multiple", controllers.InsertMultipleUsers)
	}
	log.Println("Server Statred!")
	//r.Run()
	r.Run("9000")
}
