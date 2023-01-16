package errcode

import (
	"fmt"
	"net/http"
)

type Error struct {
	ErrNo   int      `json:"code"`
	Msg     string   `json:"msg"`
	Details []string `json:"details"`
}

// codes map类型
var codes = map[int]string{}

// NewError 实例化错误结构体
func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码 %d 已存在，请更换一个", code))
	}
	codes[code] = msg
	return &Error{ErrNo: code, Msg: msg}
}

// Error 输出错误信息
func (e *Error) Error() string {
	return fmt.Sprintf("错误码：%d, 错误信息:%s", e.Code(), e.ErrMsg())
}

// Code 输出code
func (e *Error) Code() int {
	return e.ErrNo
}

// ErrMsg 输出msg
func (e *Error) ErrMsg() string {
	return e.Msg
}

// Msgf 输出更多的错误信息和详情
func (e *Error) Msgf(args []interface{}) string {
	return fmt.Sprintf(e.Msg, args...)
}

// ErrDetails 错误详情的内容
func (e *Error) ErrDetails() []string {
	return e.Details
}

// WithDetails 将详情拼接起来
func (e *Error) WithDetails(details ...string) *Error {
	newError := *e
	newError.Details = []string{}
	for _, d := range details {
		newError.Details = append(newError.Details, d)
	}
	return &newError
}

// StatusCode 状态码
func (e *Error) StatusCode() int {
	switch e.Code() {
	case Success.Code():
		return http.StatusOK
	case ServerError.Code():
		return http.StatusInternalServerError
	case InvalidParams.Code():
		return http.StatusBadRequest
	case NotFound.Code():
		fallthrough
	case TooManyRequests.Code():
		return http.StatusTooManyRequests
	}
	return http.StatusInternalServerError
}
