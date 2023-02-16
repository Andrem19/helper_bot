package switcher

import (
	"errors"
	"strings"

	"github.com/Andrem19/Timer_Job/coin"
	com "github.com/Andrem19/Timer_Job/commands"
	"github.com/Andrem19/Timer_Job/helpers"
	"github.com/Andrem19/Timer_Job/timer"
	"github.com/Andrem19/Timer_Job/variables"
)
var last_command string

func Switcher(message string, chat_id string) (string, error) {
	if chat_id != variables.Config.USER_ID {
		return "You are not the chosen one", errors.New("you are enemy, go away")
	}
	commands := helpers.Decode(message)
	if strings.ToLower(commands[0]) != "repeat" && strings.ToLower(commands[1]) != "stop"  {
		last_command = message
	}

	switch strings.ToLower(commands[0]) {
	case "timer":
		return timer.Timer(commands, chat_id)
	case "repeat":
		return Switcher(last_command, chat_id)
	case "cmd":
		return com.Command(commands)
	case "coin":
		return coin.CountPriceAndAmounts(commands)
	default:
		return "", nil
	}
}