package internal_router

import (
	"auth/controller"

	"github.com/gin-gonic/gin"
)

func RouterInternalGroups(rg *gin.RouterGroup) {
	rg.POST("/create", controller.GroupInsert)
	rg.POST("/insert/user", controller.UserGroupInsert)
	rg.POST("/insert/endpoind", controller.GroupEndpoindInsert)
	rg.DELETE("/delete", controller.GroupDelete)
	rg.PUT("/update", controller.GroupUpdate)
	rg.GET("/", controller.GroupGetALl)
}
