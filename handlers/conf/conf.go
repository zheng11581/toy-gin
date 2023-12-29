package conf

import (
	"encoding/json"
	"net/http"
	"zheng11581/toy-gin/handlers"
	"zheng11581/toy-gin/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Get(ctx *gin.Context) {
	confID := ctx.Param("id")
	conf := models.IngMonitorConf{}
	// 查询数据 IngMonitorConf
	result := models.DB.Where("id = ?", confID).First(&conf)
	if result.Error != nil {
		switch result.Error {
		case gorm.ErrRecordNotFound:
			handlers.WrapContext(ctx).Error(http.StatusNotFound, "记录未找到")
		}
	}

	// 保存到 ConfBase
	respConf := handlers.ConfBase{}
	confBytes, err := json.Marshal(conf)
	if err != nil {
		handlers.WrapContext(ctx).Error(http.StatusInternalServerError, "返回失败")
	}
	json.Unmarshal(confBytes, &respConf)
	handlers.WrapContext(ctx).Success(respConf)
	// models.IngressDB.Select("host", "name", "host")
}

func List(ctx *gin.Context) {

}
