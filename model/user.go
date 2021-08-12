package model

import (
	"auth/config"
	"auth/helper"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	UUID                string                 `gorm:"not null;index;unique"`
	Username            string                 `json:"username,omitempty" binding:"required" gorm:"index; size:50; unique;" validate:"min=10"`
	Password            string                 `json:"password" binding:"required" gorm:"->;<-:create; size:300;"  validate:"min=10"`
	Status              bool                   `gorm:"default:true"`
	IsSuperAdmin        bool                   `gorm:"default:false;<-:false"`
	UserPermissionModel UserPermissionModel    `gorm:"foreignKey:UserModelUUID;references:uuid;constraint:OnDelete:CASCADE"`
	UserGroupsModel     []GroupPermissionModel `json:"user_groups_model" gorm:"many2many:user_groups;"`
	Token               Token                  `gorm:"foreignKey:UserModelUUID;constraint:OnDelete:CASCADE"`
}

type userModelResponse struct {
	UUID         string `json:"uuid"`
	Username     string `json:"username"`
	Status       bool   `json:"status"`
	IsSuperAdmin bool   `json:"is_super_admin"`
}

// HOOKS

func (usermodel UserModel) TableName() string { return "user_auth" }

func (this *UserModel) BeforeCreate(tx *gorm.DB) (err error) {
	pass, _ := helper.PasswordBecrypt(this.Password)
	this.Password = pass
	this.UUID = uuid.New().String()
	return
}

func (user *UserModel) AfterDelete(tx *gorm.DB) (err error) {
	conn, db := config.DataBase()
	defer db.Close()
	if err := conn.Where("uuid = ?", user.UUID).Unscoped().Delete(&user.Token).Delete(&user.UserPermissionModel).Error; err != nil {
		fmt.Println(err)
	}
	if err := conn.Where("user_model_id = ? ", user.ID).Unscoped().Delete(&user.UserPermissionModel).Error; err != nil {
		fmt.Println(err)
	}
	return
}

// ============ STRUCT FUNCTION CRUD ============= //
