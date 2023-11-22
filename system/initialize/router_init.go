package initialize

import (
	core_global "osstp-go-hive/go-core/pkg/global"
	"osstp-go-hive/system/routers/admin"
	"osstp-go-hive/system/routers/shop"

	// _ "go_code/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouters() *gin.Engine {
	engine := gin.Default()

	// swagger：生产环境注释掉
	if core_global.Core_config.ApplicationConfig.Mode != "prod" {
		engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	// 注册后台管理系统路由
	admin.RegisterSystemRouters(engine)

	// 注册商城路由
	shop.RegisterShopRouters(engine)

	return engine
}
