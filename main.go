package main

import (
	"antifraud/cmd/AntiFraudServer"
	_ "antifraud/docs"
	"antifraud/pkg/serverman"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 设置随机种子
	rand.Seed(time.Now().UnixNano())

	// 加载配置

	// 项目启动
	serverman.RegisterServer(AntiFraudServer.New())
	if err := serverman.Start(); err != nil {
		fmt.Println(err)
	}

}
