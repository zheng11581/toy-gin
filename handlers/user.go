package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(ctx *gin.Context) {
	userId := ctx.Param("id")
	// 暂时只返回请求参数，无数据库处理
	WrapContext(ctx).Success(userId)
}

func GetV2User(ctx *gin.Context) {
	userInfo, ok := ctx.Get("user_info")
	if !ok {
		WrapContext(ctx).Error(http.StatusInternalServerError, "获取用户信息失败")
		return
	}
	WrapContext(ctx).Success(userInfo)
}

type userListReq struct {
	Keyword string `json:"keyword"`
}

func ListUsers(ctx *gin.Context) {
	req := &userListReq{}
	err := ctx.ShouldBindJSON(req)
	if err != nil {
		WrapContext(ctx).Error(http.StatusInternalServerError, "绑定参数失败")
		return
	}
	// 暂时只返回请求参数，无数据库处理
	WrapContext(ctx).Success(req)

}

func DeleteUser(ctx *gin.Context) {
	userId := ctx.Param("id")
	// 暂时只返回请求参数，无数据库处理
	WrapContext(ctx).Success(userId)
}

type userReq struct {
	UserName string
	Password string
	Age      int
	Sex      string
}

func AddUser(ctx *gin.Context) {
	req := &userReq{}
	err := ctx.ShouldBindJSON(req)
	if err != nil {
		WrapContext(ctx).Error(http.StatusInternalServerError, "绑定参数失败")
		return
	}
	// 暂时只返回请求参数，无数据库处理
	WrapContext(ctx).Success(req)

}

func UpdateUser(ctx *gin.Context) {
	req := &userReq{}
	err := ctx.ShouldBindJSON(req)
	if err != nil {
		WrapContext(ctx).Error(http.StatusInternalServerError, "绑定参数失败")
		return
	}
	// 暂时只返回请求参数，无数据库处理
	WrapContext(ctx).Success(req)
}
