package course

import (
	"net/http"
	"zheng11581/toy-gin/handlers"

	"github.com/gin-gonic/gin"
)

func Get(ctx *gin.Context) {
	courseId := ctx.Param("id")
	// 暂时只返回请求参数，无数据库处理
	handlers.WrapContext(ctx).Success(courseId)
}

type courseListReq struct {
	Keyword string `json:"keyword"`
}

func List(ctx *gin.Context) {
	req := &courseListReq{}
	err := ctx.ShouldBindJSON(req)
	if err != nil {
		handlers.WrapContext(ctx).Error(http.StatusInternalServerError, "绑定参数失败")
		return
	}
	// 暂时只返回请求参数，无数据库处理
	handlers.WrapContext(ctx).Success(req)
}

func Delete(ctx *gin.Context) {
	courseId := ctx.Param("id")
	// 暂时只返回请求参数，无数据库处理
	handlers.WrapContext(ctx).Success(courseId)
}

type courseReq struct {
	Name     string `json:"name"`
	Teacher  string `json:"teacher"`
	Duration int    `json:"duration"`
}

func Add(ctx *gin.Context) {
	req := &courseReq{}
	err := ctx.ShouldBindJSON(req)
	if err != nil {
		handlers.WrapContext(ctx).Error(http.StatusInternalServerError, "绑定参数失败")
		return
	}
	// 暂时只返回请求参数，无数据库处理
	handlers.WrapContext(ctx).Success(req)

}

func Update(ctx *gin.Context) {
	req := &courseReq{}
	err := ctx.ShouldBindJSON(req)
	if err != nil {
		handlers.WrapContext(ctx).Error(http.StatusInternalServerError, "绑定参数失败")
		return
	}
	// 暂时只返回请求参数，无数据库处理
	handlers.WrapContext(ctx).Success(req)
}
