package command

import (
	"errors"
	"strings"
)

var ErrorInvalidCommand = errors.New("invalid command")

type ActionName string

const (
	Invalid    ActionName = "INVALID"
	Set        ActionName = "SET"
	Get        ActionName = "GET"
	Unset      ActionName = "UNSET"
	NumEqualTo ActionName = "NUMEQUALTO"
)

type Command struct {
	Action ActionName
	Args   []string
}

func Parse(line string) (*Command, error) {
	if len(line) == 0 {
		return nil, ErrorInvalidCommand
	}

	parts := strings.Split(line, " ")
	actionName := mapToActionName(strings.ToUpper(parts[0]))
	if actionName == Invalid {
		return nil, ErrorInvalidCommand
	}

	return &Command{
		Action: actionName,
		Args:   parts[1:],
	}, nil
}

func mapToActionName(action string) ActionName {
	switch action {
	case "SET":
		return Set
	case "GET":
		return Get
	case "UNSET":
		return Unset
	case "NUMEQUALTO":
		return NumEqualTo
	}

	return Invalid
}
