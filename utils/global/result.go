package global

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Builder 统一返回结构体
type Builder struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type Result struct {
	Ctx *gin.Context
}

func NewSysError(c *gin.Context) {
	NewResult().SetCode(500).SetMsg("System busy, please try again later!").Build(c)
}

func NewResult() *Builder {
	return &Builder{
		Code: 200,
		Msg:  "Success",
		Data: nil,
	}
}

func (b *Builder) SetCode(code int) *Builder {
	b.Code = code
	return b
}

func (b *Builder) SetMsg(msg string) *Builder {
	b.Msg = msg
	return b
}

func (b *Builder) SetData(data interface{}) *Builder {
	b.Data = data
	return b
}

func (b *Builder) Build(ctx *gin.Context) {
	//ctx.JSON(http.StatusOK, b)
	//ctx.Abort()
	if b.Data == nil { // data 为空时，将 null => {} 空对象
		b.Data = gin.H{}
	}
	ctx.AbortWithStatusJSON(http.StatusOK, b)
}

// ReturnJson 返回JSON
//func ReturnJson(ctx *gin.Context) *Result {
//	return &Result{Ctx: ctx}
//}

// Success 成功
//func (r *Result) Success(msg string, data interface{}) {
//	if data == nil {
//		data = gin.H{}
//	}
//	res := ResultContent{}
//	res.Code = 200
//	res.Msg = msg
//	res.Data = data
//	r.Ctx.JSON(http.StatusOK, res)
//	r.Ctx.Abort()
//}
