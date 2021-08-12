package router

import (
	"auth/router/external_router"
	"auth/router/internal_router"
	"net/http"

	"github.com/gin-gonic/gin"
)

func healtRouter(rg *gin.RouterGroup) {
	rg.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(
			http.StatusOK,
			gin.H{
				"message": "pong",
			})
	})

}

func SetupServer() *gin.Engine {
	r := gin.New()
	auth := r.Group("/auth")
	permission := r.Group("/permission")
	services := r.Group("/services")
	groups := r.Group("/groups")
	api_gateway := r.Group("/api_gateway")
	healt := r.Group("/healt")

	internal_router.RouterInternalAuth(auth)
	internal_router.RouterInternalService(services)
	internal_router.RouterInternalPermission(permission)
	internal_router.RouterInternalGroups(groups)
	internal_router.RouterInternalApiGateway(api_gateway)

	external_router.RouterExternalAuth(auth)
	external_router.RouterInternalPermission(permission)

	healtRouter(healt)

	return r
}
