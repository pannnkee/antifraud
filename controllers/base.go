package controllers

import (
	"antifraud/pkg/consterrors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type (
	// Resp 响应
	Resp struct {
		Code int         `json:"code"` // 应答码
		Msg  string      `json:"msg"`  // 应答信息
		Data interface{} `json:"data"` // 应答数据
	}

	// RespPage 分页相应
	RespPage struct {
		Resp
		Total    int64 `json:"total"`     // 分页总数
		Page     int64 `json:"page"`      // 当前页
		PageSize int64 `json:"page_size"` // 每页数量
	}

	Base struct{}
)

func (m *Base) Success(c *gin.Context, message string, data interface{}) {
	if message == "" {
		message = consterrors.ResponseMap[consterrors.Success]
	}
	r := Resp{
		Code: consterrors.Success,
		Msg:  message,
		Data: data,
	}
	c.JSON(http.StatusOK, r)
}
