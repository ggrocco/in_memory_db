package command

import (
	"errors"
	"strings"
)

var ErrorInvalidCommand = errors.New("invalid command")

type ActionName string

const (
	Set        ActionName = "SET"
	Get        ActionName = "GET"
	Unset      ActionName = "UNSET"
	NumEqualTo ActionName = "NUMEQUALTO"

	Begin    ActionName = "BEGIN"
	Commit   ActionName = "COMMIT"
	Rollback ActionName = "ROLLBACK"
)

var mapActions = map[string]ActionName{
	"SET":        Set,
	"GET":        Get,
	"UNSET":      Unset,
	"NUMEQUALTO": NumEqualTo,
	"BEGIN":      Begin,
	"COMMIT":     Commit,
	"ROLLBACK":   Rollback,
}

type Command struct {
	Action ActionName
	Args   []string
}

func Parse(line string) (*Command, error) {
	if len(line) == 0 {
		return nil, ErrorInvalidCommand
	}

	parts := strings.Split(line, " ")
	actionName := strings.ToUpper(parts[0])

	action, ok := mapActions[actionName]
	if !ok {
		return nil, ErrorInvalidCommand
	}

	return &Command{
		Action: action,
		Args:   parts[1:],
	}, nil
}
