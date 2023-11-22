package initialize

import (
	core_database "osstp-go-hive/go-core/pkg/database"
	core_global "osstp-go-hive/go-core/pkg/global"
	"osstp-go-hive/system/app/admin/model"

	"gorm.io/gorm"
)

func InitDatabse() *gorm.DB {
	return core_database.InitMysql(core_global.Core_config.MysqlConfig)
}

func DatabaseAutoMigrate() {
	core_database.MigrateModel(model.SysUser{})
}
