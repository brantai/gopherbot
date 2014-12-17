package main

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/daph/goslack"
)

func handleMessage(msg goslack.Event, client *goslack.Client) {
	// If the message isn't of type message or is sent my the bot user or whas not sent by @<user>
	// then return
	if msg.Type != "message" || !strings.Contains(msg.Text, client.Self.Id) {
		return
	}

	command := strings.Split(msg.Text, " ")
	if len(command) < 2 {
		client.PushMessage(msg.Channel, "herp")
		return
	}

	switch command[1] {
	case "figlet":
		if len(command) < 3 {
			client.PushMessage(msg.Channel, "herp")
			return
		}
		output, err := figlet(command[2:])
		if err != nil {
			client.PushMessage(msg.Channel, fmt.Sprintf("There was an error running your command. ERR: %v", err))
			return
		}
		client.PushMessage(msg.Channel, "```"+output+"```")

	default:
		client.PushMessage(msg.Channel, "derp")
	}

}

func figlet(command []string) (string, error) {
	figletCmd := exec.Command("figlet", strings.Join(command, " "))
	figletOut, err := figletCmd.Output()
	if err != nil {
		return "", err
	}
	return string(figletOut), nil
}
