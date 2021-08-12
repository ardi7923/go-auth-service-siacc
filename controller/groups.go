package controller

import (
	"auth/config"
	_struct "auth/controller/struct"
	"auth/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GroupGetALl(ctx *gin.Context) {
	var data []model.GroupPermissionModel
	conn, db := config.DataBase()
	defer db.Close()
	conn.Preload("EndpoindID").Find(&data)
	ctx.SecureJSON(
		http.StatusNotAcceptable,
		gin.H{
			"data":    data,
			"message": "success",
		})
}

func GroupInsert(ctx *gin.Context) {
	var data_groups_user model.GroupPermissionModel
	conn, db := config.DataBase()
	defer db.Close()
	if err := ctx.ShouldBindJSON(&data_groups_user); err != nil {
		ctx.SecureJSON(
			http.StatusNotAcceptable,
			gin.H{
				"message": err.Error(),
			})
		return
	}
	if err := conn.Create(&data_groups_user).Error; err != nil {
		ctx.SecureJSON(
			http.StatusNotAcceptable,
			gin.H{
				"data":    nil,
				"message": err.Error(),
			})
		return
	}
	ctx.SecureJSON(
		http.StatusOK,
		gin.H{
			"data":    nil,
			"message": data_groups_user,
		})
}

func GroupEndpoindInsert(ctx *gin.Context) {
	conn, db := config.DataBase()
	defer db.Close()
	var data_json _struct.GroupsEndpoindInsert
	var data_endpoind []model.ServicesEndpoind
	var data_groups model.GroupPermissionModel

	if err := ctx.ShouldBindJSON(&data_json); err != nil {
		ctx.SecureJSON(
			http.StatusNotAcceptable,
			gin.H{
				"message": err.Error(),
			})
		return
	}

	if err := conn.Where("id IN (?)", data_json.Endpoind).Find(&data_endpoind).Error; err != nil {
		ctx.SecureJSON(
			http.StatusNotAcceptable,
			gin.H{
				"message": err.Error(),
			})
		return
	}

	if err := conn.Where("id = ?", data_json.ID).First(&data_groups).Association("EndpoindID").Append(&data_endpoind); err != nil {
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
			"data":    data_endpoind,
			"message": "success insert endpoind",
		})
}

func UserGroupInsert(ctx *gin.Context) {
	conn, db := config.DataBase()
	defer db.Close()
	var data_json _struct.UserGroup
	var user model.UserModel
	var data_groups []model.GroupPermissionModel

	if err := ctx.ShouldBindJSON(&data_json); err != nil {
		ctx.SecureJSON(
			http.StatusNotAcceptable,
			gin.H{
				"message": err.Error(),
			})
		return
	}

	if err := conn.Where("id IN (?)", data_json.GroupsID).Find(&data_groups).Error; err != nil {
		ctx.SecureJSON(
			http.StatusNotAcceptable,
			gin.H{
				"message": err.Error(),
			})
		return
	}

	if err := conn.Where("uuid = ? ", data_json.UUID_User).First(&user).Association("UserGroupsModel").Append(&data_groups); err != nil {
		ctx.SecureJSON(
			http.StatusNotAcceptable,
			gin.H{
				"message": err.Error(),
			})
		return
	}

	ctx.SecureJSON(
		http.StatusNotAcceptable,
		gin.H{
			"message": data_groups,
		})
}

func GroupDelete(ctx *gin.Context) {
	var data_json _struct.GroupsDelete
	var group model.GroupPermissionModel
	if err := ctx.ShouldBindJSON(&data_json); err != nil {
		ctx.SecureJSON(
			http.StatusNotAcceptable,
			gin.H{
				"message": err.Error(),
			})
		return
	}

	response, err := group.Delete(data_json.ID)
	fmt.Println("response")
	fmt.Println(response)
	if err != nil {
		ctx.SecureJSON(
			http.StatusNotAcceptable,
			gin.H{
				"data":    nil,
				"message": err.Error(),
			})
		return
	}

	ctx.SecureJSON(
		http.StatusOK,
		gin.H{
			"data":    response,
			"message": "success delete",
		})
}

func GroupUpdate(ctx *gin.Context) {
	var group model.GroupPermissionModel

	if err := ctx.ShouldBindJSON(&group); err != nil {
		ctx.SecureJSON(
			http.StatusNotAcceptable,
			gin.H{
				"message": err.Error(),
			})
		return
	}
	response, err := group.Update()
	if err != nil {
		ctx.SecureJSON(
			http.StatusNotAcceptable,
			gin.H{
				"data":    nil,
				"message": err.Error(),
			})
		return
	}
	ctx.SecureJSON(
		http.StatusOK,
		gin.H{
			"data":    response,
			"message": "success insert",
		})
}
