package main

import (
	"application-form/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/persons", controllers.CreatePerson)
	router.GET("/persons", controllers.FindPersons)
	router.GET("/persons/:identity_number", controllers.FindPersonByIdentityNumber)
	router.PUT("/persons/:identity_number", controllers.UpdatePersonByIdentityNumber)
	router.DELETE("/persons/:identity_number", controllers.DeletePersonByIdentityNumber)

	router.Run(":8080")
}
