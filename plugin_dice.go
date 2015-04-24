package main

import (
	"fmt"
	"strings"
        "math/rand"
	"strconv"
)

type DicePlugin struct{}

func init() {
	Plugins = append(Plugins, DicePlugin{})
}

func (p DicePlugin) Help() string {
	return "Roll the specified type of die. Usage: !dice <die type>"
}

func (p DicePlugin) Name() string {
	return "dice"
}
func (p DicePlugin) Execute(command string) string {
	if len(command) < 1 {
		return "herp"
	}
        
	DiceCmd := Roll(command)
	DiceOut, err := DiceCmd.Output()
	if err != nil {
		return fmt.Sprintf("Couldn't get output. ERR %v", err)
	}
	return "```" + string(DiceOut) + "```"
}
func (p DicePlugin) Roll(command string) string {
        sides, _ := strconv.Atoi(command)
        result := rand.Intn(sides)
        return result
}
