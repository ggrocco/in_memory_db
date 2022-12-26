package interpreter

import (
	"errors"
	"in_memory_db/internal/command"
	"in_memory_db/internal/db"
	"testing"
)

func TestArgumentError(t *testing.T) {
	err := ArgumentError{command.Set, 0}
	if err.Error() != "the command: SET, require 0 args" {
		t.Errorf("Fail: %v", err)
	}
}

func TestExample1(t *testing.T) {
	testCases := []struct {
		command string
		output  string
		err     error
	}{
		{"SET test-var-name 100", "", nil},
		{"GET test-var-name", "100", nil},
		{"UNSET test-var-name", "", nil},
		{"GET test-var-name", "Nil", nil},
		{"SET test-var-name-1 50", "", nil},
		{"SET test-var-name-2 50", "", nil},
		{"NUMEQUALTO 50", "2", nil},
		{"SET test-var-name-2 10", "", nil},
		{"NUMEQUALTO 50", "1", nil},
		{"INVALID_COMMAND 50", "", command.ErrorInvalidCommand},
	}

	db := db.New()

	for _, tc := range testCases {
		t.Run(tc.command, func(t *testing.T) {
			output, err := Run(db, tc.command)

			if err != nil {
				if !errors.Is(err, tc.err) {
					t.Errorf("Unexpected error: %v", err)
				}
				return
			}

			if tc.output != output {
				t.Errorf("Fail: expected %v to be %v", tc.output, output)
			}
		})
	}
}
