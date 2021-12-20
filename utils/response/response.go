package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func Result(code int, data interface{}, msg string, context *gin.Context) {
	context.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func Ok(context *gin.Context) {
	Result(SUCCESS, nil, GetErrMsg(SUCCESS), context)
}
func Fail(context *gin.Context) {
	Result(ERROR, nil, GetErrMsg(ERROR), context)
}

func OkWithData(data interface{}, context *gin.Context) {
	Result(SUCCESS, data, GetErrMsg(SUCCESS), context)
}

func OkWithMsg(msg string, context *gin.Context) {
	Result(SUCCESS, nil, msg, context)
}

func OkWithDataMsg(data interface{}, msg string, context *gin.Context) {
	Result(SUCCESS, data, msg, context)
}

func FailWithData(data interface{}, context *gin.Context) {
	Result(ERROR, data, GetErrMsg(ERROR), context)
}

func FailWithMsg(msg string, context *gin.Context) {
	Result(ERROR, nil, msg, context)
}

func FailWithDataMsg(data interface{}, msg string, context *gin.Context) {
	Result(ERROR, data, msg, context)
}
func FailWithCode(code int, context *gin.Context) {
	Result(code, nil, GetErrMsg(code), context)
}

func ResultWithCode(code int, context *gin.Context) {
	Result(code, nil, GetErrMsg(code), context)
}
