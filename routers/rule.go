package routers

import (
	"zheng11581/toy-gin/handlers"

	"github.com/gin-gonic/gin"
)

func InitRule(group *gin.RouterGroup) {
	v1 := group.Group("/v1")
	v1.GET("/rule/:id", handlers.GetRule)
	v1.DELETE("/rule/:id", handlers.DeleteRule)
	v1.GET("/rule", handlers.ListRules)
	v1.POST("/rule", handlers.AddRule)
	v1.PUT("/rule/:id", handlers.UpdateRule)

	v1.GET("/rule/special/:id", handlers.GetSpecialRule)
	v1.DELETE("/rule/special/:id", handlers.DeleteSpecialRule)
	v1.GET("/rule/special", handlers.ListSpecialRule)
	v1.POST("/rule/special", handlers.AddSpecialRule)
	v1.PUT("/rule/special/:id", handlers.UpdateSpecialRule)

	v1.GET("/rule/silence/:id", handlers.GetSilenceRule)
	v1.DELETE("/rule/silence/:id", handlers.DeleteSilenceRule)
	v1.GET("/rule/silence", handlers.ListSilenceRule)
	v1.POST("/rule/silence", handlers.AddSilenceRule)
	v1.PUT("/rule/silence/:id", handlers.UpdateSilenceRule)

}
