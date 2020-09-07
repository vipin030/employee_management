package main

import (
	"github.com/vipin030/employee_management/controllers"
	"github.com/vipin030/employee_management/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Connect to database
	models.ConnectDatabase()

	// Routes
	r.GET("/employees", controllers.FindEmployees)
	r.GET("/employees/:id", controllers.FindEmployee)
	r.POST("/employees", controllers.CreateEmployee)
	r.PATCH("/employees/:id", controllers.UpdateEmployee)
	r.DELETE("/employees/:id", controllers.DeleteEmployee)

	// Run the server
	r.Run(":8081")
}
