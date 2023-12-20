package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

func AuthCheck(ctx *gin.Context) {
	userName, _ := ctx.Get("user_name")
	userId := ctx.GetString("user_id")
	log.Printf("AuthCheck: get userName: %s, userId: %s", userName, userId)
	ctx.Next()
}
