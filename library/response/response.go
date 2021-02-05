package response

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"net/http"
)

var (
	OK           = http.StatusOK
	OkMessage    = "正常"
	ERROR        = http.StatusInternalServerError
	ErrorMessage = "服务器发生了一些错误"
)

type ResultRes struct {
	Code    int         `json:"code"` // response code, default: success-200 error-500 响应码
	Message string      `json:"msg"`  // response message 响应信息，若返回的是错误码，则此处对应相关错误信息
	Data    interface{} `json:"data"` // response result data 返回数据
}

// RestResult to write response, but not exit. 返回统一格式对象，但不退出。
func RestResult(r *ghttp.Request, code int, msg string, data interface{}) {
	if err := r.Response.WriteJson(ResultRes{
		Code:    code,
		Message: msg,
		Data:    data,
	}); err != nil {
		// response write error, print log. 返回错误，记录日志
		g.Log().Error(err)
	}
}

// RestResultExit to write response and exit. 返回统一格式对象，并退出。
func RestResultExit(r *ghttp.Request, code int, msg string, data interface{}) {
	RestResult(r, code, msg, data)
	r.Exit()
}

// Success result default success message, but not exit. 返回默认成功信息，但不退出。
func Success(r *ghttp.Request) {
	RestResult(r, OK, OkMessage, nil)
}

// SuccessExit result default success message and exit. 返回默认成功信息，并退出。
func SuccessExit(r *ghttp.Request) {
	RestResultExit(r, OK, OkMessage, nil)
}

// SuccessMsg result success and custom message but not exit. 返回成功和自定义信息，但不退出。
func SuccessMsg(r *ghttp.Request, msg string) {
	RestResult(r, OK, msg, nil)
}

// SuccessMsgExit result success and custom message and exit. 返回成功和自定义信息，并退出。
func SuccessMsgExit(r *ghttp.Request, msg string) {
	RestResultExit(r, OK, msg, nil)
}

// SuccessData result success and custom data but not exit. 返回成功和数据，但不退出。
func SuccessData(r *ghttp.Request, data interface{}) {
	RestResult(r, OK, OkMessage, data)
}

// SuccessDataExit result success and custom data and exit. 返回成功和数据，并退出。
func SuccessDataExit(r *ghttp.Request, data interface{}) {
	RestResultExit(r, OK, OkMessage, data)
}

// Fail result default fail message, but not exit. 返回默认失败信息，但不退出。
func Fail(r *ghttp.Request) {
	RestResult(r, ERROR, ErrorMessage, nil)
}

// FailExit result default fail message and exit. 返回默认失败信息，并退出。
func FailExit(r *ghttp.Request) {
	RestResultExit(r, ERROR, ErrorMessage, nil)
}

// FailMsg result error and custom message but not exit. 返回失败和信息，但不退出
func FailMsg(r *ghttp.Request, msg string) {
	RestResult(r, ERROR, msg, nil)
}

// FailMsgExit result error and custom message and exit. 返回失败和信息，并退出
func FailMsgExit(r *ghttp.Request, msg string) {
	RestResultExit(r, ERROR, msg, nil)
}
