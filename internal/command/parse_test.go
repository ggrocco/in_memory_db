package command

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	testCases := []struct {
		line   string
		action ActionName
		args   []string
		err    error
	}{
		{"SET test-var-name 100", Set, []string{"test-var-name", "100"}, nil},
		{"GET test-var-name", Get, []string{"test-var-name"}, nil},
		{"UNSET test-var-name", Unset, []string{"test-var-name"}, nil},
		{"NUMEQUALTO 50", NumEqualTo, []string{"50"}, nil},
		{"", "", nil, ErrorInvalidCommand},
		{"BUUM", "", nil, ErrorInvalidCommand},
		{"BEGIN", Begin, []string{}, nil},
		{"COMMIT", Commit, []string{}, nil},
		{"ROLLBACK", Rollback, []string{}, nil},
	}

	for _, tc := range testCases {
		t.Run(tc.line, func(t *testing.T) {
			command, err := Parse(tc.line)

			if err != nil {
				if tc.err != err {
					t.Errorf("Error: %v", err)
				}
			} else {
				if tc.action != command.Action {
					t.Errorf("Expected action %v to be %v", tc.action, command.Action)
				}

				if !reflect.DeepEqual(tc.args, command.Args) {
					t.Errorf("Expected args %v to be %v", tc.args, command.Args)
				}
			}
		})
	}
}
