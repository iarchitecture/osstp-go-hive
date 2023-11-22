package admin

import (
	"osstp-go-hive/system/app/admin/api"

	"github.com/gin-gonic/gin"
)

func RegisterSystemRouters(engine *gin.Engine) {

	engine.GET("/auth/captcha", api.CaptchaApi{}.GenerateCaptchaApiHandler)
	engine.POST("/auth/register", api.UserApi{}.Register)
	engine.POST("/auth/login", api.UserApi{}.Login)

	// adminRouter := engine.Group("/admin/auth")
	// adminRouter.Use()
	// {
	// 	adminRouter.POST("/login", admin.LoginApi{}.Login)
	// }

}
