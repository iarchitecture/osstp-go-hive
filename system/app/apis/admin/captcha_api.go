package admin

import (
	"net/http"
	core_captcha "osstp-go-hive/go-core/pkg/captcha"
	pkg_constant "osstp-go-hive/system/pkg/constant"
	pkg_handle "osstp-go-hive/system/pkg/handle"

	"github.com/gin-gonic/gin"
)

type CaptchaApi struct{}

// GenerateCaptchaApiHandler 获取验证码
//
//	@Summary		获取验证码
//	@Description	获取验证码
//	@Tags			登录
//	@Success		200	{object}	response.Response{data=string,id=string,msg=string}	"{"code": 200, "data": [...]}"
//	@Router			/auth/captcha [get]
func (e CaptchaApi) GenerateCaptchaApiHandler(ctx *gin.Context) {
	c := pkg_handle.Ctx{Context: ctx}
	result := core_captcha.DriverStringHandle()
	if result.Err != nil {
		// log.Logger.Error("DriverDigitFunc error, %s", err.Error())
		c.Response(http.StatusInternalServerError, pkg_constant.Error_captcha, result)
	}
	c.Response(http.StatusOK, pkg_constant.Success, result)
}
