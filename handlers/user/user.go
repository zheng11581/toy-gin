package user

import (
	"net/http"
	"zheng11581/toy-gin/handlers"

	"github.com/gin-gonic/gin"
)

func Get(ctx *gin.Context) {
	userId := ctx.Param("id")
	// 暂时只返回请求参数，无数据库处理
	handlers.WrapContext(ctx).Success(userId)
}

func GetV2(ctx *gin.Context) {
	userInfo, ok := ctx.Get("user_info")
	if !ok {
		handlers.WrapContext(ctx).Error(http.StatusInternalServerError, "获取用户信息失败")
		return
	}
	handlers.WrapContext(ctx).Success(userInfo)
}

type userListReq struct {
	Keyword string `json:"keyword"`
}

func List(ctx *gin.Context) {
	req := &userListReq{}
	err := ctx.ShouldBindJSON(req)
	if err != nil {
		handlers.WrapContext(ctx).Error(http.StatusInternalServerError, "绑定参数失败")
		return
	}
	// 暂时只返回请求参数，无数据库处理
	handlers.WrapContext(ctx).Success(req)

}

func Delete(ctx *gin.Context) {
	userId := ctx.Param("id")
	// 暂时只返回请求参数，无数据库处理
	handlers.WrapContext(ctx).Success(userId)
}

type userReq struct {
	UserName string
	Password string
	Age      int
	Sex      string
}

func Add(ctx *gin.Context) {
	req := &userReq{}
	err := ctx.ShouldBindJSON(req)
	if err != nil {
		handlers.WrapContext(ctx).Error(http.StatusInternalServerError, "绑定参数失败")
		return
	}
	// 暂时只返回请求参数，无数据库处理
	handlers.WrapContext(ctx).Success(req)

}

func Update(ctx *gin.Context) {
	req := &userReq{}
	err := ctx.ShouldBindJSON(req)
	if err != nil {
		handlers.WrapContext(ctx).Error(http.StatusInternalServerError, "绑定参数失败")
		return
	}
	// 暂时只返回请求参数，无数据库处理
	handlers.WrapContext(ctx).Success(req)
}
