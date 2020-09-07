package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vipin030/employee_management/models"
)

type CreateEmployeeInput struct {
	Name   string `json:"name" binding:"required"`
	City   string `json:"city" binding:"required"`
	Salary int    `json:"salary" binding:"required"`
}

type UpdateEmployeeInput struct {
	Name   string `json:"name"`
	City  string `json:"name"`
	Salary int    `json:"salary"`
}

// Find all employees
func FindEmployees(c *gin.Context) {
	var employees []models.Employee
	models.DB.Find(&employees)

	c.JSON(http.StatusOK, gin.H{"data": employees})
}

// Find an employee
func FindEmployee(c *gin.Context) {
	// Get model if exist
	var employee models.Employee
	if err := models.DB.Where("id = ?", c.Param("id")).First(&employee).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": employee})
}

// Create new employee
func CreateEmployee(c *gin.Context) {
	// Validate input
	var input CreateEmployeeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create employee
	employee := models.Employee{Name: input.Name, City: input.City, Salary: input.Salary}
	models.DB.Create(&employee)

	c.JSON(http.StatusOK, gin.H{"data": employee})
}

// Update an employee
func UpdateEmployee(c *gin.Context) {
	// Get model if exist
	var employee models.Employee
	if err := models.DB.Where("id = ?", c.Param("id")).First(&employee).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateEmployeeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&employee).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": employee})
}

// Delete an employee
func DeleteEmployee(c *gin.Context) {
	// Get model if exist
	var employee models.Employee
	if err := models.DB.Where("id = ?", c.Param("id")).First(&employee).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&employee)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
