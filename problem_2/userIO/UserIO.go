package userIO

import (
	"bufio"
	"fmt"
	"os"
)

type UserIO struct {
	message string
}

func (userio *UserIO) Print(message string) {
	fmt.Printf(message)
}

func (userio *UserIO) Input(message string) string {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		message = scanner.Text()
	}
	return message
}
