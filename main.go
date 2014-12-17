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
	go client.ReadMessages()
	go client.SendMessages()
	for {
		select {
		case msg := <-client.MsgIn:
			go handleMessage(msg, &client) //in handleMessage.go
		}
	}
}
