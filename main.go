package main

import (
	"go-gin-jwt/models"
	"go-gin-jwt/routes"
)

func init() {
	// 初始化DB连接
	models.ConnectDB()
}

func main() {
	// 启动路由
	r := routes.Routes()
	r.Run("0.0.0.0:8000")
}
