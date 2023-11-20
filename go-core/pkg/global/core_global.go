package core_global

import (
	core_config "osstp-go-hive/go-core/pkg/config"

	"gorm.io/gorm"
)

var (
	/// setting.yaml 文件读取
	Core_config core_config.Config
	/// 全局database 获取
	Core_DB  *gorm.DB
	Core_DBD *gorm.DB
)
