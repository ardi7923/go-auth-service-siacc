package internal_router

import (
	"auth/controller"

	"github.com/gin-gonic/gin"
)

func RouterInternalApiGateway(rg *gin.RouterGroup) {
	rg.GET("/check/permission", controller.CheckPermissionUser)
}
