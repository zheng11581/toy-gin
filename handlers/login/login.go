package login

import (
	"net/http"
	"time"
	"zheng11581/toy-gin/handlers"
	"zheng11581/toy-gin/middleware/plugin"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func Ping(ctx *gin.Context) {

}

type loginReq struct {
	UserName string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(ctx *gin.Context) {
	req := &loginReq{}
	err := ctx.ShouldBindJSON(req)
	if err != nil {
		handlers.WrapContext(ctx).Error(http.StatusInternalServerError, "参数不合法，绑定参数失败")
		return
	}
	data := plugin.Data{
		Name:   req.UserName,
		Age:    100,
		Gender: 1,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	sign, err := plugin.Sign(data)
	if err != nil {
		handlers.WrapContext(ctx).Error(http.StatusInternalServerError, "token 签名失败")
		return
	}
	// 暂时只返回假的token，无数据库处理
	handlers.WrapContext(ctx).Success(sign)
}

type registerReq struct {
	Username string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"omitempty,email"`
	Phone    string `json:"phone" binding:"required,e164"`
}

func Register(ctx *gin.Context) {
	req := &registerReq{}
	err := ctx.ShouldBindJSON(req)
	if err != nil {
		handlers.WrapContext(ctx).Error(http.StatusInternalServerError, "参数不合法，绑定参数失败")
		return
	}
	// 暂时只返回请求参数，无数据库处理
	handlers.WrapContext(ctx).Success(req)

}
