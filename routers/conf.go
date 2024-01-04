package routers

import (
	"zheng11581/toy-gin/handlers"

	"github.com/gin-gonic/gin"
)

func InitConf(group *gin.RouterGroup) {
	v1 := group.Group("/v1")
	v1.GET("/conf/:id", handlers.GetConf)
	v1.DELETE("/conf/:id", handlers.DeleteConf)
	v1.GET("/conf", handlers.ListConfs)
	v1.POST("/conf", handlers.AddConf)
	v1.PUT("/conf/:id", handlers.UpdateConf)

}
