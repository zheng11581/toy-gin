package handlers

import (
	"net/http"
	"strconv"
	"zheng11581/toy-gin/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetConf(ctx *gin.Context) {
	// 绑定参数
	confID, err := strconv.ParseUint(ctx.Param("id"), 8, 64)
	if err != nil {
		WrapContext(ctx).Error(http.StatusInternalServerError, "参数获取失败")
		return
	}

	// 查询数据 IngMonitorConf
	var conf models.Conf
	conf.ID = uint(confID)
	result := models.DB.First(&conf)
	if result.Error != nil {
		switch result.Error {
		case gorm.ErrRecordNotFound:
			WrapContext(ctx).Error(http.StatusNotFound, "查询数据失败")
		}
		return
	}

	// 返回结果
	WrapContext(ctx).Success(&conf)
}

func DeleteConf(ctx *gin.Context) {
	// 绑定参数
	confID, err := strconv.ParseUint(ctx.Param("id"), 8, 64)
	if err != nil {
		WrapContext(ctx).Error(http.StatusInternalServerError, "参数获取失败")
		return
	}

	// 删除数据 IngMonitorConf
	var conf models.Conf
	conf.ID = uint(confID)
	result := models.DB.Delete(&conf)
	if result.Error != nil {
		switch result.Error {
		case gorm.ErrRecordNotFound:
			WrapContext(ctx).Error(http.StatusNotFound, "删除数据失败")
		}
		return
	}

	// 返回结果
	WrapContext(ctx).Success(&conf)
}

func ListConfs(ctx *gin.Context) {
	// 绑定参数
	reqConf := &ConfBase{}
	err := ctx.ShouldBindJSON(reqConf)
	if err != nil {
		WrapContext(ctx).Error(http.StatusInternalServerError, "参数获取失败")
		return
	}

	// 查询列表 IngMonitorConf
	var confList []models.Conf
	result := models.DB.Where("host = ?", reqConf.Host).Find(&confList)
	if result.Error != nil {
		switch result.Error {
		case gorm.ErrRecordNotFound:
			WrapContext(ctx).Error(http.StatusNotFound, "查询数据失败")
		}
		return
	}

	// 返回结果
	WrapContext(ctx).Success(confList)

}

func AddConf(ctx *gin.Context) {
	// 绑定参数
	reqConf := ConfBase{}
	err := ctx.ShouldBindJSON(&reqConf)
	if err != nil {
		WrapContext(ctx).Error(http.StatusInternalServerError, "参数获取失败")
		return
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
			WrapContext(ctx).Error(http.StatusNotFound, "新增数据失败")
		}
		return
	}
	// 返回结果
	WrapContext(ctx).Success(conf)
}

func UpdateConf(ctx *gin.Context) {
	// 绑定参数
	confID, err := strconv.ParseUint(ctx.Param("id"), 8, 64)
	if err != nil {
		WrapContext(ctx).Error(http.StatusInternalServerError, "参数获取失败")
		return
	}
	reqConf := ConfBase{}
	err = ctx.ShouldBindJSON(&reqConf)
	if err != nil {
		WrapContext(ctx).Error(http.StatusInternalServerError, "参数获取失败")
		return
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
			WrapContext(ctx).Error(http.StatusNotFound, "更新数据失败")
		}
		return
	}
	// 返回结果
	WrapContext(ctx).Success(conf)
}
