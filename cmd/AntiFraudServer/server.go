package AntiFraudServer

import (
	"antifraud/pkg/middleware"
	"antifraud/router"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// AntiFraudService 反欺诈服务
type AntiFraudService struct {
	server *http.Server
}

func New() *AntiFraudService {
	return &AntiFraudService{}
}

// Serve 启动服务
func (m *AntiFraudService) Serve() (err error) {
	engine := gin.Default()
	// 跨域
	engine.Use(middleware.Cors())
	// 捕获异常
	engine.Use(middleware.CatchError())
	// 设置路由
	router.InitRouter(engine)

	m.server = &http.Server{
		Addr:         ":8080",
		Handler:      engine,
		ReadTimeout:  time.Duration(60) * time.Second,
		WriteTimeout: time.Duration(60) * time.Second,
	}

	if err = m.server.ListenAndServe(); err != nil {
		return err
	}
	return
}

// Stop 退出
func (m *AntiFraudService) Stop() (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return m.server.Shutdown(ctx)
}
