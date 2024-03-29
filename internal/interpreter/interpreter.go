package interpreter

import (
	"fmt"
	"strconv"

	"in_memory_db/internal/command"
	"in_memory_db/internal/db"
)

type ArgumentError struct {
	action    command.ActionName
	numOfArgs int
}

func (e ArgumentError) Error() string {
	return fmt.Sprintf("the command: %v, require %d args", e.action, e.numOfArgs)
}

func Run(db db.DB, line string) (string, error) {
	cmd, err := command.Parse(line)
	if err != nil {
		return "", err
	}

	switch cmd.Action {
	case command.Set:
		return runIfValid(cmd, 2, func() { db.Set(cmd.Args[0], cmd.Args[1]) })
	case command.Get:
		return runIfValidWithOutput(cmd, 1, func() string { return db.Get(cmd.Args[0]) })
	case command.Unset:
		return runIfValid(cmd, 1, func() { db.Unset(cmd.Args[0]) })
	case command.NumEqualTo:
		return runIfValidWithOutput(cmd, 1, func() string { return strconv.Itoa(db.NumEqualTo(cmd.Args[0])) })

	case command.Begin:
		return runIfValid(cmd, 0, func() { db.Begin() })
	case command.Commit:
		return runIfValidWithError(cmd, 0, func() error { return db.Commit() })
	case command.Rollback:
		return runIfValidWithError(cmd, 0, func() error { return db.Rollback() })
	}

	return "", nil
}

func runIfValid(cmd *command.Command, numOfArgs int, action func()) (string, error) {
	return runIfValidWithOutput(cmd, numOfArgs, func() string {
		action()
		return ""
	})
}

func runIfValidWithError(cmd *command.Command, numOfArgs int, action func() error) (string, error) {
	if len(cmd.Args) != numOfArgs {
		return "", ArgumentError{cmd.Action, numOfArgs}
	}

	return "", action()
}

func runIfValidWithOutput(cmd *command.Command, numOfArgs int, action func() string) (string, error) {
	if len(cmd.Args) != numOfArgs {
		return "", ArgumentError{cmd.Action, numOfArgs}
	}

	return action(), nil
}
