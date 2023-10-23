package models

import (
	"time"

	"gorm.io/gorm"
)

type Item struct {
	ID   		uint   	`json:"id" form:"id" gorm:"primaryKey;autoIncrement"`
	Item_code 	string  `json:"item_code"`
	Description string 	`json:"description" gorm:"type:varchar(500)"`
	Quantity  	uint	`json:"quantity"`
	OrderID   	uint	`json:"OrderID" form:"OrderID"`
	Created_at 	time.Time
	Updated_at 	time.Time
	Deleted_at 	gorm.DeletedAt `gorm:"index"`
}