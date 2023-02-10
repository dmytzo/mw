package main

import (
	"fmt"
	"os"

	"golang.org/x/exp/maps"
	"mw/action"
	"mw/std"
)

const (
	commandJWT          command = "jwt"
	commandBase64Encode command = "base64e"
	commandBase64Decode command = "base64d"
	commandDateTime     command = "dt"
	commandUUID         command = "uuid"
	commandCounter      command = "counter"
)

var (
	commands = map[command]commandFunc{
		commandJWT:          action.JWT,
		commandBase64Encode: action.Base64Encode,
		commandBase64Decode: action.Base64Decode,
		commandDateTime:     action.DateTime,
		commandUUID:         action.UUID,
		commandCounter:      action.Counter,
	}

	commandsList = maps.Keys(commands)
)

type (
	command     string
	commandFunc func([]string) (string, error)
)

func execute() {
	args := os.Args

	if len(args) < 2 {
		std.WriteError(fmt.Sprintf("no command provided: %v\n", commandsList))
		return
	}

	f, ok := commands[command(args[1])]
	if !ok {
		std.WriteError(fmt.Sprintf("unknown command: %v\n", commandsList))
		return
	}

	var commandArgs []string
	if len(args) > 2 {
		commandArgs = args[2:]
	}

	res, err := f(commandArgs)
	if err != nil {
		std.WriteError(fmt.Sprintf("command %s: %s\n", args[1], err.Error()))
		return
	}

	std.WriteResult(res, false)
}
