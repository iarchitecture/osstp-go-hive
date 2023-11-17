package core_captcha

import (
	"image/color"

	"github.com/mojocn/base64Captcha"
)

// 设置自带的store
var store = base64Captcha.DefaultMemStore

type CaptchaResult struct {
	Id         string `json:"id"`
	Base64Blob string `json:"base_64_blob"`
	Err        error  `json:"err"`
}

func DriverDigitHandle() (captchaResult CaptchaResult) {
	driver := base64Captcha.NewDriverDigit(80, 240, 4, 0.7, 80)
	captcha := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := captcha.Generate()

	result := CaptchaResult{
		Id:         id,
		Base64Blob: b64s,
		Err:        err,
	}

	return result
}

func DriverStringHandle() (captchaResult CaptchaResult) {
	driverString := base64Captcha.NewDriverString(46, 140, 2, 2, 4, "234567890abcdefghjkmnpqrstuvwxyz", &color.RGBA{240, 240, 246, 246}, nil, []string{"wqy-microhei.ttc"})
	driver := driverString.ConvertFonts()
	captcha := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := captcha.Generate()

	result := CaptchaResult{
		Id:         id,
		Base64Blob: b64s,
		Err:        err,
	}

	return result
}

// Verify 校验验证码
func DriverVerify(id, code string, clear bool) bool {
	return store.Verify(id, code, clear)
}
