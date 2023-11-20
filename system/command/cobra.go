package command

import (
	"fmt"
	"os"
	"osstp-go-hive/system/routers"

	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use:     "HIVE",
	Short:   "HIVE root command",
	Example: "go start server -c config/settings.yml",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		appInfo()
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run HIVE server...")

	},
}

func appInfo() {
	fmt.Printf("%s\n", `欢迎使用 `+`HIVE`)
}

func Init() {
	routers.InitRouters()
}

func Execute() {
	if err := rootCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
