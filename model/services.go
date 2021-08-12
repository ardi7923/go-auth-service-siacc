package model

import (
	"fmt"

	"gorm.io/gorm"
)

type ServicesModel struct {
	gorm.Model
	Name         string             `json:"service_name" binding:"required" gorm:"unique; index; type:varchar(100);"`
	Url_services string             `json:"url_services" binding:"required" gorm:"unique;"`
	Endpoind     []ServicesEndpoind `json:"endpoind" gorm:"foreigenKey:ServicesModelID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Status       bool               `gorm:"default:true"`
}

type ServicesEndpoind struct {
	gorm.Model
	ServicesModelID uint   `json:"services_id" gorm:"index"`
	Name            string `json:"endpoind" binding:"required"`
	Method          string `json:"method" binding:"required"`
	Status          bool
}

// HOOKS FUNCTION

func (services *ServicesModel) TableName() string { return "list_services" }

func (services *ServicesModel) AfterDelete(tx *gorm.DB) (err error) {
	fmt.Println("delete services")
	fmt.Println("delete after")
	return
}
