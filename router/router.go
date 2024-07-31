package router

import (
	router_static "github.com/PerryHuan9/learn_go/router/static"
	router_user "github.com/PerryHuan9/learn_go/router/user"
	"github.com/gin-gonic/gin"
)

func RegisterRoute() *gin.Engine {
	server := gin.Default()
	router_user.RegisterUserRoute(server)
	router_static.RegisterStaticRoute(server)
	server.Run("localhost:8080")
	return server
}
