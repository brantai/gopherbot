package main

import (
	"strings"

	"github.com/daph/goslack"
)

func handleMessage(msg goslack.Event, client *goslack.Client) {

	if msg.Type != "message" || /* If the event isn't a message */
		msg.User == client.Self.Id || /* If the message was sent by the bot */
		len(strings.Split(msg.Text, "")) <= 0 || /* If somehow the message is less than or equal to zero characters */
		strings.Split(msg.Text, "")[0] != "!" /* Or lastly, if the message we got does not begin with an '!' */ {
		return /* Then return, this isn't a message we want to bother handling. */
	}

	temp := strings.Join(strings.Split(msg.Text, "")[1:], "") // Split and Join the message text in order to remove the leading '!'
	command := strings.Split(temp, " ")                       // Split the message into words seperated by spaces.
	if len(command) <= 0 {                                    // If there was nothing after the '!', then send a herp
		client.PushMessage(msg.Channel, "herp")
		return
	}

	// We check if the user wants help here
	if command[0] == "help" {
		if len(command) > 1 { // If there's something after 'help',
			for _, v := range Plugins { // we iterate through the list of plugins,
				if command[1] == v.Name() { // to see if we have the command they want help with,
					client.PushMessage(msg.Channel, v.Help()) // and then send a message with the help from the plugin.
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

	/*
		Iterate through the list of Plugins
		and check if what the user asked for it there.
		If it is, then execute the plugin and give it whatever
		collection of words came after the plugin name and send a message
		with the return of the plugin (plugins return a string)
	*/
	for _, v := range Plugins {
		if command[0] == v.Name() {
			client.PushMessage(msg.Channel, v.Execute(command[1:]))
			return
		}
	}

	client.PushMessage(msg.Channel, "derp") // If we've got down here, then the user tried to use a plugin that does not exist.
}
