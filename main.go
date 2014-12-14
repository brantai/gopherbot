package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/daph/goslack"
)

type SlackMessage struct {
	Id      int    `json:"id"`
	Type    string `json:"type"`
	Channel string `json:"channel"`
	Text    string `json:"text"`
}

func main() {

	ws, err := goslack.Connect("")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Println(ws)
	defer ws.Close()
	var message SlackMessage
	message.Id += 1
	message.Type = "message"
	message.Channel = ""
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
