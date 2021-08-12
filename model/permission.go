package model

import (
	"auth/config"

	"gorm.io/gorm"
)

type UserPermissionModel struct {
	gorm.Model
	UserModelUUID      string             `gorm:"unique" json:"user_model_uuid,omitempty"`
	ServicesEndpoindID []ServicesEndpoind `json:"endpoind_id,omitempty" gorm:"many2many:user_endpoind_permission"`
	Status             bool               `gorm:"default:true" json:"status,omitempty"`
}

type GroupPermissionModel struct {
	gorm.Model
	Name       string             `json:"name" binding:"required" gorm:"unique"`
	EndpoindID []ServicesEndpoind `json:"endpoind,string" gorm:"many2many:group_endpoind_permission"`
	Status     bool               `gorm:"default:true"`
}

func (permissionUser UserPermissionModel) BeforeUpdate(tx *gorm.DB) (err error) {
	conn, db := config.DataBase()
	defer db.Close()
	conn.Model(&permissionUser).Association("ServicesEndpoindID").Clear()
	return
}
