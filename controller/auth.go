package controller

import (
	"auth/config"
	_struct "auth/controller/struct"
	"auth/helper"
	"auth/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

var i = 0

func LoginList(ctx *gin.Context) {
	user, err := model.UserModel{}.GetAll()
	if err != nil {
		ctx.SecureJSON(
			http.StatusNotAcceptable,
			gin.H{
				"data":    nil,
				"message": err,
			})
		return
	}
	ctx.SecureJSON(
		http.StatusOK,
		gin.H{
			"data":    user,
			"message": "success get data",
		})
}

func LoginCreate(ctx *gin.Context) {
	var data_json model.UserModel
	if err := ctx.ShouldBindJSON(&data_json); err != nil {
		ctx.SecureJSON(
			http.StatusNotAcceptable,
			gin.H{
				"message": err.Error(),
			})
		return
	}
	response, err := data_json.Insert()
	if err != nil {
		ctx.SecureJSON(
			http.StatusNotAcceptable,
			gin.H{
				"message": err,
			})
		return
	}
	ctx.SecureJSON(
		http.StatusOK,
		gin.H{
			"message": "success input user",
			"data":    response,
		})
}

func LoginUpdate(ctx *gin.Context) {
	var user_model model.UserModel
	var data_json _struct.UpdateUser
	if err := ctx.ShouldBindJSON(&data_json); err != nil {
		ctx.SecureJSON(
			http.StatusNotAcceptable,
			gin.H{
				"message": err.Error(),
			})
		return
	}
	response, err := user_model.Update(data_json.UUID, user_model)
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
			"message": "success update",
			"data":    response,
		})
	return
}

func LoginDelete(ctx *gin.Context) {
	var struct_json _struct.UserDelete
	if err := ctx.ShouldBindJSON(&struct_json); err != nil {
		ctx.SecureJSON(
			http.StatusNotAcceptable,
			gin.H{
				"message": err.Error(),
			})
		return
	}
	response, err := model.UserModel{}.Delete(struct_json.UUID)
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

func LoginCheck(ctx *gin.Context) {
	var data_json model.UserModel
	var old_data model.UserModel
	if err := ctx.ShouldBindJSON(&data_json); err != nil {
		ctx.SecureJSON(
			http.StatusNotAcceptable,
			gin.H{
				"message": err.Error(),
			})
		return
	}
	conn, db := config.DataBase()
	defer db.Close()
	res := conn.Where(&model.UserModel{Username: data_json.Username}).First(&old_data)

	if res.Error != nil {
		ctx.SecureJSON(
			http.StatusOK,
			gin.H{
				"message": res.Error,
				"data":    nil,
			})
		return
	}
	if !old_data.Status {
		ctx.SecureJSON(
			http.StatusUnauthorized,
			gin.H{
				"message": "user not active",
				"data": gin.H{
					"token":  nil,
					"result": false,
				},
			})
		return
	}
	if helper.PasswordMatch(data_json.Password, []byte(old_data.Password)) == false {
		ctx.SecureJSON(
			http.StatusUnauthorized,
			gin.H{
				"message": "error login",
				"data": gin.H{
					"token":  nil,
					"result": false,
				},
			})
		return
	}
	token, err := helper.GetToken(old_data.UUID)

	if err == nil {
		ctx.SecureJSON(
			http.StatusOK,
			gin.H{
				"message": "login success",
				"data": gin.H{
					"token":  token,
					"result": true,
				},
			})
	}

	old_data.Token = model.Token{Token: token}
	conn.Updates(&old_data)
	// config.DataBase().Save(&old_data)
}
