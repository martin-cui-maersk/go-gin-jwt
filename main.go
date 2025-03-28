package main

import (
	"github.com/gin-gonic/gin"
	"go-gin-jwt/controllers"
	"go-gin-jwt/middlewares"
	"go-gin-jwt/models"
)

func init() {
	models.ConnectDB()
}

func main() {
	r := gin.Default()
	public := r.Group("/api")
	{
		public.POST("/register", controllers.Register)
		public.POST("/login", controllers.Login)
	}
	protected := r.Group("/api/admin")
	{
		protected.Use(middlewares.JwtAuthMiddleware())
		protected.GET("/user", controllers.CurrentUser)
	}
	r.Run("0.0.0.0:8000")
}
