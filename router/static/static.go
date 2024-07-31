package router_static

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterStaticRoute(router *gin.Engine) {
	// 目录内文件都能访问
	router.Static("/assets", "./assets")
	// list 目录内文件
	router.StaticFS("/proj", http.Dir("./"))
	// 暴露某个文件
	router.StaticFile("/20200510.jpg", "./assets/20200510.jpg")
	router.LoadHTMLGlob("./templates/*")
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{"title": "梦幻"})
	})

}
