package controller

import (
	"auth/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type token struct {
	Token string `json:"token" binding:"required"`
}

func TokenCheck(ctx *gin.Context) {
	var data_json token
	if err := ctx.ShouldBindJSON(&data_json); err != nil {
		ctx.SecureJSON(
			http.StatusNotAcceptable,
			gin.H{
				"message": err.Error(),
			})
		return
	}
	match, err := helper.TokenCheck(data_json.Token)
	if err != nil {
		ctx.SecureJSON(
			http.StatusBadRequest,
			gin.H{
				"message": "Token Check Error",
				"data":    nil,
			})
		return
	} else if match == false {
		ctx.SecureJSON(
			http.StatusUnauthorized,
			gin.H{
				"message": "token is invalid",
				"data":    nil,
			})
		return
	}
	ctx.SecureJSON(
		http.StatusOK,
		gin.H{
			"message": "token is Valid",
			"data":    match,
		})
}
