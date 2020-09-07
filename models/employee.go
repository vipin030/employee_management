package models

type Employee struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Name   string `json:"name"`
	City   string `json:"city"`
	Salary int    `json:"salary"`
}
