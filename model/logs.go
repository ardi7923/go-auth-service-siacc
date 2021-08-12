package model

import (

	"gorm.io/gorm"
)

type LoginAtivity struct {
	gorm.Model
	UserModelUUID string
	ip_address    net.IPAddr
}

func (loginActivity LoginAtivity) TableName() string { return "login_activity" }
