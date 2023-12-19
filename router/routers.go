package router

import (
	"zheng11581/toy-gin/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouters(r *gin.Engine) {
	r.Use(middleware.LogInput)
	InitApi(r)
	InitCourse(r)
}
