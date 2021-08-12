package internal_router

import (
	"auth/controller"

	"github.com/gin-gonic/gin"
)

func RouterInternalService(rg *gin.RouterGroup) {
	rg.POST("/create", controller.ServicesInsert)
	rg.POST("/endpoind/insert", controller.EndpoindInsert)
	rg.PUT("/update", controller.ServiceUpdate)
	rg.PUT("/endpoind/update", controller.EndpointUpdate)
	rg.DELETE("/delete", controller.ServiceDelete)
	rg.GET("/", controller.ServicesGetAll)
}
