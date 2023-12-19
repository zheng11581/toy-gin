package router

import (
	"zheng11581/toy-gin/web"

	"github.com/gin-gonic/gin"
)

func InitApi(r *gin.Engine) {
	api := r.Group("/api")
	v1 := api.Group("/v1")
	v1.GET("/ping", web.Ping)
	v1.POST("/login", web.Login)
	v1.POST("/register", web.Register)
}
