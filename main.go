package main

import (
	core_database "osstp-go-hive/go-core/pkg/database"
	core_viper "osstp-go-hive/go-core/pkg/viper"
	"osstp-go-hive/system/command"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//或 		  			GOPROXY='https://proxy.golang.com.cn,direct'
//go:generate go mod tidy
//go:generate go mod download

func main() {
	//go: cobra方式 启动
	command.Execute()

	core_viper.Viper(core_database.DBdd, "/system/config/settings.yaml")
	// print(core_database.Config.Database.port)

	core_database.Init()
}
