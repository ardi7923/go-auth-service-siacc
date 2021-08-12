package controller

import (
	"auth/config"
	_struct "auth/controller/struct"
	"auth/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PermissionUserGetAll(ctx *gin.Context) {
	var data_json []model.UserPermissionModel
	conn, db := config.DataBase()
	defer db.Close()
	if err := conn.Preload("ServicesEndpoindID").Find(&data_json).Error; err != nil {
		ctx.SecureJSON(
			http.StatusNotFound,
			gin.H{
				"message": err.Error(),
				"data":    nil,
			})
		return
	}
	ctx.SecureJSON(
		http.StatusNotFound,
		gin.H{
			"message": "success",
			"data":    data_json,
		})
}

func PermissionUserInsert(ctx *gin.Context) {

	var data_json _struct.PermissionUserInsert
	var data_permissoin_user model.UserPermissionModel
	if err := ctx.ShouldBindJSON(&data_json); err != nil {
		ctx.SecureJSON(
			http.StatusNotAcceptable,
			gin.H{
				"message": err.Error(),
			})
		return
	}
	response, err := data_permissoin_user.Insert(&data_json)
	fmt.Println(data_json)

	if err != nil {
		ctx.SecureJSON(
			http.StatusNotAcceptable,
			gin.H{
				"message": err.Error(),
			})
		return
	}
	ctx.SecureJSON(
		http.StatusOK,
		gin.H{
			"message": "success",
			"data":    response,
		})
	return
}

func PermissionUserUpdate(ctx *gin.Context) {
	var data_json model.UserPermissionModel
	if err := ctx.ShouldBindJSON(&data_json); err != nil {
		ctx.SecureJSON(
			http.StatusNotAcceptable,
			gin.H{
				"message": err.Error(),
			})
		return
	}
	data_json.Update()
}

func PermissionUserDelete(ctx *gin.Context) {
	var data_json _struct.PermisisonUserDelete
	if err := ctx.ShouldBindJSON(&data_json); err != nil {
		ctx.SecureJSON(
			http.StatusNotAcceptable,
			gin.H{
				"message": err.Error(),
			})
		return
	}
	model.UserPermissionModel{}.Delete(data_json.UUID)
}
