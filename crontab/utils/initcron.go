package utils

import (
	"context"
	"fmt"
	"github.com/flyerxp/lib/v2/config"
	"github.com/flyerxp/lib/v2/logger"
	"gopkg.in/yaml.v3"
	"io"
	"log"
	"os"
)

type AppConfig struct {
	App struct {
		Name   string `yaml:"name" json:"name"`
		Type   string `yaml:"type" json:"type"`
		Logger struct {
			OutputPaths []string `json:"outputPaths" yaml:"outputPaths"`
		} `yaml:"logger"`
		ErrLog struct {
			OutputPaths []string `json:"outputPaths" yaml:"outputPaths"`
		} `yaml:"errlog"`
	}
}

func GetLogCtx(name string) context.Context {
	yamlFile := config.GetConfFile("cron.yml")
	fd, _ := os.Open(yamlFile)
	yamlContent, err := io.ReadAll(fd)
	if err != nil {
		log.Println(err)
	}
	fd.Close()
	data := new(AppConfig)
	yaml.Unmarshal(yamlContent, data)
	config.GetConf().App.Type = data.App.Type
	config.GetConf().App.Name = data.App.Name
	config.GetConf().App.Logger.OutputPaths = data.App.Logger.OutputPaths
	config.GetConf().App.ErrLog.OutputPaths = data.App.ErrLog.OutputPaths
	fmt.Println("notice logs path:", data.App.Logger.OutputPaths[0]+"_"+name+"_*")
	fmt.Println("notice logs path:", data.App.ErrLog.OutputPaths[0]+"_"+name+"_*")
	logger.SetLogPathPrefix(name)
	return logger.GetContext(context.Background(), "cron")
}
