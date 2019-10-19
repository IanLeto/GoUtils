package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"goUtils/utils"
	"transfer/logging"
)

func GetVersion() string {
	return "v0.0.0.1"
}

var RootCmd = &cobra.Command{
	Use: "debug",
	// cmd cobra 对象,掌控cmd args 输入参数
	Run: func(cmd *cobra.Command, args []string) {
		// 拿到run 这个命令的输入
		var (
			flags = cmd.Flags()
		)
		// 这个run 就是 debug -run ${param} 中的参数
		if utils.MustGetString(flags.GetString("run")) != "" {
			logging.GetStdLogger().Infof("%v", Logo)
		}

		version := utils.MustGetString(flags.GetString("version"))
		fmt.Printf("%v", version)
	},
}

// 被main 函数执行
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		panic(err)
	}
}

// 一般在main func 中初始化参数变量
func init() {
	RootCmd.Flags().StringP("run", "", "", "run dev version")
	RootCmd.Flags().String("version", "v", "")

}
