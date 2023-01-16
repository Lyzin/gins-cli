package errcode

var (
	Success                   = NewError(0, "成功")
	ServerError               = NewError(50000, "服务内部错误")
	InvalidParams             = NewError(50001, "参数错误")
	NotFound                  = NewError(50002, "未找到")
	TooManyRequests           = NewError(50007, "请求过多")
)
