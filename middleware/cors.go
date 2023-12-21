package middleware

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORS1(ctx *gin.Context) {
	cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods: []string{
			"GET", "POST", "PUT", "DELETE", "OPTOPNS", "HEAD",
		},
		AllowHeaders: []string{
			"Origin", "Content-Length", "Content-Type",
		},
	})
	log.Printf("CORS: config CORS")
	ctx.Next()
}

func CORS() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods: []string{
			"GET", "POST", "PUT", "DELETE", "OPTOPNS", "HEAD",
		},
		AllowHeaders: []string{
			"Origin", "Content-Length", "Content-Type",
		},
	})
}
