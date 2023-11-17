package pkg_constant

var MsgMap = map[int]string{
	Success:             "success",
	Fail:                "fail",
	Invalid_params:      "请求参数异常",
	Error_database_CRUD: "数据库CRUD失败",

	Error_user_exist:     "该用户已存在",
	Error_user_not_exist: "该用户不存在",
	Error_password:       "密码错误",
	Error_captcha:        "验证码获取失败",
	Error_logout:         "退出失败",

	// ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	// ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
	// ERROR_AUTH_TOKEN:               "Token生成失败",
	// ERROR_AUTH:                     "Token错误",
	// ERROR_AUTH_CHECK_FAIL:          "无权限，请联系管理员",

	// ERROR_UPLOAD_SAVE_IMAGE_FAIL:    "保存图片失败",
	// ERROR_UPLOAD_CHECK_IMAGE_FAIL:   "检查图片失败",
	// ERROR_UPLOAD_CHECK_IMAGE_FORMAT: "校验图片错误，图片格式或大小有问题",
}

func StatusMsg(code int) string {
	msg, ok := MsgMap[code]
	if ok {
		return msg
	}

	return MsgMap[Fail]
}
