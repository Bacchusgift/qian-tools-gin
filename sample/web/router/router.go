package router

import (
	"github.com/gin-gonic/gin"
	"qian-tools-gin/middleware"
)

// 注册路由

func Register(engine *gin.Engine) {
	mainRouter := engine.Group("/", func(context *gin.Context) {
	}).Use(middleware.LimitMiddleware(1))

	mainRouter.GET("/demo", func(context *gin.Context) {
		context.JSON(200, "success")
	})
}
