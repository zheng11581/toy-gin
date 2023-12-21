package routers

import (
	"zheng11581/toy-gin/handlers/course"
	"zheng11581/toy-gin/middleware"

	"github.com/gin-gonic/gin"
)

func InitCourse(group *gin.RouterGroup) {
	v1 := group.Group("/v1")
	v1.Use(middleware.AuthCheck)
	v1.GET("/course/:id", course.Get)
	v1.GET("/course", course.List)
	v1.POST("/course", course.Add)
	v1.PUT("/course/:id", course.Update)
	v1.DELETE("/course/:id", course.Delete)
}
