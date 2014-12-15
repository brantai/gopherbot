package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/daph/goslack"
)

type Config struct {
	token   string
	channel string
	user    string
}

func NewConfig() Config {
	token, err := ioutil.ReadFile("token")
	if err != nil {
		os.Exit(-1)
	}
	channel, err := ioutil.ReadFile("channel")
	if err != nil {
		os.Exit(-1)
	}
	user, err := ioutil.ReadFile("user")
	if err != nil {
		os.Exit(-1)
	}

	return Config{string(token), string(channel), string(user)}
}

func main() {

	msgId := 1
	conf := NewConfig()

	ws, err := goslack.Connect(conf.token)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer ws.Close()
	var message goslack.MessageSend
	message.Id += 1
	message.Type = "message"
	message.Channel = conf.channel
	message.Text = "hello slack"
	goslack.SendMessage(ws, message)
	chat_ch := make(chan goslack.MessageRecv)
	go goslack.ReadMessages(ws, chat_ch)
	for {
		select {
		case msg := <-chat_ch:
			if msg.Type == "message" && msg.User != conf.user && strings.Contains(msg.Text, conf.user) {
				goslack.SendMessage(ws, goslack.MessageSend{msgId, "message", msg.Channel, "hello"})
				msgId++
				fmt.Println(msg)
				time.Sleep(time.Second * 1)
			}
		}
	}
}
