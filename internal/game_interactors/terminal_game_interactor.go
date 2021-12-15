package game_interactors

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type TerminalGameInteractor struct {
	reader *bufio.Reader
}

func NewTerminalGameInteractor() *TerminalGameInteractor {
	return &TerminalGameInteractor{reader: bufio.NewReader(os.Stdin)}
}

func (t *TerminalGameInteractor) GetString(prompt string) (string, error) {
	fmt.Println(prompt)

	// Reading from stdin
	str, err := t.reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	// Removing "\n" from the end of the string and returning it
	return str[:len(str)-1], nil
}

func (t *TerminalGameInteractor) GetNumber(prompt string) (int, error) {
	// Getting string from stdin
	str, err := t.GetString(prompt)
	if err != nil {
		return 0, err
	}

	// Converting string to int
	number, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	return number, nil
}

func (t *TerminalGameInteractor) Say(message string) {
	fmt.Println(message)
}
