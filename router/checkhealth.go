package router

import (
	"antifraud/controllers"
	"antifraud/pkg/consterrors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CheckHealth
// @Description 探针路由
// @Summary 反欺诈系统探针接口
// @Router /checkhealth [GET]
// @Success 200 {object} controllers.Resp
func CheckHealth(c *gin.Context) {
	r := controllers.Resp{
		Code: consterrors.Success,
		Msg:  fmt.Sprintf("%s state health...", "antifraud"),
		Data: []int{},
	}
	c.JSON(http.StatusOK, r)
}
