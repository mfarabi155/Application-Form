package main

import (
	"application-form/controllers"
	"application-form/database"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDatabase()

	r := gin.Default()

	config := cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}

	r.Use(cors.New(config))

	r.POST("/persons", controllers.CreatePerson)
	r.GET("/persons", controllers.FindPersons)
	r.GET("/persons/:identity_number", controllers.FindPersonByIdentityNumber)
	r.PUT("/persons/:identity_number", controllers.UpdatePersonByIdentityNumber)
	r.DELETE("/persons/:identity_number", controllers.DeletePersonByIdentityNumber)

	r.Run(":8080")
}
