package conf

import (
	"net/http"
	"strconv"
	"zheng11581/toy-gin/handlers"
	"zheng11581/toy-gin/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Get(ctx *gin.Context) {
	// 绑定参数
	confID, err := strconv.ParseUint(ctx.Param("id"), 8, 64)
	if err != nil {
		handlers.WrapContext(ctx).Error(http.StatusInternalServerError, "参数获取失败")
	}

	// 查询数据 IngMonitorConf
	var conf models.Conf
	conf.ID = uint(confID)
	result := models.DB.First(&conf)
	if result.Error != nil {
		switch result.Error {
		case gorm.ErrRecordNotFound:
			handlers.WrapContext(ctx).Error(http.StatusNotFound, "查询数据失败")
		}
	}

	// 返回结果
	handlers.WrapContext(ctx).Success(&conf)
}

func Delete(ctx *gin.Context) {
	// 绑定参数
	confID, err := strconv.ParseUint(ctx.Param("id"), 8, 64)
	if err != nil {
		handlers.WrapContext(ctx).Error(http.StatusInternalServerError, "参数获取失败")
	}

	// 删除数据 IngMonitorConf
	var conf models.Conf
	conf.ID = uint(confID)
	result := models.DB.Delete(&conf)
	if result.Error != nil {
		switch result.Error {
		case gorm.ErrRecordNotFound:
			handlers.WrapContext(ctx).Error(http.StatusNotFound, "删除数据失败")
		}
	}

	// 返回结果
	handlers.WrapContext(ctx).Success(&conf)
}

func List(ctx *gin.Context) {
	// 绑定参数
	reqConf := &handlers.ConfBase{}
	err := ctx.ShouldBindJSON(reqConf)
	if err != nil {
		handlers.WrapContext(ctx).Error(http.StatusInternalServerError, "参数获取失败")
	}

	// 查询列表 IngMonitorConf
	var confList []models.Conf
	result := models.DB.Where("host = ?", reqConf.Host).Find(&confList)
	if result.Error != nil {
		switch result.Error {
		case gorm.ErrRecordNotFound:
			handlers.WrapContext(ctx).Error(http.StatusNotFound, "查询数据失败")
		}
	}

	// 返回结果
	handlers.WrapContext(ctx).Success(confList)

}

func Add(ctx *gin.Context) {
	// 绑定参数
	reqConf := handlers.ConfBase{}
	err := ctx.ShouldBindJSON(&reqConf)
	if err != nil {
		handlers.WrapContext(ctx).Error(http.StatusInternalServerError, "参数获取失败")
	}
	// 新增数据
	var conf models.Conf
	conf.Host = reqConf.Host
	conf.Name = reqConf.Name
	conf.Conf = reqConf.Conf
	result := models.DB.Create(&conf)
	if result.Error != nil {
		switch result.Error {
		case gorm.ErrRecordNotFound:
			handlers.WrapContext(ctx).Error(http.StatusNotFound, "新增数据失败")
		}
	}
	// 返回结果
	handlers.WrapContext(ctx).Success(conf)
}

func Update(ctx *gin.Context) {
	// 绑定参数
	confID, err := strconv.ParseUint(ctx.Param("id"), 8, 64)
	if err != nil {
		handlers.WrapContext(ctx).Error(http.StatusInternalServerError, "参数获取失败")
	}
	reqConf := handlers.ConfBase{}
	err = ctx.ShouldBindJSON(&reqConf)
	if err != nil {
		handlers.WrapContext(ctx).Error(http.StatusInternalServerError, "参数获取失败")
	}
	// 更新数据
	var conf models.Conf
	conf.ID = uint(confID)
	conf.Host = reqConf.Host
	conf.Name = reqConf.Name
	conf.Conf = reqConf.Conf
	result := models.DB.Updates(&conf)
	if result.Error != nil {
		switch result.Error {
		case gorm.ErrRecordNotFound:
			handlers.WrapContext(ctx).Error(http.StatusNotFound, "更新数据失败")
		}
	}
	// 返回结果
	handlers.WrapContext(ctx).Success(conf)
}
