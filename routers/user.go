package routers

import (
	"zheng11581/toy-gin/handlers"

	"github.com/gin-gonic/gin"
)

func InitUser(group *gin.RouterGroup) {
	v1 := group.Group("/v1")
	v1.GET("/user/:id", handlers.GetUser)
	v1.GET("/user", handlers.ListUsers)
	v1.POST("/user", handlers.AddUser)
	v1.PUT("/user/:id", handlers.UpdateUser)
	v1.DELETE("/user/:id", handlers.DeleteUser)

	v2 := group.Group("/v2")
	v2.GET("/user/:id", handlers.GetV2User)
}
