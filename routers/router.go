package routers

import (
	"zheng11581/toy-gin/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouters(r *gin.Engine) {
	api := r.Group("/api")
	api.Use(middleware.LogInput, middleware.CORS(), middleware.AuthCheck)
	InitCourse(api)
	InitUser(api)

	noAuthApi := r.Group("/api")
	InitLogin(noAuthApi)

}
