package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID      		uint `json:"id" form:"id" gorm:"primaryKey;autoIncrement"`
	Customer_name 	string `json:"customer_name" form:"customer_name" gorm:"not null;type:varchar(200)"`
	Ordered_at    	time.Time
	Created_at		time.Time
	Updated_at		time.Time
	Deleted_at		gorm.DeletedAt `gorm:"index"`
	Items			[]Item `gorm:"constraint:OnDelete:CASCADE;"`
}