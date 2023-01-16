package response

import (
	"github.com/gin-gonic/gin"
	"{{.ProjectName}}/pkg/errcode"
)

// BaseResponse 基础响应
type BaseResponse struct {
	ErrNo int `json:"errNo"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}

// Success 成功的响应
func (b *BaseResponse) Success(c *gin.Context, data interface{}) {
	res := &BaseResponse {
		ErrNo: errcode.Success.Code(),
		Msg: errcode.Success.ErrMsg(),
		Data: data,
	}
	c.JSON(errcode.Success.StatusCode(), res)
}

// Failure 公共失败的响应，已经错误了，没有必要再传data了
// 传入一个errcode.Error，那么在调用时，支持传入common_err里提前设置好的错误信息
// 这样就可以获取到code和msg
// 并且响应里也根据传入进来的Error对象来调用StatusCode，进一步获取到定义的状态码
func (b *BaseResponse) Failure(c *gin.Context, err *errcode.Error) {
	res := &BaseResponse {
		ErrNo: err.Code(),
		Msg: err.ErrMsg(),
		Data: nil,
	}
	c.JSON(err.StatusCode(), res)
}

// FailureWithMsg 返回响应，但是是自定义的错误code和错误信息，不是公共的错误和信息
// 走到自定义错误，其实还是可以返回正常的
func (b *BaseResponse) FailureWithMsg(c *gin.Context, errNo int, msg string) {
	res := &BaseResponse {
		ErrNo: errNo,
		Msg: msg,
		Data: nil,
	}
	c.JSON(errcode.Success.StatusCode(), res)
}

