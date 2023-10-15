/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"consumer/service"
	"fmt"
	"github.com/spf13/cobra"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "run test --cnt=100",
	Run: func(cmd *cobra.Command, args []string) {
		cnt, ok := cmd.Flags().GetInt("cnt")
		if ok != nil {
			fmt.Println(ok)
		}
		service.Test(cnt)
	},
}

func init() {
	testCmd.Flags().Int("cnt", 0, "0")
	rootCmd.AddCommand(testCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
