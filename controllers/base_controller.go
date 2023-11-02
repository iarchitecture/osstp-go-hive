package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type JsonSuccessStruct struct {
	Code  int         `json:"code"`
	Msg   interface{} `json:"msg"`
	Data  interface{} `json:"data"`
	Count int64       `json:"count"`
}
type JsonErrStruct struct {
	Code int         `json:"code"`
	Msg  interface{} `json:"msg"`
}

func SuccessResponse(ctx *gin.Context, code int, msg interface{}, data interface{}, count int64) {
	json := &JsonSuccessStruct{Code: code, Msg: msg, Data: data, Count: count}

	ctx.JSON(http.StatusOK, json)
}

// ErrorResponse 返回错误信息
func ErrorResponse(ctx *gin.Context, code int, msg interface{}) {
	json := &JsonErrStruct{Code: code, Msg: msg}

	ctx.JSON(http.StatusOK, json)
}
