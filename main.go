package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tdytaylor/go-gin-web-all/middleware"
)

func main() {
	web := gin.Default()
	web.Use(middleware.Auth())
	// 监听并在 0.0.0.0:8080 上启动服务
	web.Run(":8080")
}
