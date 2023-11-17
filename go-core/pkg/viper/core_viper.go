package core_viper

import (
	"flag"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

const (
	ConfigEnvMode     = "OSSTP_CONFIG"
	ConfigDebugMode   = "config.debug.yaml"
	ConfigTestMode    = "config.test.yaml"
	ConfigReleaseMode = "config.release.yaml"
)

// Viper
// 优先级: 命令行 > 环境变量 > 默认值
func Viper(configg interface{}, path ...string) *viper.Viper {
	var config string

	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "choose config file.")
		flag.Parse()
		if config == "" { // 判断命令行参数是否为空
			if configEnv := os.Getenv(ConfigEnvMode); configEnv == "" { // 判断 coreinternal.ConfigEnv 常量存储的环境变量是否为空
				switch gin.Mode() {
				case gin.DebugMode:
					config = ConfigDebugMode
				case gin.ReleaseMode:
					config = ConfigReleaseMode
				case gin.TestMode:
					config = ConfigTestMode
				}
				fmt.Printf("gin model: %s, config path: %s\n", gin.EnvGinMode, config)
			} else { // coreinternal.ConfigEnv 常量存储的环境变量不为空 将值赋值于config
				config = configEnv
				fmt.Printf("gin model: %s, config path: %s\n", ConfigEnvMode, config)
			}
		} else { // 命令行参数不为空 将值赋值于config
			fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%s\n", config)
		}
	} else { // 函数传递的可变参数的第一个值赋值于config
		roott, _ := os.Getwd()
		config = roott + path[0]
		fmt.Printf("您正在使用func Viper()传递的值,config的路径为%s\n", config)
	}

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&configg); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&configg); err != nil {
		fmt.Println(err)
	}

	// root 适配性 根据root位置去找到对应迁移位置,保证root路径有效
	// coreglobal.Config.AutoCode.Root, _ = filepath.Abs("..")
	return v
}
