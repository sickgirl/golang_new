package main

/**
路由分发
*/
import (
	"first_go/controller"
	"github.com/gin-gonic/gin"
)

func CollectRoutes(r *gin.Engine) *gin.Engine {
	r.POST("/user/auth/register", controller.Register)
	return r
}
