package router

import (
	"github.com/gin-gonic/gin"
	"github.com/my-org/user-service/internal/handler"
)

// SetupRoutes 配置路由
func SetupRoutes(r *gin.RouterGroup) {
	// 用户相关路由
	users := r.Group("/users")
	{
		users.GET("", handler.GetUsers)
		users.GET("/:id", handler.GetUser)
		users.POST("", handler.CreateUser)
		users.PUT("/:id", handler.UpdateUser)
		users.DELETE("/:id", handler.DeleteUser)
	}

	// 认证相关路由
	auth := r.Group("/auth")
	{
		auth.POST("/login", handler.Login)
	}
}

