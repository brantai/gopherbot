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
	return "Roll the specified type of di(c)e. Usage: !roll <number of dice>d<die type>"
}

func (p DicePlugin) Name() string {
	return "roll"
}

func fateDice() string {
	total := 0
	for i := 0;i<4;i++{
		roll, err := dice.Roll("1d3")
		if err != nil {
			return  fmt.Sprintf("Couldn't get output. ERR %v", err)
		}
		switch {
			case roll.Total == 1:
				total--
			case roll.Total == 3:
				total++
		}
	}
	return strconv.Itoa(total)
}

func otherDice(dieType string) string {
	roll, err := dice.Roll(dieType)
	if err != nil {
		return fmt.Sprintf("Couldn't get output. ERR %v", err)
	}
	return strconv.Itoa(roll.Total)
}

func (p DicePlugin) Execute(command []string) string {
	if len(command) < 1 {
		return "herp"
	}
	var dieType string = command[0]
	var result string = ""
	if dieType == "dF"{
		result = fateDice()
	} else {
		result = otherDice(dieType)
	}

	return "```" + result + "```"
}
