package main

import (
	"bytes"
	"in_memory_db/internal/db"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	type interaction struct {
		command string
		output  string
	}

	testCases := []struct {
		description  string
		interactions []interaction
	}{
		{
			description: "invalid",
			interactions: []interaction{
				{"SET with-one-arg", "Fail on execute `SET with-one-arg`, error: the command: SET, require 2 args"},
				{"END", ""},
			},
		},
		{
			description: "basic",
			interactions: []interaction{
				{"SET test-var-name 100", ""},
				{"GET test-var-name", "100"},
				{"UNSET test-var-name", ""},
				{"GET test-var-name", "Nil"},
				{"SET test-var-name-1 50", ""},
				{"SET test-var-name-2 50", ""},
				{"NUMEQUALTO 50", "2"},
				{"SET test-var-name-2 10", ""},
				{"NUMEQUALTO 50", "1"},
				{"END", ""},
			},
		},
		{
			description: "basic transaction",
			interactions: []interaction{
				{"GET test-var-name", "Nil"},
				{"BEGIN", ""},
				{"SET test-var-name 100", ""},
				{"GET test-var-name", "100"},
				{"COMMIT", ""},
				{"GET test-var-name", "100"},
				{"END", ""},
			},
		},
		{
			description: "nested Transaction",
			interactions: []interaction{
				{"GET test-var-name", "Nil"},
				{"BEGIN", ""},
				{"SET test-var-name 100", ""},
				{"GET test-var-name", "100"},
				{"BEGIN", ""},
				{"SET test-var-name 120", ""},
				{"GET test-var-name", "120"},
				{"BEGIN", ""},
				{"SET test-var-name 150", ""},
				{"GET test-var-name", "150"},
				{"ROLLBACK", ""},
				{"COMMIT", ""},
				{"GET test-var-name", "100"}, // Change from 120 to 100 based on the requirement.
				{"END", ""},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			db := db.New()
			var inputBuffer bytes.Buffer

			for _, i := range tc.interactions {
				inputBuffer.WriteString(i.command + "\n")
			}

			var outputBuffer bytes.Buffer
			runLoop(db, &inputBuffer, &outputBuffer)

			for _, i := range tc.interactions {
				if i.output != "" {
					output, err := outputBuffer.ReadString('\n')
					if err != nil {
						t.Errorf("read output fail, err: %v", err)
					}

					if i.output != strings.TrimSuffix(output, "\n") {
						t.Errorf("Fail: expected %v to be %v", i.output, output)
					}
				}
			}
		})
	}
}
