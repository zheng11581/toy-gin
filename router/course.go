package router

import (
	"zheng11581/toy-gin/middleware"
	"zheng11581/toy-gin/web"

	"github.com/gin-gonic/gin"
)

func InitCourse(r *gin.Engine) {
	course := r.Group("/course")
	course.Use(middleware.TokenCheck)

	v1 := course.Group("/v1")
	v1.GET("/detail/:id", web.GetCourseDetail)
	v1.GET("/view/:id", web.GetCourseVideo)

	admin := course.Group("/admin")
	admin.Use(middleware.AuthCheck)
	adminV1 := admin.Group("/v1")
	adminV1.POST("/add", web.AddCourse)
	adminV1.POST("/publish", web.PublishCourse)
	adminV1.POST("/upload", web.UploadCourse)
}
