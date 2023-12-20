package middleware

import (
	"bytes"
	"io"
	"log"

	"github.com/gin-gonic/gin"
)

func LogInput(ctx *gin.Context) {
	requestBody, _ := ctx.GetRawData()
	ctx.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))

	mp := make(map[string]interface{})
	mp["request_url"] = ctx.Request.RequestURI
	mp["status"] = ctx.Writer.Status()
	mp["method"] = ctx.Request.Method
	mp["access_token"] = ctx.GetHeader("access_token")
	mp["body"] = string(requestBody)
	log.Printf("LogInput: %v", mp)
	ctx.Next()
	// log.Println("LogInput: After Next()")

}
