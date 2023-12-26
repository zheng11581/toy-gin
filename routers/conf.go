package routers

import (
	"zheng11581/toy-gin/handlers/conf"

	"github.com/gin-gonic/gin"
)

func InitConf(group *gin.RouterGroup) {
	v1 := group.Group("/v1")
	v1.GET("/conf/:id", conf.Get)
}
