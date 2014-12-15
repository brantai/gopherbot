package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/daph/goslack"
	"golang.org/x/net/websocket"
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

func handleMessage(msg goslack.MessageRecv, ws *websocket.Conn, conf Config) {
	if msg.Type != "message" || msg.Type == conf.user || !strings.Contains(msg.Text, conf.user) {
		return
	}
	goslack.SendMessage(ws, goslack.MessageSend{msgId, "message", msg.Channel, "hello"})
	msgId++
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
			go handleMessage(msg, ws, conf)
		case <-time.After(30 * time.Second):
			goslack.SendMessage(ws, goslack.MessageSend{msgId, "ping", "", ""})
			msgId++
		}
	}
}
