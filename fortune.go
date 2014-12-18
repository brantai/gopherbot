package main

import (
	"fmt"
	"os/exec"
)

type FortunePlugin struct{}

func (p FortunePlugin) Name() string {
	return "fortune"
}
func (p FortunePlugin) Execute(command []string) string {
	_ = command
	fortuneCmd := exec.Command("fortune")
	fortuneOut, err := fortuneCmd.Output()
	if err != nil {
		return fmt.Sprintf("Couldn't get fortune output. ERR %v", err)
	}
	return "```" + string(fortuneOut) + "```"
}
