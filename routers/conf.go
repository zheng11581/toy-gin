package routers

import (
	"zheng11581/toy-gin/handlers/conf"

	"github.com/gin-gonic/gin"
)

func InitConf(group *gin.RouterGroup) {
	v1 := group.Group("/v1")
	v1.GET("/conf/:id", conf.Get)
	v1.DELETE("/conf/:id", conf.Delete)
	v1.GET("/conf", conf.List)
	v1.POST("/conf", conf.Add)
	v1.PUT("/conf/:id", conf.Update)

}
