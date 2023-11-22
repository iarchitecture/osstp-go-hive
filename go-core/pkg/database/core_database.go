package core_database

import (
	"fmt"
	"log"
	"os"
	core_config "osstp-go-hive/go-core/pkg/config"
	core_global "osstp-go-hive/go-core/pkg/global"
	"time"

	"github.com/spf13/cobra"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

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

func InitMysql(config core_config.MysqlConfig) (db *gorm.DB) {
	var connStr = fmt.Sprintf(config.Dsn())

	mysqlConfig := mysql.Config{
		DSN:                       connStr, // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}

	loggerConfig := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer 日志输出的目标
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 彩色打印
		},
	)

	db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		Logger: loggerConfig,
	})

	// 判断连接情况
	if err != nil {
		fmt.Println("🐳🐳🐳🐳🐳🐳🐳🐳🐳🐳数据库连接失败🐳🐳🐳🐳🐳🐳🐳🐳🐳🐳")
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("🐳🐳🐳🐳🐳🐳🐳🐳🐳🐳数据库创建失败🐳🐳🐳🐳🐳🐳🐳🐳🐳🐳")
	}

	// 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(config.MaxIdleConns)

	// 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(config.MaxOpenConns)

	// 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	fmt.Println("🐬🐬🐬🐬🐬🐬🐬🐬🐬🐬数据库连接成功🐬🐬🐬🐬🐬🐬🐬🐬🐬🐬")

	core_global.Core_DB = db
	core_global.Core_DBD = db.Debug()

	return db
}

func MigrateModel(dst ...interface{}) {
	err := core_global.Core_DBD.AutoMigrate(dst...)
	if err != nil {
		fmt.Printf("🐳🐳🐳🐳🐳🐳🐳🐳🐳🐳数据库AutoMigrate异常 %v🐳🐳🐳🐳🐳🐳🐳🐳🐳🐳\n", err)
		os.Exit(0)
	}
	fmt.Println("🐬🐬🐬🐬🐬🐬🐬🐬🐬🐬数据库AutoMigrate完成🐬🐬🐬🐬🐬🐬🐬🐬🐬🐬")
}
