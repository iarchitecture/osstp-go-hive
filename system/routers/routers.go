package routers

import (
	"osstp-go-hive/system/routers/admin"

	"github.com/gin-gonic/gin"
)

func InitRouters() *gin.Engine {
	engine := gin.Default()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	// 注册后台管理系统路由
	admin.RegisterSystemRouters(engine)

	return engine
}
