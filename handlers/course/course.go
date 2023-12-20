package course

import (
	"net/http"
	"zheng11581/toy-gin/handlers"

	"github.com/gin-gonic/gin"
)

func Get(ctx *gin.Context) {
	courseId := ctx.Param("id")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "获取课程详情成功",
		"id":      courseId,
	})
	mp := make(map[string]string)
	mp["id"] = courseId
	handlers.WrapContext(ctx).Success(mp)
}

func List(ctx *gin.Context) {

}

func Update(ctx *gin.Context) {

}

func Delete(ctx *gin.Context) {

}

func Add(ctx *gin.Context) {

}
