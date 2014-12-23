package main

import (
	"strings"

	"github.com/daph/goslack"
)

type Plugin interface {
	Name() string
	Help() string
	Execute(args []string) string
}

var Plugins = make([]Plugin, 0)

func handleMessage(msg goslack.Event, client *goslack.Client) {
	// If the message isn't of type message or is sent my the bot user or whas not sent by @<user>
	// then return
	if msg.Type != "message" || msg.User == client.Self.Id || len(strings.Split(msg.Text, "")) <= 0 || strings.Split(msg.Text, "")[0] != "!" {
		return
	}

	temp := strings.Join(strings.Split(msg.Text, "")[1:], "")
	command := strings.Split(temp, " ")
	if len(command) <= 0 {
		client.PushMessage(msg.Channel, "herp")
		return
	}

	if command[0] == "help" {
		if len(command) > 1 {
			for _, v := range Plugins {
				if command[1] == v.Name() {
					client.PushMessage(msg.Channel, v.Help())
					return
				}
			}
		}

		// If we get to this point, the user was either
		// asking for general help, or help on a plugin
		// that does not exist
		pluginNames := "| "
		for _, v := range Plugins {
			pluginNames += v.Name() + " | "
		}
		client.PushMessage(msg.Channel, pluginNames)
		return
	}

	for _, v := range Plugins {
		if command[0] == v.Name() {
			client.PushMessage(msg.Channel, v.Execute(command[1:]))
			return
		}
	}

	client.PushMessage(msg.Channel, "derp")
}
