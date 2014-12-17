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
	if msg.Type != "message" || msg.User == client.Self.Id || len(strings.Split(msg.Text, "")) <= 0 || strings.Split(msg.Text, "")[0] != "!" {
		return
	}

	temp := strings.Join(strings.Split(msg.Text, "")[1:], "")
	command := strings.Split(temp, " ")
	if len(command) < 2 {
		client.PushMessage(msg.Channel, "herp")
		return
	}

	switch command[0] {
	case "figlet":
		if len(command) < 2 {
			client.PushMessage(msg.Channel, "herp")
			return
		}
		output, err := figlet(command[1:])
		if err != nil {
			client.PushMessage(msg.Channel, fmt.Sprintf("There was an error running your command. ERR: %v", err))
			return
		}
		client.PushMessage(msg.Channel, "```"+output+"```")

	case "ud":
		if len(command) < 2 {
			client.PushMessage(msg.Channel, "herp")
			return
		}
		client.PushMessage(msg.Channel, ud(command[1:]))

	case "google":
		if len(command) < 2 {
			client.PushMessage(msg.Channel, "herp")
			return
		}
		client.PushMessage(msg.Channel, google(command[1:]))

	case "giphy":
		if len(command) < 2 {
			client.PushMessage(msg.Channel, "herp")
			return
		}
		client.PushMessage(msg.Channel, giphy(command[1:]))
	case "randomgif":
		client.PushMessage(msg.Channel, randomgif())

	default:
		client.PushMessage(msg.Channel, string(temp)+" "+string(command[1]))
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

func ud(command []string) string {
	query := strings.Join(command, "%20")
	return "http://www.urbandictionary.com/define.php?term=" + query
}

func google(command []string) string {
	query := strings.Join(command, "%20")
	return "https://www.google.com/search?q=" + query
}
