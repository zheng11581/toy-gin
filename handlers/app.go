package handlers

import (
	"net/http"
	"zheng11581/toy-gin/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetApps(ctx *gin.Context) {
	// 获取、绑定请求参数
	reqApp := AppBase{}
	err := ctx.ShouldBindJSON(&reqApp)
	if err != nil {
		WrapContext(ctx).Error(http.StatusInternalServerError, "绑定参数失败")
		return
	}
	// 根据app_code查询App
	var appList []models.App
	result := models.DB.Where("specify_app_code LIKE ?", "%"+reqApp.SpecifyAppCode+"%").Or("specify_app_name LIKE ?", "%"+reqApp.SpecifyAppName+"%").Find(&appList)
	if result.Error != nil {
		switch result.Error {
		case gorm.ErrRecordNotFound:
			WrapContext(ctx).Error(http.StatusNotFound, "查询数据失败")
		}
		return
	}
	// 返回查询App结果
	var reqAppList []AppBase
	BindReqAndM(&appList, &reqAppList)
	WrapContext(ctx).Success(&reqAppList)
}

func AddApp(ctx *gin.Context) {
	// 获取、绑定请求参数
	reqApp := AppBase{}
	err := ctx.ShouldBindJSON(&reqApp)
	if err != nil {
		WrapContext(ctx).Error(http.StatusInternalServerError, "绑定参数失败")
		return
	}
	// 根据请求参数specify_app_code，查询specify_app_name
	var pips []Pipeline
	result := models.DB1.Raw("SELECT pipeline_code, pipeline_name FROM devops_pipeline_main WHERE dr = 0 AND pipeline_code = ? LIMIT 10", reqApp.SpecifyAppCode).Scan(&pips)
	if result.Error != nil {
		switch result.Error {
		case gorm.ErrRecordNotFound:
			WrapContext(ctx).Error(http.StatusInternalServerError, "查询流水线数据失败")
		}
		return
	}
	pip := pips[0]
	// 新增App
	app := models.App{}
	app.SpecifyAppName = pip.PipelineName
	app.SpecifyAppCode = reqApp.SpecifyAppCode
	result = models.DB.Create(&app)
	if result.Error != nil {
		switch result.Error {
		case gorm.ErrRecordNotFound:
			WrapContext(ctx).Error(http.StatusInternalServerError, "新增微服务失败")
		case gorm.ErrDuplicatedKey:
			WrapContext(ctx).Error(http.StatusInternalServerError, "新增微服务失败，微服务已存在")
		default:
			WrapContext(ctx).Error(http.StatusInternalServerError, "新增微服务失败")
		}
		return

	}
	// 返回新增App
	BindReqAndM(&app, &reqApp)
	WrapContext(ctx).Success(&reqApp)

}
