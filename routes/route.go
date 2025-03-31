package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin-jwt/controllers"
	"go-gin-jwt/middlewares"
	"go-gin-jwt/utils/global"
	"runtime/debug"
	"time"
)

func Routes() *gin.Engine {
	r := gin.Default()

	//处理找不到路由
	r.NoRoute(HandleNotFound)
	r.NoMethod(HandleNotFound)

	// 处理发生异常
	r.Use(Recover)

	// public路由
	public := r.Group("/api/auth")
	{
		public.POST("/register", controllers.Register)
		public.POST("/login", controllers.Login)
	}
	// protected路由
	protected := r.Group("/api")
	{
		protected.Use(middlewares.JwtAuthMiddleware())
		protected.GET("/user/info", controllers.CurrentUser)
	}

	test := r.Group("/api/test")
	{
		test.GET("/", func(c *gin.Context) {
			global.NewSysError(c)
			global.NewResult().SetMsg("text api1").Build(c)
		})
	}
	return r
}

// HandleNotFound 404 找不到路径时的处理
func HandleNotFound(c *gin.Context) {
	c.JSON(404, gin.H{"code": 500, "msg": "Page not found"})
}

// Recover 500 内部发生异常时的处理
func Recover(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			//打印错误堆栈信息
			timeStr := time.Now().Format("2006-01-02 15:04:05")
			fmt.Println("当前时间:", timeStr)
			fmt.Println("当前访问path:", c.FullPath())
			fmt.Println("当前完整地址:", c.Request.URL.String())
			fmt.Println("当前协议:", c.Request.Proto)
			//fmt.Println("当前get参数:", global.GetAllGetParams(c))
			//fmt.Println("当前post参数:", global.GetAllPostParams(c))
			fmt.Println("当前访问方法:", c.Request.Method)
			fmt.Println("当前访问Host:", c.Request.Host)
			fmt.Println("当前IP:", c.ClientIP())
			fmt.Println("当前浏览器:", c.Request.UserAgent())
			fmt.Println("发生异常:", err)
			//global.Logger.Errorf("stack: %v",string(debug.Stack()))
			debug.PrintStack()
			//return
			c.JSON(200, gin.H{"code": 500, "msg": "System Error"})
		}
	}()
	//继续后续接口调用
	c.Next()
}
