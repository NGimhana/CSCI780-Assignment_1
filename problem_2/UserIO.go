package main

import (
	"bufio"
	"fmt"
	"os"
)

type UserIO struct {
	message string
}

func (userio *UserIO) print(message string) {
	fmt.Printf(message)
}

func (userio *UserIO) input(message string) string {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		message = scanner.Text()
	}
	return message
}
