package router

import (
	"antifraud/controllers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter(engine *gin.Engine) *gin.Engine {
	// 缺失路由
	engine.NoRoute(NotFoundHandler)

	// 健康检查
	engine.GET("/checkhealth", CheckHealth)

	// swag 文档
	engine.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 反欺诈路由
	c := controllers.AntiFraudController{}
	group := engine.Group("/api/v1")
	group.GET("/events", c.Events)

	return engine
}
