package game_interactors

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"syscall"
)

type TerminalGameInteractor struct {
	reader *bufio.Reader
}

func NewTerminalGameInteractor() *TerminalGameInteractor {
	return &TerminalGameInteractor{reader: bufio.NewReader(os.Stdin)}
}

func (t *TerminalGameInteractor) GetString(prompt string) (string, error) {
	fmt.Println(prompt)

	// Common settings and variables for both stty calls.
	attrs := syscall.ProcAttr{
		Dir:   "",
		Env:   []string{},
		Files: []uintptr{os.Stdin.Fd(), os.Stdout.Fd(), os.Stderr.Fd()},
		Sys:   nil,
	}
	var ws syscall.WaitStatus

	// Disable echoing.
	pid, err := syscall.ForkExec(
		"/bin/stty",
		[]string{"stty", "-echo"},
		&attrs,
	)
	if err != nil {
		panic(err)
	}

	// Wait for the stty process to complete.
	_, err = syscall.Wait4(pid, &ws, 0, nil)
	if err != nil {
		panic(err)
	}

	// Reading from stdin
	str, err := t.reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	// Re-enable echo.
	pid, err = syscall.ForkExec(
		"/bin/stty",
		[]string{"stty", "echo"},
		&attrs)
	if err != nil {
		panic(err)
	}

	// Wait for the stty process to complete.
	_, err = syscall.Wait4(pid, &ws, 0, nil)
	if err != nil {
		panic(err)
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
