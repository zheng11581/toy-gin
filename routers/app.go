package routers

import "github.com/gin-gonic/gin"

func InitApp(group *gin.RouterGroup) {
	v1 := group.Group("/v1")
	v1.GET("/app/:id")
	v1.DELETE("/app/:id")
	v1.PUT("/app/:id")
	v1.GET("/app")
	v1.POST("/app")
}
