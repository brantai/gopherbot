package main

import (
	"fmt"
        "github.com/tonio-ramirez/dice"
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

func intsToStrings(ints []int) (strings []string) {
        strings = make([]string, len(ints))
        for i, v := range ints {
                strings[i] = strconv.Itoa(v)
        }
        return
}

func (p DicePlugin) Execute(command []string) string {
	if len(command) < 1 {
		return "herp"
	}
	var dieType string = command[0]
	roll, err := dice.Roll(dieType)
	if err != nil {
		return fmt.Sprintf("Couldn't get output. ERR %v", err)
	}
	return "```" + strconv.Itoa(roll.Total) + "```"
}
