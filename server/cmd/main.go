package main

import (
	"fmt"
	"log"
	"tennis-booking-system/api"
	"tennis-booking-system/internal/config"

	"github.com/gin-gonic/gin"
)

func main() {
	// 加载配置
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 初始化数据库
	if err := config.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// 设置Gin模式
	gin.SetMode(cfg.Server.Mode)

	// 创建Gin引擎
	r := gin.Default()

	// 注册路由
	api.SetupRoutes(r)

	// 启动服务器
	addr := fmt.Sprintf(":%s", cfg.Server.Port)
	log.Printf("Server starting on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
