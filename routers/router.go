package routers

import (
	"zheng11581/toy-gin/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouters(r *gin.Engine) {
	api := r.Group("/api")
	api.Use(middleware.LogInput)
	InitLogin(api)
	InitCourse(api)
	InitUser(api)
}
