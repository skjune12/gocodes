package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("> ")

		// Read the keyboard input
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		// Handle the execution of the input
		err = execInput(input)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func execInput(input string) error {
	// Remove newline character.
	input = strings.TrimSuffix(input, "\n")

	// Split the input to separate the command and the arguments.
	args := strings.Split(input, " ")

	switch args[0] {
	case "cd":
		// 'cd' to homedir with empty path not yet supported
		if len(args) < 2 {
			return errors.New("Path required")
		}
		err := os.Chdir(args[1])
		if err != nil {
			return err
		}
		// Stop further processing
		return nil
	case "exit":
		os.Exit(0)
	}

	// Pass the program and the arguments separately.
	cmd := exec.Command(args[0], args[1:]...)

	// Execute the command and save it's output.
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	// Print the output
	fmt.Printf("%s", stdoutStderr)
	return nil
}
