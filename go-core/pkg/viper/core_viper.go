package core_viper

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Viper
func Viper(rawVal any, path ...string) *viper.Viper {
	var config string

	roott, _ := os.Getwd()
	config = roott + path[0]
	fmt.Printf("配置文件路径是：%s\n", config)

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
		if err = v.Unmarshal(rawVal); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(rawVal); err != nil {
		fmt.Println(err)
	}

	fmt.Println("⭐⭐⭐⭐⭐⭐⭐⭐⭐⭐ Viper init success ⭐⭐⭐⭐⭐⭐⭐⭐⭐⭐")
	return v
}
