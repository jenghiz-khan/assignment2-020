package main

import (
	"github.com/jenghiz-khan/assignment2-020/initializers"
	"github.com/jenghiz-khan/assignment2-020/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Item{})
	initializers.DB.AutoMigrate(&models.Order{})
}
