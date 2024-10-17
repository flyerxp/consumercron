/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"cron/service/nacos"
	"github.com/flyerxp/lib/v2/app"
	"github.com/spf13/cobra"
)

// nacosCmd represents the nacos command
var nacosCmd = &cobra.Command{
	Use:   "nacos",
	Short: "nasos listen",
	Long:  `更新config 的 redis缓存`,
	Run: func(cmd *cobra.Command, args []string) {
		defer func() {
			app.Shutdown(context.Background())
		}()
		name, _ := cmd.Flags().GetString("name")
		act, _ := cmd.Flags().GetString("act")
		switch act {
		case "listen":
			nacos.Listen(cmd, name)
		default:
			cmd.Printf("act 错误,目前只支持 listen")
		}
	},
}

func init() {
	nacosCmd.Flags().String("name", "nacosConf", "nacosConf")
	nacosCmd.Flags().String("act", "listen", "listen")
	rootCmd.AddCommand(nacosCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// nacosCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// nacosCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
