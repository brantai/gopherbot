package main

import (
	"fmt"
	"os"
	"flag"
	"github.com/daph/goslack"
)

func main() {
        flag.StringVar(&configFile, "config", "gopher.cfg", "configuration file to load")
        flag.StringVar(&logPath, "log", ".", "path to log file")
	flag.Parse()
	InitLogger()
	err := loadConfig()
	if err != nil {
		debugLog.Printf("Could not load config. ERR: %v", err)
		os.Exit(1)
	}
	debugLog.Println("Startup")

	if configMap["slack_token"] == "" {
		debugLog.Printf("No slack_token in config file")
		os.Exit(1)
	}
	client, err := goslack.NewClient(configMap["slack_token"])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer client.Close()
	for {
		select {
		case msg := <-client.MsgIn:
			go handleMessage(msg, client) //in handleMessage.go
		}
	}
}
