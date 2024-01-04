package routers

import (
	"zheng11581/toy-gin/handlers"
	"zheng11581/toy-gin/middleware"

	"github.com/gin-gonic/gin"
)

func InitCourse(group *gin.RouterGroup) {
	v1 := group.Group("/v1")
	v1.Use(middleware.AuthCheck)
	v1.GET("/course/:id", handlers.GetCourse)
	v1.GET("/course", handlers.ListCourses)
	v1.POST("/course", handlers.AddCourse)
	v1.PUT("/course/:id", handlers.UpdateCourse)
	v1.DELETE("/course/:id", handlers.DeleteCourse)
}
