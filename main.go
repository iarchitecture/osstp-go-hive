package main

import (
	"osstp-go-hive/system/command"
	"osstp-go-hive/system/global"
	"osstp-go-hive/system/initialize"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//或 		  			GOPROXY='https://proxy.golang.com.cn,direct'
//go:generate go mod tidy
//go:generate go mod download

func init() {
	initialize.InitViper()
	global.DB = initialize.InitDatabse()

	initialize.DatabaseAutoMigrate()

	global.Router = initialize.InitRouters()
}

//	@title			HIVE
//	@version		1.0
//	@description	Go 语言编程
//	@termsOfService	https://github.com/

func main() {

	//go: cobra方式 启动
	command.Execute()

}
