package helpers

import (
	"errors"
	"strings"

	"github.com/Andrem19/Timer_Job/timer"
	"github.com/Andrem19/Timer_Job/variables"
)

func Switcher(message string, chat_id string) (string, error) {
	if chat_id != variables.Config.USER_ID {
		return "You are not the chosen one", errors.New("you are enemy, go away")
	}
	commands := Decode(message)

	switch strings.ToLower(commands[0]) {
	case "timer":
		return timer.Timer(commands, chat_id)
	default:
		return "", nil
	}
}