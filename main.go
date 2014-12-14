package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/daph/goslack"
)

type SlackMessage struct {
	Id      int    `json:"id"`
	Type    string `json:"type"`
	Channel string `json:"channel"`
	Text    string `json:"text"`
}

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
	fmt.Println(ws)
	defer ws.Close()
	var message SlackMessage
	message.Id += 1
	message.Type = "message"
	message.Channel = conf.channel
	message.Text = "hello slack"
	b_message, err := json.Marshal(message)
	if err != nil {
		fmt.Println("Could not marshal the message. ERR: %v", err)
		os.Exit(-1)
	}
	_, err = ws.Write(b_message)
	if err != nil {
		fmt.Println("Couldn't write message. ERR: %v", err)
		os.Exit(-1)
	}

}
