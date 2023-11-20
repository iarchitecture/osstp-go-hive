package main

import (
	core_database "osstp-go-hive/go-core/pkg/database"
	core_global "osstp-go-hive/go-core/pkg/global"
	core_viper "osstp-go-hive/go-core/pkg/viper"
	"osstp-go-hive/system/command"
	"osstp-go-hive/system/global"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//或 		  			GOPROXY='https://proxy.golang.com.cn,direct'
//go:generate go mod tidy
//go:generate go mod download

func init() {
	// 配置全局变量：core_global.Core_config
	core_viper.Viper(&core_global.Core_config, "/system/config/config_settings.yaml")
	global.DB = core_database.InitMysql(core_global.Core_config.MysqlConfig)
}

func main() {

	//go: cobra方式 启动
	command.Execute()

}
