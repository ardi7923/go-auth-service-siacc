package middleware

import (
	"auth/helper"
	"auth/model"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func cekTokenAuthDatabase(token string) (bool, error) {

	var uuid string
	var user model.UserModel
	claims, err_claims := helper.GetPayloadToken(token)

	if err_claims != nil {
		return false, err_claims
	}

	uuid = claims["UUID"].(string)
	user.UUID = uuid

	if _, err := user.TokenCheck(uuid); err != nil {
		return false, err
	}
	return true, nil
}

func RequiredLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.Request.Header["Authorization"]
		if len(header) == 0 {
			ctx.Abort()
			ctx.SecureJSON(
				http.StatusUnauthorized,
				gin.H{
					"message": "required Authorization",
					"data":    nil,
				})
			return
		}
		token := strings.Split(header[0], "bearer ")
		if match, err := helper.TokenCheck(token[1]); match == false && err != nil {
			ctx.Abort()
			ctx.SecureJSON(
				http.StatusUnauthorized,
				gin.H{
					"message": "your token not valid",
					"data":    nil,
				})
			return
		}
		if _, err := cekTokenAuthDatabase(token[1]); err != nil {
			ctx.Abort()
			ctx.SecureJSON(
				http.StatusUnauthorized,
				gin.H{
					"message": err.Error(),
					"data":    nil,
				})
			return
		}
		ctx.Next()
	}
}
