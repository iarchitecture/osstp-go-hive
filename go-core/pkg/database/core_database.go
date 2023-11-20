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
	loggerConfig := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer 日志输出的目标
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 彩色打印
		},
	)

	db, err := gorm.Open(mysql.Open(connStr), &gorm.Config{
		Logger: loggerConfig,
	})

	// 判断连接情况
	if err != nil {
		fmt.Println("🐳🐳🐳🐳🐳🐳🐳🐳🐳🐳数据库连接失败🐳🐳🐳🐳🐳🐳🐳🐳🐳🐳")
		panic(err)
	}
	fmt.Println("🐬🐬🐬🐬🐬🐬🐬🐬🐬🐬数据库连接成功🐬🐬🐬🐬🐬🐬🐬🐬🐬🐬")

	core_global.Core_DB = db
	core_global.Core_DBD = db.Debug()

	return db
}
