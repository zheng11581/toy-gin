package routers

import (
	"zheng11581/toy-gin/handlers"

	"github.com/gin-gonic/gin"
)

func InitApp(group *gin.RouterGroup) {
	v1 := group.Group("/v1")
	v1.GET("/app", handlers.GetApps)
	v1.POST("/app", handlers.AddApp)
}
