package main

import (
	v1 "github.com/zhystc888/glorious-version-repository/internal/api/v1"
	"github.com/zhystc888/glorious-version-repository/internal/config"
)

func main() {
	cfg := config.Load()     // 加载配置
	r := v1.SetupRouter(cfg) // 初始化路由
	r.Run(cfg.Port)          // 启动服务
}
