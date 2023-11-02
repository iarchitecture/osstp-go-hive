package routers

import (
	"osstp-go-hive/controllers/front"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	v1 := r.Group("/api/user")
	{
		v1.GET("/register", front.UserController{}.UserRegisterHandle())
		v1.GET("/userInfo", front.UserController{}.GetUserInfo)

	}

	return r

}
