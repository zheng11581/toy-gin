package routers

import (
	"zheng11581/toy-gin/handlers/user"

	"github.com/gin-gonic/gin"
)

func InitUser(group *gin.RouterGroup) {
	v1 := group.Group("/v1")
	v1.GET("/user/:id", user.Get)
	v1.GET("/user", user.List)
	v1.POST("/user", user.Add)
	v1.PUT("/user/:id", user.Update)
	v1.DELETE("/user/:id", user.Delete)
}
