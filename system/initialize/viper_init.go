package initialize

import (
	core_global "osstp-go-hive/go-core/pkg/global"
	core_viper "osstp-go-hive/go-core/pkg/viper"
)

func InitViper() {
	// 配置全局变量：core_global.Core_config
	core_viper.Viper(&core_global.Core_config, "/system/config/config_settings.yaml")

}
