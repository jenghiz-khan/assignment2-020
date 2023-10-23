package controllers

import (
	"github.com/jenghiz-khan/assignment2-020/initializers"
	"github.com/jenghiz-khan/assignment2-020/models"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateOrder(c *gin.Context) {
	var orderData struct {
		Customer_name string
		Items []models.Item
	}

	c.ShouldBindJSON(&orderData)
	
	order := models.Order{Customer_name: orderData.Customer_name,
		Items: orderData.Items}

	result := initializers.DB.Create(&order)

	if result.Error != nil {
		c.Status(400) 
		return
	}

	c.JSON(200, gin.H{
		"order": order,
	})
}

func GetAll(c *gin.Context) {
	var orders []models.Order

	result := initializers.DB.Preload("Items").Find(&orders)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"Data": orders,
	})
}

func GetOrder(c *gin.Context) {
	id := c.Param("id")

	var order models.Order

	result := initializers.DB.Preload("Items").First(&order, id)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("Order with id %v not found", id),
		})
		return
	}

	c.JSON(200, gin.H{
		"Data": order,
	})
}

func PatchOrder(c *gin.Context) {
	id := c.Param("id")

	var orderData struct {
		Customer_name string
		Items []models.Item
	}

	c.ShouldBindJSON(&orderData)
	
	var order models.Order
	initializers.DB.Preload("Items").Find(&order, id)

	fmt.Println(id, order)
	
	initializers.DB.Model(&order).Updates(models.Order{
		Customer_name: orderData.Customer_name,
		Items: orderData.Items,
	})

	initializers.DB.Session(&gorm.Session{FullSaveAssociations: true}).Model(&order).Association("Items").Replace(orderData.Items)

	c.JSON(200, gin.H{
		"code": 200,
		"data": order,
	})

}

func PutOrder(c *gin.Context) {
	id := c.Param("id")

	var orderData struct {
		Customer_name string
		Items []models.Item
	}

	c.ShouldBindJSON(&orderData)

	var order models.Order
	initializers.DB.Preload("Items").Find(&order, id)

	initializers.DB.Model(&order).Updates(models.Order{
		Customer_name: orderData.Customer_name,
		Items: orderData.Items,
	})

	initializers.DB.Session(&gorm.Session{FullSaveAssociations: true}).Model(&order).Association("Items").Replace(orderData.Items)

	c.JSON(200, gin.H{
		"code": 200,
		"data": order,
	})
}

func DeleteOrder(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Order{}, id)

	c.JSON(200, gin.H{
		"result": fmt.Sprintf("Order with id %v not found", id),
	})
}
