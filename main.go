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
	routes.Route()
}
