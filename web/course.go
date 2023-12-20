package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCourseDetail(ctx *gin.Context) {
	courseId := ctx.Param("id")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "获取课程详情成功",
		"id":      courseId,
	})
	mp := make(map[string]string)
	mp["id"] = courseId
	WrapContext(ctx).Success(mp)
}

func GetCourseVideo(ctx *gin.Context) {

}

func AddCourse(ctx *gin.Context) {

}

func PublishCourse(ctx *gin.Context) {

}

func UploadCourse(ctx *gin.Context) {

}
