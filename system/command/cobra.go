package command

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	core_global "osstp-go-hive/go-core/pkg/global"
	"osstp-go-hive/system/global"
	"time"

	"github.com/spf13/cobra"
)

type server interface {
	ListenAndServe() error
}

var rootCommand = &cobra.Command{
	Use:     "HIVE",
	Short:   "HIVE root command",
	Example: "go start server -c config/settings.yml",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running HIVE server...")
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return appRunInfo()
	},
}

func appRunInfo() error {
	fmt.Printf("%s\n", `欢迎使用 `+`HIVE`)

	address := fmt.Sprintf(":%d", core_global.Core_config.ApplicationConfig.Port)

	server := initServer(address, global.Router)

	fmt.Printf(`
	swagger文档: http://127.0.0.1%s/swagger/index.html
	服务器默认地址: http://127.0.0.1%s
	`, address, address)

	fmt.Println(server.ListenAndServe().Error())

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println("Server exiting")

	return nil

}

func Execute() {
	if err := rootCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
