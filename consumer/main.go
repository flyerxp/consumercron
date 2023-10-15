/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"consumer/cmd"
	"time"
)

func main() {
	time.LoadLocation("Asia/Shanghai")
	cmd.Execute()
}
