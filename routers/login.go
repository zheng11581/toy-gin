package routers

import (
	"zheng11581/toy-gin/handlers/login"

	"github.com/gin-gonic/gin"
)

func InitLogin(group *gin.RouterGroup) {
	v1 := group.Group("/v1")
	v1.GET("/ping", login.Ping)
	v1.POST("/login", login.Login)
	v1.POST("/register", login.Register)
}
