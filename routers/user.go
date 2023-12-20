package routers

import "github.com/gin-gonic/gin"

func InitUser(group *gin.RouterGroup) {
	v1 := group.Group("/v1")
	v1.GET("/user")
	v1.POST("/user")
	v1.PUT("/user")
	v1.DELETE("/user")
}
