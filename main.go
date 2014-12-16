package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/daph/goslack"
)

func getToken() string {
	token, err := ioutil.ReadFile("token")
	if err != nil {
		fmt.Println("No 'token' file!")
		os.Exit(-1)
	}

	return string(token)
}

func main() {
	InitLogger()
	debugLog.Println("Startup")
	client, err := goslack.NewClient(getToken())
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer client.Ws.Close()

	for {
		msg, err := client.ReadMessages()
		if err != nil {
			debugLog.Printf("Could not read messages. ERR: %v", err)
		}
		if (msg != goslack.Event{}) {
			go handleMessage(msg, &client) //in handleMessage.go
		}
	}
}
