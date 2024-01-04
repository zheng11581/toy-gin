package routers

import (
	"zheng11581/toy-gin/handlers"

	"github.com/gin-gonic/gin"
)

func InitLogin(group *gin.RouterGroup) {
	v1 := group.Group("/v1")
	v1.GET("/ping", handlers.Ping)
	v1.POST("/login", handlers.Login)
	v1.POST("/register", handlers.Register)
}
