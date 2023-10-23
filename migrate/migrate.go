package main

import (
	"assignment-2/initializers"
	"assignment-2/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Item{})
	initializers.DB.AutoMigrate(&models.Order{})
}