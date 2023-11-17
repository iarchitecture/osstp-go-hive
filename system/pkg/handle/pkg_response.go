package pkg_handle

import (
	pkg_constant "osstp-go-hive/system/pkg/constant"

	"github.com/gin-gonic/gin"
)

type Ctx struct {
	Context *gin.Context
}

type Response struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

type ResponsePage struct {
	Status    int         `json:"status"`
	Msg       string      `json:"msg"`
	Data      interface{} `json:"data"`
	Count     int         `json:"count"`
	PageIndex int         `json:"page_index"`
	PageSize  int         `json:"page_size"`
}

func (ctx *Ctx) Response(httpCode int, status int, data interface{}) {
	ctx.Context.JSON(httpCode, Response{
		Status: status,
		Msg:    pkg_constant.StatusMsg(status),
		Data:   data,
	})
	return
}

func (ctx *Ctx) ResponsePage(httpCode int, status interface{}, data interface{}, count, pageIndex, pageSize int) {
	sta := status.(int)
	ctx.Context.JSON(httpCode, ResponsePage{
		Status:    sta,
		Msg:       pkg_constant.StatusMsg(sta),
		Data:      data,
		Count:     count,
		PageIndex: pageIndex,
		PageSize:  pageSize,
	})
	return
}
