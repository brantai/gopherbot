package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

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

func reconnect(ws *websocket.Conn, conf Config) {
	ws.Close()
	ws, err := goslack.Connect(conf.token)
	if err != nil {
		debugLog.Printf("Could not reconnect")
		os.Exit(-1)
	}
}

func main() {
	InitLogger()

	msgId = 1
	conf := newConfig()

	ws, err := goslack.Connect(conf.token)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer ws.Close()
	var message goslack.Event
	message.Id = msgId
	message.Type = "message"
	message.Channel = conf.channel
	message.Text = "hello slack"
	goslack.SendMessage(ws, message)
	msgId++

	for ws.IsClientConn() {
		msg, err := goslack.ReadMessages(ws)
		if err != nil {
			if err == io.EOF {
				debugLog.Printf("Got EOF from server. Reconnect and connect")
				reconnect(ws, conf)
			}
			debugLog.Printf("Could not read messages. ERR: %v", err)
		}
		if (msg != goslack.Event{}) {
			go handleMessage(msg, ws, conf) //in handleMessage.go
		}
	}
}
