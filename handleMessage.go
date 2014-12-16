package main

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/daph/goslack"
	"golang.org/x/net/websocket"
)

func handleMessage(msg goslack.Event, ws *websocket.Conn, conf Config) {
	// If the message isn't of type message or is sent my the bot user or whas not sent by @<user>
	// then return
	if msg.Type != "message" || msg.Type == conf.user || !strings.Contains(msg.Text, conf.user) {
		return
	}

	command := strings.Split(msg.Text, " ")
	if len(command) < 2 {
		goslack.SendMessage(ws, goslack.Event{msgId, "message", msg.Channel, "herp", "", ""})
		return
	}

	switch command[1] {
	case "figlet":
		if len(command) < 3 {
			goslack.SendMessage(ws, goslack.Event{msgId, "message", msg.Channel, "herp", "", ""})
			return
		}
		output, err := figlet(command[2:])
		if err != nil {
			goslack.SendMessage(ws, goslack.Event{msgId, "message", msg.Channel,
				fmt.Sprintf("There was an error running your command. ERR: %v", err), "", ""})
			return
		}
		goslack.SendMessage(ws, goslack.Event{msgId, "message", msg.Channel, "```" + output + "```", "", ""})

	default:
		goslack.SendMessage(ws, goslack.Event{msgId, "message", msg.Channel, "derp", "", ""})
	}

	msgId++
}

func figlet(command []string) (string, error) {
	figletCmd := exec.Command("figlet", strings.Join(command, " "))
	figletOut, err := figletCmd.Output()
	if err != nil {
		return "", err
	}
	return string(figletOut), nil
}
