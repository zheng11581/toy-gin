package routers

import (
	"github.com/gin-gonic/gin"
)

func InitRule(group *gin.RouterGroup) {
	v1 := group.Group("/v1")
	v1.GET("/rule/:id")
	v1.DELETE("/rule/:id")
	v1.GET("/rule")
	v1.POST("/rule")
	v1.PUT("/rule/:id")

}
