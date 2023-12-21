package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	StatusCode int         `json:"status_code" ` // 业务状态码
	Message    string      `json:"message" `     // 提示信息
	Data       interface{} `json:"data" `        // 任何数据
	// Meta       Meta        `json:"meta" `        // 源数据，存储比如请求ID、分页信息等
	Errors []ErrorItem `json:"errors" ` // 错误提示，比如xx字段不能为空等
}

type Meta struct {
	RequestID string `json:"request_id" `
	Page      int    `json:"page" `
}

type ErrorItem struct {
	Key   string `json:"key" `
	Value string `json:"value" `
}

func NewResponse() *Response {
	return &Response{
		StatusCode: 200,
		Message:    "success",
		Data:       nil,
		// Meta: Meta{
		// 	RequestID: "1234", // 可以是uuid
		// 	Page:      1,
		// },
		Errors: []ErrorItem{},
	}
}

// Wrapper 封装了gin.Context
type Wrapper struct {
	ctx *gin.Context
}

func WrapContext(ctx *gin.Context) *Wrapper {
	return &Wrapper{ctx: ctx}
}

// Success 输出成功信息
func (wrapper *Wrapper) Success(data interface{}) {
	resp := NewResponse()
	resp.Data = data
	wrapper.ctx.JSON(http.StatusOK, resp)
}

// Error 输出错误信息
func (wrapper *Wrapper) Error(statusCode int, errMessage string) {
	resp := NewResponse()
	resp.StatusCode = statusCode
	resp.Message = errMessage
	wrapper.ctx.JSON(statusCode, resp)
}
