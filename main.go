package main

import (
	"github.com/jenghiz-khan/assignment2-020/controllers"
	"github.com/jenghiz-khan/assignment2-020/initializers"

	"github.com/gin-gonic/gin"
)

var PORT = ":8080"

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/order", controllers.CreateOrder)
	r.GET("/orders", controllers.GetAll)
	r.GET("/order/:id", controllers.GetOrder)
	r.PATCH("/order/:id", controllers.PatchOrder)
	r.PUT("/order/:id", controllers.PutOrder)
	r.DELETE("/order/:id", controllers.DeleteOrder)

	r.Run()
}
