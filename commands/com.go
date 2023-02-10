package commands

import (
	"context"
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"github.com/Andrem19/Timer_Job/helpers"
	"github.com/Andrem19/Timer_Job/variables"
	db "github.com/Andrem19/Timer_Job/db/sqlc"
)

func Command(cmd []string) (string, error) {
	switch strings.ToLower(cmd[1]) {
	case "exec":
		return Command_exec(cmd)
	case "s":
		return Command_save(cmd)
	case "ls":
		return Cmd_ls()
	case "run":
		return Command_run(cmd[2])
	default:
		return "Unknown command", nil
	}
}

func Cmd_ls() (string, error) {
	ls_com, err := variables.Queries.ListCommands(context.Background())
	if err != nil {
		helpers.AddToLog(err.Error())
		return "Something wrong", err
	}
	var sb strings.Builder
	for i := 0; i < len(ls_com); i++ {
		sb.WriteString("Id: ")
		str_id := strconv.FormatInt(ls_com[i].ID, 10)
		sb.WriteString(str_id)
		sb.WriteString(" ")
		sb.WriteString("CMD:")
		sb.WriteString(ls_com[i].Cmd)
		sb.WriteString("\n")
	}
return sb.String(), nil
}

func Command_exec(cmd []string) (string, error) {
	var sb strings.Builder
	for i := 2; i < len(cmd); i++ {
		sb.WriteString(cmd[i])
		sb.WriteString(" ")
	}
	fmt.Println(sb.String())
	out, _ := exec.Command("/bin/sh", "-c", sb.String()).Output()

    output := string(out[:])
    return output, nil
}

func Command_run(cmd string) (string, error) {
	id, err := strconv.Atoi(cmd)
	if err != nil {
		helpers.AddToLog(err.Error())
		return "Something wrong", err
	}
	var com db.Command
	com, err = variables.Queries.GetCommand(context.Background(), int64(id))

	if err != nil {
		helpers.AddToLog(err.Error())
		return "Something wrong", err
	}
	out, err := exec.Command("/bin/sh", "-c", com.Cmd).Output()
	if err != nil {
		helpers.AddToLog(err.Error())
		return "Something wrong", err
	}
    output := string(out[:])
    return output, nil
}
func Command_save(cmd []string) (string, error) {
	var sb strings.Builder
	for i := 2; i < len(cmd); i++ {
		sb.WriteString(cmd[i])
		sb.WriteString(" ")
	}
	id, err := variables.Queries.CreateCommand(context.Background(), sb.String())
	if err != nil {
		helpers.AddToLog(err.Error())
	}
	msg := fmt.Sprintf("Command successfuly created. Id: %d", id)
	return msg, nil
}