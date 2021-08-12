package model

import (
	"auth/config"
)

func Migrations() {
	conn, db := config.DataBase()
	defer db.Close()
	conn.AutoMigrate(
		&UserModel{},
		&Token{},
		&LoginAtivity{},
		&ServicesModel{},
		&ServicesEndpoind{},
		&UserPermissionModel{},
		&GroupPermissionModel{})
}
