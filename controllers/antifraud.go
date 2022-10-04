package controllers

import "github.com/gin-gonic/gin"

type AntiFraudController struct {
	Base
}

func (m *AntiFraudController) Events(c *gin.Context) {
	m.Success(c, "", []string{"1", "2"})
}
