package main

import (
	"strings"

	"github.com/daph/goslack"
	"golang.org/x/net/websocket"
)

func handleMessage(msg goslack.MessageRecv, ws *websocket.Conn, conf Config) {
	// If the message isn't of type message or is sent my the bot user or whas not sent by @<user>
	// then return
	if msg.Type != "message" || msg.Type == conf.user || !strings.Contains(msg.Text, conf.user) {
		return
	}

	command := strings.Split(msg.Text, " ")
	if len(command) < 2 {
		goslack.SendMessage(ws, goslack.MessageSend{msgId, "message", msg.Channel, "derp"})
		return
	}

	goslack.SendMessage(ws, goslack.MessageSend{msgId, "message", msg.Channel, command[1]})

	msgId++
}
