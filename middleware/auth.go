package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func AuthCheck(ctx *gin.Context) {
	userName, _ := ctx.Get("user_name")
	userId := ctx.GetString("user_id")
	fmt.Printf("called auth check, userName: %s, userId: %s", userName, userId)
	ctx.Next()
}
