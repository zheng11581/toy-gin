package routers

import (
	"zheng11581/toy-gin/handlers/rule"

	"github.com/gin-gonic/gin"
)

func InitRule(group *gin.RouterGroup) {
	v1 := group.Group("/v1")
	v1.GET("/rule/:id", rule.Get)
	v1.DELETE("/rule/:id", rule.Delete)
	v1.GET("/rule", rule.List)
	v1.POST("/rule", rule.Add)
	v1.PUT("/rule/:id")

}
