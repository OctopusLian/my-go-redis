/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-09-03 23:15:29
 * @LastEditors: neozhang
 * @LastEditTime: 2022-09-03 23:18:17
 */
package main

import (
	"fmt"
	"mygoredis/config"
	"mygoredis/lib/logger"
	"mygoredis/tcp"
	"os"
)

const configFile string = "redis.conf"

var defaultProperties = &config.ServerProperties{
	Bind: "0.0.0.0",
	Port: 6379,
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	return err == nil && !info.IsDir()
}

func main() {
	logger.Setup(&logger.Settings{
		Path:       "logs",
		Name:       "mygodis",
		Ext:        "log",
		TimeFormat: "2006-01-02",
	})

	if fileExists(configFile) {
		config.SetupConfig(configFile)
	} else {
		config.Properties = defaultProperties
	}

	err := tcp.ListenAndServeWithSignal(
		&tcp.Config{
			Address: fmt.Sprintf("%s:%d",
				config.Properties.Bind,
				config.Properties.Port),
		},
		tcp.MakeHandler())
	if err != nil {
		logger.Error(err)
	}
}
