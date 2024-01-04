package rule

import (
	"net/http"
	"strconv"
	"zheng11581/toy-gin/handlers"
	"zheng11581/toy-gin/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Get(ctx *gin.Context) {
	// 获取、绑定参数
	ruleID, err := strconv.ParseUint(ctx.Param("id"), 8, 64)
	if err != nil {
		handlers.WrapContext(ctx).Error(http.StatusInternalServerError, "获取参数失败")
	}
	// 查询 Rule
	rule := models.Rule{}
	rule.ID = uint(ruleID)
	result := models.DB.Find(&rule)
	if result.Error != nil {
		switch result.Error {
		case gorm.ErrRecordNotFound:
			handlers.WrapContext(ctx).Error(http.StatusNotFound, "查询数据失败")

		}
	}
	// 返回查询到的数据
	handlers.WrapContext(ctx).Success(&rule)
}

func Delete(ctx *gin.Context) {
	// 获取、绑定参数
	ruleID, err := strconv.ParseUint(ctx.Param("id"), 8, 64)
	if err != nil {
		handlers.WrapContext(ctx).Error(http.StatusInternalServerError, "获取参数失败")
	}
	// 删除 Rule
	rule := models.Rule{}
	rule.ID = uint(ruleID)
	result := models.DB.Delete(&rule)
	if result.Error != nil {
		switch result.Error {
		case gorm.ErrRecordNotFound:
			handlers.WrapContext(ctx).Error(http.StatusNotFound, "删除数据失败")

		}
	}
	// 返回被删除的数据
	handlers.WrapContext(ctx).Success(&rule)
}

func List(ctx *gin.Context) {
	// 获取、绑定参数
	reqRule := handlers.RuleBase{}
	err := ctx.ShouldBindJSON(&reqRule)
	if err != nil {
		handlers.WrapContext(ctx).Error(http.StatusInternalServerError, "绑定参数失败")
	}
	// 查询 Rule 列表
	ruleList := []models.Rule{}
	result := models.DB.Where("specify_app_code = ?", reqRule.SpecifyAppCode).Find(&ruleList)
	if result != nil {
		switch result.Error {
		case gorm.ErrRecordNotFound:
			handlers.WrapContext(ctx).Error(http.StatusNotFound, "查询数据失败")
		}
	}
	// 返回查询到的结果
	handlers.WrapContext(ctx).Success(&ruleList)
}

func Add(ctx *gin.Context) {
	// 获取、绑定参数
	reqRule := handlers.RuleBase{}
	err := ctx.ShouldBindJSON(&reqRule)
	if err != nil {
		handlers.WrapContext(ctx).Error(http.StatusInternalServerError, "绑定参数失败")
	}
	// 新增 Rule
	rule := models.Rule{}
	err = handlers.BindReqToM(&reqRule, &rule)
	if err != nil {
		handlers.WrapContext(ctx).Error(http.StatusInternalServerError, "新增数据失败")
	}
	result := models.DB.Create(&rule)
	if result != nil {
		switch result.Error {
		case gorm.ErrRecordNotFound:
			handlers.WrapContext(ctx).Error(http.StatusNotFound, "新增数据失败")
		}
	}
	// 返回新增的数据
	handlers.WrapContext(ctx).Success(&rule)

}

func Update(ctx *gin.Context) {
	// 获取、绑定参数
	ruleID, err := strconv.ParseUint(ctx.Param("id"), 8, 64)
	if err != nil {
		handlers.WrapContext(ctx).Error(http.StatusInternalServerError, "获取参数失败")
	}
	reqRule := handlers.RuleBase{}
	err = ctx.ShouldBindJSON(&reqRule)
	if err != nil {
		handlers.WrapContext(ctx).Error(http.StatusInternalServerError, "绑定参数失败")
	}
	// 更新 Rule
	rule := models.Rule{}
	err = handlers.BindReqToM(&reqRule, &rule)
	if err != nil {
		handlers.WrapContext(ctx).Error(http.StatusInternalServerError, "更新数据失败")
	}
	rule.ID = uint(ruleID)
	models.DB.Updates(&rule)
	// 返回更新的数据
	handlers.WrapContext(ctx).Success(&rule)
}
