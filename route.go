package main

import (
	"ginessential/controller"

	"github.com/gin-gonic/gin"
)

// CollectRoute 路由分发
func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controller.Register)
	return r
}
