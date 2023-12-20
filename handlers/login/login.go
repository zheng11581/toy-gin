package login

import (
	"net/http"
	"zheng11581/toy-gin/handlers"

	"github.com/gin-gonic/gin"
)

func Ping(ctx *gin.Context) {

}

func Login(ctx *gin.Context) {

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
	handlers.WrapContext(ctx).Success(req)

}
