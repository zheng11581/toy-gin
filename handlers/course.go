package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCourse(ctx *gin.Context) {
	courseId := ctx.Param("id")
	// 暂时只返回请求参数，无数据库处理
	WrapContext(ctx).Success(courseId)
}

type courseListReq struct {
	Keyword string `json:"keyword"`
}

func ListCourses(ctx *gin.Context) {
	req := &courseListReq{}
	err := ctx.ShouldBindJSON(req)
	if err != nil {
		WrapContext(ctx).Error(http.StatusInternalServerError, "绑定参数失败")
		return
	}
	// 暂时只返回请求参数，无数据库处理
	WrapContext(ctx).Success(req)
}

func DeleteCourse(ctx *gin.Context) {
	courseId := ctx.Param("id")
	// 暂时只返回请求参数，无数据库处理
	WrapContext(ctx).Success(courseId)
}

type courseReq struct {
	Name     string `json:"name"`
	Teacher  string `json:"teacher"`
	Duration int    `json:"duration"`
}

func AddCourse(ctx *gin.Context) {
	req := &courseReq{}
	err := ctx.ShouldBindJSON(req)
	if err != nil {
		WrapContext(ctx).Error(http.StatusInternalServerError, "绑定参数失败")
		return
	}
	// 暂时只返回请求参数，无数据库处理
	WrapContext(ctx).Success(req)

}

func UpdateCourse(ctx *gin.Context) {
	req := &courseReq{}
	err := ctx.ShouldBindJSON(req)
	if err != nil {
		WrapContext(ctx).Error(http.StatusInternalServerError, "绑定参数失败")
		return
	}
	// 暂时只返回请求参数，无数据库处理
	WrapContext(ctx).Success(req)
}
