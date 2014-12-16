package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/daph/goslack"
)

type Config struct {
	token string
}

func newConfig() Config {
	token, err := ioutil.ReadFile("token")
	if err != nil {
		fmt.Println("No 'token' file!")
		os.Exit(-1)
	}

	return Config{string(token)}
}

func main() {
	InitLogger()
	debugLog.Println("Startup")
	conf := newConfig()

	client, err := goslack.NewClient(conf.token)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer client.Ws.Close()

	for client.Ws.IsClientConn() {
		msg, err := client.ReadMessages()
		if err != nil {
			debugLog.Printf("Could not read messages. ERR: %v", err)
		}
		if (msg != goslack.Event{}) {
			go handleMessage(msg, &client, conf) //in handleMessage.go
		}
	}
}
