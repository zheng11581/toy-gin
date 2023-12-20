package web

import (
	"net/http"

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
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, req)

}
