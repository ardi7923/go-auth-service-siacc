package external_router

import (
	"auth/controller"
	"auth/middleware"

	"github.com/gin-gonic/gin"
)

func RouterExternalAuth(rg *gin.RouterGroup) {
	rg.POST("/login", controller.LoginCheck)
	rg.POST("/token/check", controller.TokenCheck)
	rg.Use(middleware.RequiredLogin())
	{
		rg.PUT("/login/update", controller.LoginUpdate)
	}
}
