package main

/**
路由分发
*/
import (
	"first_go/controller"
	"first_go/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoutes(r *gin.Engine) *gin.Engine {
	r.POST("/user/auth/register", controller.Register)
	r.POST("/user/auth/login", controller.Login)
	r.GET("/user/auth/info", middleware.AuthMiddleware(), controller.Info)
	return r
}
