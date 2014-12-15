package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/daph/goslack"
)

// Since each message id needs to be unique for the session
// This gets to be global and gets to be incremented any where
// I feel like. Methinks this will cause problems, but oh well.
var msgId int

type Config struct {
	token   string
	channel string
	user    string
}

func newConfig() Config {
	token, err := ioutil.ReadFile("token")
	if err != nil {
		fmt.Println("No 'token' file!")
		os.Exit(-1)
	}
	channel, err := ioutil.ReadFile("channel")
	if err != nil {
		fmt.Println("No 'channel' file!")
		os.Exit(-1)
	}
	user, err := ioutil.ReadFile("user")
	if err != nil {
		fmt.Println("No 'user' file!")
		os.Exit(-1)
	}

	return Config{string(token), string(channel), string(user)}
}

func main() {

	msgId = 1
	conf := newConfig()

	ws, err := goslack.Connect(conf.token)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer ws.Close()
	var message goslack.MessageSend
	message.Id = msgId
	message.Type = "message"
	message.Channel = conf.channel
	message.Text = "hello slack"
	goslack.SendMessage(ws, message)
	msgId++
	chat_ch := make(chan goslack.MessageRecv)
	go goslack.ReadMessages(ws, chat_ch)
	for {
		select {
		case msg := <-chat_ch:
			go handleMessage(msg, ws, conf) //in handleMessage.go
		case <-time.After(20 * time.Second):
			err := goslack.SendMessage(ws, goslack.MessageSend{msgId, "ping", "", ""})
			msgId++
			if err != nil {
				fmt.Printf("Problem sending ping. ERR: %v", err)
				os.Exit(-1)
			}
		}
	}
}
