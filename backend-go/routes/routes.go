package routes

import (
	"github.com/ZhangYongChang/backend-go/controller"
	"github.com/ZhangYongChang/backend-go/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CorsMiddleware())
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.POST("/api/auth/info", middleware.AuthMiddleware(), controller.QueryUser)
	return r
}
