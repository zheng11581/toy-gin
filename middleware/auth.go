package middleware

import (
	"log"
	"net/http"
	"zheng11581/toy-gin/handlers"
	"zheng11581/toy-gin/middleware/plugin"

	"github.com/gin-gonic/gin"
)

func AuthCheck(ctx *gin.Context) {
	accessToken := ctx.GetHeader("access_token")
	log.Printf("AuthCheck: get access_token=%s", accessToken)
	data := &plugin.Data{}
	err := plugin.Verify(accessToken, data)
	if err != nil {
		handlers.WrapContext(ctx).Error(http.StatusInternalServerError, "身份认证失败")
		ctx.Abort()
	}
	ctx.Set("user_info", data)
	ctx.Next()
}
