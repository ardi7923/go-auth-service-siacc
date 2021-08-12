package model

import (
	"auth/config"
	"fmt"
)

func (token Token) TableName() string { return "token" }

func deleteOldData(uuid string) {
	var token Token
	conn, db := config.DataBase()
	defer db.Close()
	if err := conn.Where("user_model_uuid = ? ", uuid).Unscoped().Delete(&token).Error; err != nil {
		fmt.Println(err)
	}
}
