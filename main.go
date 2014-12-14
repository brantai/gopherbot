package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/daph/goslack"
)

type Config struct {
	token   string
	channel string
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

	return Config{string(token), string(channel)}
}

func main() {

	conf := NewConfig()

	ws, err := goslack.Connect(conf.token)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer ws.Close()
	var message goslack.SlackMessage
	message.Id += 1
	message.Type = "message"
	message.Channel = conf.channel
	message.Text = "hello slack"
	goslack.SendMessage(ws, message)
}
