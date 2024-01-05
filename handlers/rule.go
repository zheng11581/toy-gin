package handlers

import (
	"net/http"
	"strconv"
	"zheng11581/toy-gin/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRule(ctx *gin.Context) {
	// 获取、绑定参数
	ruleID, err := strconv.ParseUint(ctx.Param("id"), 8, 64)
	if err != nil {
		WrapContext(ctx).Error(http.StatusInternalServerError, "获取参数失败")
		return
	}
	// 查询 Rule
	rule := models.Rule{}
	rule.ID = uint(ruleID)
	result := models.DB.Find(&rule)
	if result.Error != nil {
		switch result.Error {
		case gorm.ErrRecordNotFound:
			WrapContext(ctx).Error(http.StatusNotFound, "查询数据失败")
		}
		return
	}
	// 返回查询到的数据
	reqRule := RuleBase{}
	BindReqAndM(&rule, &reqRule)
	WrapContext(ctx).Success(&reqRule)
}

func DeleteRule(ctx *gin.Context) {
	// 获取、绑定参数
	ruleID, err := strconv.ParseUint(ctx.Param("id"), 8, 64)
	if err != nil {
		WrapContext(ctx).Error(http.StatusInternalServerError, "获取参数失败")
		return
	}
	// 删除 Rule
	rule := models.Rule{}
	rule.ID = uint(ruleID)
	result := models.DB.Delete(&rule)
	if result.Error != nil {
		switch result.Error {
		case gorm.ErrRecordNotFound:
			WrapContext(ctx).Error(http.StatusNotFound, "删除数据失败")

		}
		return
	}
	// 返回被删除的数据
	reqRule := RuleBase{}
	BindReqAndM(&rule, &reqRule)
	WrapContext(ctx).Success(&reqRule)
}

func ListRules(ctx *gin.Context) {
	// 获取、绑定参数
	reqRule := RuleBase{}
	err := ctx.ShouldBindQuery(&reqRule)
	if err != nil {
		WrapContext(ctx).Error(http.StatusInternalServerError, "绑定参数失败")
		return
	}
	// 查询 Rule 列表
	ruleList := []models.Rule{}
	result := models.DB.Where("specify_app_code = ?", reqRule.SpecifyAppCode).Find(&ruleList)
	if result.Error != nil {
		switch result.Error {
		case gorm.ErrRecordNotFound:
			WrapContext(ctx).Error(http.StatusNotFound, "查询数据失败")
		}
		return
	}
	// 返回查询到的结果
	var reqRuleList []RuleBase
	BindReqAndM(&ruleList, &reqRuleList)
	WrapContext(ctx).Success(&reqRuleList)
}

func AddRule(ctx *gin.Context) {
	// 获取、绑定参数
	reqRule := RuleBase{}
	err := ctx.ShouldBindJSON(&reqRule)
	if err != nil {
		WrapContext(ctx).Error(http.StatusInternalServerError, "绑定参数失败")
		return
	}
	// 新增 Rule
	rule := models.Rule{}
	err = BindReqAndM(&reqRule, &rule)
	if err != nil {
		WrapContext(ctx).Error(http.StatusInternalServerError, "新增数据失败")
		return
	}
	result := models.DB.Create(&rule)
	if result.Error != nil {
		switch result.Error {
		case gorm.ErrRecordNotFound:
			WrapContext(ctx).Error(http.StatusNotFound, "新增数据失败")
		}
		return
	}
	// 返回新增的数据
	BindReqAndM(&rule, &reqRule)
	WrapContext(ctx).Success(&reqRule)

}

func UpdateRule(ctx *gin.Context) {
	// 获取、绑定参数
	ruleID, err := strconv.ParseUint(ctx.Param("id"), 8, 64)
	if err != nil {
		WrapContext(ctx).Error(http.StatusInternalServerError, "获取参数失败")
		return
	}
	reqRule := RuleBase{}
	err = ctx.ShouldBindJSON(&reqRule)
	if err != nil {
		WrapContext(ctx).Error(http.StatusInternalServerError, "绑定参数失败")
		return
	}
	// 更新 Rule
	rule := models.Rule{}
	err = BindReqAndM(&reqRule, &rule)
	if err != nil {
		WrapContext(ctx).Error(http.StatusInternalServerError, "更新数据失败")
		return
	}
	rule.ID = uint(ruleID)
	result := models.DB.Updates(&rule)
	if result.Error != nil {
		switch result.Error {
		case gorm.ErrRecordNotFound:
			WrapContext(ctx).Error(http.StatusNotFound, "更新数据失败")
		}
		return
	}
	// 返回更新的数据
	BindReqAndM(&rule, &reqRule)
	WrapContext(ctx).Success(&reqRule)
}

func GetSpecialRule(ctx *gin.Context) {
	// 获取、绑定请求参数
	specID, err := strconv.ParseUint(ctx.Param("id"), 8, 64)
	if err != nil {
		WrapContext(ctx).Error(http.StatusInternalServerError, "获取参数失败")
		return
	}
	// 查询SpecialRule数据
	specRule := models.SpecialRule{}
	specRule.ID = uint(specID)
	result := models.DB.Find(&specRule)
	if result.Error != nil {
		switch result.Error {
		case gorm.ErrRecordNotFound:
			WrapContext(ctx).Error(http.StatusNotFound, "查询数据失败")
		}
		return
	}
	// 返回响应数据
	reqSpecRule := SpecialRule{}
	BindReqAndM(&specRule, &reqSpecRule)
	WrapContext(ctx).Success(&reqSpecRule)

}

func DeleteSpecialRule(ctx *gin.Context) {
	// 获取、绑定请求参数
	specID, err := strconv.ParseUint(ctx.Param("id"), 8, 64)
	if err != nil {
		WrapContext(ctx).Error(http.StatusInternalServerError, "获取参数失败")
		return
	}
	// 删除SpecialRule数据
	specRule := models.SpecialRule{}
	specRule.ID = uint(specID)
	result := models.DB.Delete(&specRule)
	if result.Error != nil {
		switch result.Error {
		case gorm.ErrRecordNotFound:
			WrapContext(ctx).Error(http.StatusNotFound, "查询数据失败")
		}
		return
	}
	// 返回响应数据
	reqSpecRule := SpecialRule{}
	BindReqAndM(&specRule, &reqSpecRule)
	WrapContext(ctx).Success(&reqSpecRule)
}

func ListSpecialRule(ctx *gin.Context) {
	// 获取、绑定请求参数
	reqSpecRule := SpecialRule{}
	err := ctx.ShouldBindQuery(&reqSpecRule)
	if err != nil {
		WrapContext(ctx).Error(http.StatusInternalServerError, "绑定参数失败")
		return
	}
	// 查询SpecialRule列表
	specRuleList := []models.SpecialRule{}
	result := models.DB.Where("rule_id = ?", reqSpecRule.RuleID).Find(&specRuleList)
	if result.Error != nil {
		switch result.Error {
		case gorm.ErrRecordNotFound:
			WrapContext(ctx).Error(http.StatusNotFound, "查询数据失败")
		}
		return
	}
	// 返回响应数据
	var reqSpecRuleList []SpecialRule
	BindReqAndM(&specRuleList, &reqSpecRuleList)
	WrapContext(ctx).Success(&reqSpecRuleList)
}

func AddSpecialRule(ctx *gin.Context) {
	// 获取、绑定请求参数
	reqSpecRule := SpecialRule{}
	err := ctx.ShouldBindJSON(&reqSpecRule)
	if err != nil {
		WrapContext(ctx).Error(http.StatusInternalServerError, "绑定参数失败")
		return
	}
	// 新增SpecialRule数据
	specRule := models.SpecialRule{}
	BindReqAndM(&reqSpecRule, &specRule)
	result := models.DB.Create(&specRule)
	if result.Error != nil {
		switch result.Error {
		case gorm.ErrRecordNotFound:
			WrapContext(ctx).Error(http.StatusNotFound, "新增数据失败")
		}
		return
	}
	// 返回响应数据
	BindReqAndM(&specRule, &reqSpecRule)
	WrapContext(ctx).Success(&reqSpecRule)
}

func UpdateSpecialRule(ctx *gin.Context) {
	// 获取、绑定请求参数
	specRuleID, err := strconv.ParseUint(ctx.Param("id"), 8, 64)
	if err != nil {
		WrapContext(ctx).Error(http.StatusInternalServerError, "获取参数失败")
		return
	}
	reqSpecRule := SpecialRule{}
	err = ctx.ShouldBindJSON(&reqSpecRule)
	if err != nil {
		WrapContext(ctx).Error(http.StatusInternalServerError, "绑定参数失败")
		return
	}
	// 更新SpecialRule数据
	specRule := models.SpecialRule{}
	specRule.ID = uint(specRuleID)
	result := models.DB.Updates(&specRule)
	if result.Error != nil {
		switch result.Error {
		case gorm.ErrRecordNotFound:
			WrapContext(ctx).Error(http.StatusNotFound, "更新数据失败")
		}
		return
	}
	// 返回响应数据
	BindReqAndM(&specRule, &reqSpecRule)
	WrapContext(ctx).Success(&reqSpecRule)
}
