package admin

import (
	"osstp-go-hive/system/app/apis/admin"

	"github.com/gin-gonic/gin"
)

func RegisterSystemRouters(engine *gin.Engine) {

	engine.POST("/auth/captcha", admin.CaptchaApi{}.GenerateCaptchaApiHandler)
	engine.POST("/auth/login", admin.LoginApi{}.Login)

	// adminRouter := engine.Group("/admin/auth")
	// adminRouter.Use()
	// {
	// 	adminRouter.POST("/login", admin.LoginApi{}.Login)
	// }

}
