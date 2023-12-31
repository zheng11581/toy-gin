package middleware

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var token = "123456"

func TokenCheck(ctx *gin.Context) {
	accessToken := ctx.GetHeader("access_token")

	if accessToken != token {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "token 校验失败",
		})
		ctx.AbortWithError(http.StatusInternalServerError, errors.New("token check failed"))
	}

	ctx.Set("user_name", "nick")
	ctx.Set("user_id", "100")
	log.Printf("TokenCheck: set user_name=%s, user_id=%d", "nick", 100)
	ctx.Next()
}
