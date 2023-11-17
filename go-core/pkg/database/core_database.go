package core_database

import (
	"fmt"
	core_viper "osstp-go-hive/go-core/pkg/viper"

	"github.com/spf13/cobra"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// db 基本配置
var (
	localhost = "127.0.0.1"
	port      = 3306
	dbname    = "go_admin"
	username  = "root"
	password  = "Root@20211001"
)

// DB 全局变量，当需要数据库时可以直接调用
var DB *gorm.DB

// DBD 全局debug变量，在开发时使用DBD能够快速查看sql语句
var DBD *gorm.DB

// init初始化数据库连接与配置，当调用其他方法或变量时，会自动执行init函数
var (
	configYml string
	apiCheck  bool
	StartCmd  = &cobra.Command{
		Use:          "server",
		Short:        "Start API server",
		Example:      "go-admin server -c config/settings.yml",
		SilenceUsage: true,
		PreRun: func(cmd *cobra.Command, args []string) {

		},
	}
)

func Init() {
	StartCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "config/settings.yml", "Start server with provided configuration file")
	print(configYml)
	core_viper.Viper(DBdd, configYml)

	// fmt.Sprintf 内置格式字符串
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, localhost, port, dbname)
	// 连接数据库，mysql.Config 后面是一些高级配置参数，可以根据需要进行修改
	db, err := gorm.Open(mysql.New(mysql.Config{
		DriverName:                    "",
		ServerVersion:                 "",
		DSN:                           dsn, // 连接数据库信息
		DSNConfig:                     nil,
		Conn:                          nil,
		SkipInitializeWithVersion:     false,
		DefaultStringSize:             0,
		DefaultDatetimePrecision:      nil,
		DisableWithReturning:          false,
		DisableDatetimePrecision:      false,
		DontSupportRenameIndex:        false,
		DontSupportRenameColumn:       false,
		DontSupportForShareClause:     false,
		DontSupportNullAsDefaultValue: false,
		DontSupportRenameColumnUnique: false,
	}), &gorm.Config{ //&gorm.Config 后面的参数是相关配置，可以根据开发进行修改
		SkipDefaultTransaction:                   false,
		NamingStrategy:                           nil,
		FullSaveAssociations:                     false,
		Logger:                                   nil,
		NowFunc:                                  nil,
		DryRun:                                   false,
		PrepareStmt:                              false,
		DisableAutomaticPing:                     false,
		DisableForeignKeyConstraintWhenMigrating: false,
		IgnoreRelationshipsWhenMigrating:         false,
		DisableNestedTransaction:                 false,
		AllowGlobalUpdate:                        false,
		QueryFields:                              false,
		CreateBatchSize:                          0,
		TranslateError:                           false,
		ClauseBuilders:                           nil,
		ConnPool:                                 nil,
		Dialector:                                nil,
		Plugins:                                  nil,
	})
	// 判断连接情况
	if err != nil {
		fmt.Println("----------数据库连接失败--------------")
		//panic 抛出异常，并终止程序
		panic(err)
	}
	fmt.Println("----------数据库连接成功--------------")
	// 分别赋值给你全局变量DB 和DBD
	DB = db
	DBD = db.Debug()
}
