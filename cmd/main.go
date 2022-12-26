package main

import (
	"bufio"
	"fmt"
	"in_memory_db/internal/db"
	"in_memory_db/internal/interpreter"
	"io"
	"os"
	"strings"
)

func main() {
	db := db.New()
	runLoop(db, os.Stdin, os.Stdout)
}

func runLoop(db db.DB, input io.Reader, output io.Writer) {
	fmt.Println(" > type END to exit")

	scanner := bufio.NewScanner(input)

	var line string
	for {
		scanner.Scan()
		line = scanner.Text()
		if strings.EqualFold(line, "END") {
			break
		}

		cmdOutput, err := interpreter.Run(db, line)
		if err != nil {
			output.Write([]byte(fmt.Sprintf("Fail on execute `%v`, error: %v\n", line, err)))
			continue
		}

		if len(cmdOutput) > 0 {
			cmdOutput += "\n"
		}

		output.Write([]byte(cmdOutput))
	}

	fmt.Println("see you soon")
}
