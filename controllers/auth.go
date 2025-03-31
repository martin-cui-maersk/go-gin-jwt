package controllers

import (
	"go-gin-jwt/models"
	"go-gin-jwt/utils/global"
	"go-gin-jwt/utils/token"

	"github.com/gin-gonic/gin"
)

// ReqRegister /api/register的请求体
type ReqRegister struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// ReqLogin api/login 的请求体
type ReqLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	var req ReqRegister

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		//c.JSON(http.StatusBadRequest, gin.H{
		//	"data": err.Error(),
		//})
		global.NewResult().SetCode(503).SetMsg(err.Error()).SetData(nil).Build(c)
		return
	}

	u := models.User{
		Username: req.Username,
		Password: req.Password,
	}

	_, err := u.SaveUser()
	if err != nil {
		//c.JSON(http.StatusBadRequest, gin.H{
		//	"data": err.Error(),
		//})
		global.NewResult().SetCode(503).SetMsg(err.Error()).SetData(nil).Build(c)
		return
	}
	//c.JSON(http.StatusOK, gin.H{
	//	"message": "register success",
	//	"data":    req,
	//})
	global.NewResult().SetCode(200).SetMsg("register success").SetData(req).Build(c)
}

func Login(c *gin.Context) {
	var req ReqLogin
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		// c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		global.NewResult().SetCode(503).SetMsg(err.Error()).Build(c)
		return
	}

	u := models.User{
		Username: req.Username,
		Password: req.Password,
	}
	// 调用 models.LoginCheck 对用户名和密码进行验证
	token, err := models.LoginCheck(u.Username, u.Password)
	if err != nil {
		//c.JSON(http.StatusBadRequest, gin.H{
		//	"error": "username or password is incorrect.",
		//})
		// global.NewResult().SetCode(503).SetMsg("username or password is incorrect.").Build(c)
		global.NewResult().SetCode(503).SetMsg(err.Error()).Build(c)
		return
	}
	//c.JSON(http.StatusOK, gin.H{
	//	"token": token,
	//})
	global.NewResult().SetData(map[string]string{"token": token}).Build(c)
}

func CurrentUser(c *gin.Context) {
	// 从token中解析出user_id
	userId, err := token.ExtractTokenID(c)
	if err != nil {
		//c.JSON(http.StatusBadRequest, gin.H{
		//	"error": err.Error(),
		//})
		global.NewResult().SetCode(503).SetMsg(err.Error()).SetData(nil).Build(c)
		return
	}

	// 根据user_id从数据库查询数据
	u, err := models.GetUserByID(userId)
	if err != nil {
		//c.JSON(http.StatusBadRequest, gin.H{
		//	"error": err.Error(),
		//})
		global.NewResult().SetCode(503).SetMsg(err.Error()).SetData(nil).Build(c)
		return
	}

	//c.JSON(http.StatusOK, gin.H{
	//	"message": "success",
	//	"data":    u,
	//})
	global.NewResult().SetData(u).Build(c)
}
