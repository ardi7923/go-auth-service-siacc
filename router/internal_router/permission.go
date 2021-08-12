package internal_router

import (
	"auth/controller"

	"github.com/gin-gonic/gin"
)

func RouterInternalPermission(rg *gin.RouterGroup) {
	rg.POST("/user/create", controller.PermissionUserInsert)
	rg.PUT("/user/update", controller.PermissionUserUpdate)
	rg.DELETE("/user/delete", controller.PermissionUserDelete)
	rg.GET("user/", controller.PermissionUserGetAll)
}
