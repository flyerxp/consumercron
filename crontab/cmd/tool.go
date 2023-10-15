/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"cron/service/tool"
	"github.com/flyerxp/lib/app"
	"github.com/spf13/cobra"
)

// toolCmd represents the tool command
var toolCmd = &cobra.Command{
	Use:   "tool",
	Short: "tool makefile",
	Long:  `根据表生成文件`,
	Run: func(cmd *cobra.Command, args []string) {
		defer func() {
			app.Shutdown(context.Background())
		}()
		//name, _ := cmd.Flags().GetString("name")
		act, _ := cmd.Flags().GetString("act")
		switch act {
		case "makefile":
			tool.Make()
		default:
			cmd.Printf("act 错误,目前只支持 listen")
		}
	},
}

func init() {
	toolCmd.Flags().String("name", "toolConf", "toolConf")
	toolCmd.Flags().String("act", "listen", "listen")
	rootCmd.AddCommand(toolCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// toolCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// toolCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
