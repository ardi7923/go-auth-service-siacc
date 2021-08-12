package internal_router

import (
	"auth/controller"

	"github.com/gin-gonic/gin"
)

func RouterInternalAuth(rg *gin.RouterGroup) {
	rg.POST("/login/create", controller.LoginCreate)
	rg.GET("/login/list", controller.LoginList)
	rg.DELETE("/login/delete", controller.LoginDelete)
}
