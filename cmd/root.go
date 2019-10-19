package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"transfer/json"
)

func GetVersion() string {
	return "v0.0.0.1"
}

var RootCmd = &cobra.Command{
	Use: "run",
	// cmd cobra 对象,掌控cmd args 输入参数
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := cmd.Flags().GetBool("version"); err != nil {
			fmt.Println(GetVersion())
		}
		if _, err := cmd.Flags().GetString("xx"); err != nil {
			fmt.Println(GetVersion())

		}
		list, err := cmd.Flags().GetStringArray("tt")
		if err != nil {
			panic(err)
		}
		for _, value := range list {
			fmt.Println(value)
			ss, _ := json.Marshal(value)
			fmt.Println(string(ss))

		}
	},
}

// 一般在main func 中初始化参数变量
func init() {
	RootCmd.Flags().StringP("run", "d", "", "run dev version")
	RootCmd.Flags().BoolP("version", "v", false, "current version")
	RootCmd.Flags().String("xx", "x", "")
	RootCmd.Flags().StringArrayP("tt", "t", []string{}, "")

}
