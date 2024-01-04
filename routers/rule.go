package routers

import (
	"zheng11581/toy-gin/handlers"

	"github.com/gin-gonic/gin"
)

func InitRule(group *gin.RouterGroup) {
	v1 := group.Group("/v1")
	v1.GET("/rule/:id", handlers.GetRule)
	v1.DELETE("/rule/:id", handlers.DeleteRule)
	v1.GET("/rule", handlers.ListRules)
	v1.POST("/rule", handlers.AddRule)
	v1.PUT("/rule/:id", handlers.UpdateRule)

}
