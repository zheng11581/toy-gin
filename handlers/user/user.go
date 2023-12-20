package user

import (
	"zheng11581/toy-gin/handlers"

	"github.com/gin-gonic/gin"
)

type userReq struct {
	UserId   string
	UserName string
	Password string
	Age      int
	Sex      string
}

func Get(ctx *gin.Context) {
	userId := ctx.Param("id")
	user := &userReq{}
	user.UserId = userId
	handlers.WrapContext(ctx).Success(user)
}

func List(ctx *gin.Context) {

}

func Update(ctx *gin.Context) {

}

func Delete(ctx *gin.Context) {

}

func Add(ctx *gin.Context) {

}
