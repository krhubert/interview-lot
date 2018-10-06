package shell

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	// default prompt character.
	prompt = "$"

	// builtin functions.
	exitFn    = "exit"
	historyFn = "history"
)

// Shell (lots) handles interactive user input.
type Shell struct {
	stdout  io.Writer
	scanner *bufio.Scanner
	history []string
	isPipe  bool
}

// NewShell creates new shell with standard stdin and stdout.
func NewShell() *Shell {
	fi, _ := os.Stdin.Stat()

	return &Shell{
		stdout:  os.Stdout,
		scanner: bufio.NewScanner(os.Stdin),
		isPipe:  (fi.Mode() & os.ModeCharDevice) == 0,
	}
}

// prompt prints propmt character.
func (s *Shell) prompt() {
	if !s.isPipe {
		fmt.Fprintf(s.stdout, "%s ", prompt)
	}
}

// ReadLine reads line from shell input and returns it.
// It returns io.EOF when reading is end.
func (s *Shell) ReadLine() (string, error) {
	s.prompt()

	for s.scanner.Scan() {
		line := s.scanner.Text()
		switch line {
		case exitFn:
			fmt.Fprintln(s.stdout, "Goodbye!")
			return "", io.EOF
		case historyFn:
			fmt.Fprintln(s.stdout, strings.Join(s.history, "\n"))
			s.prompt()
			continue
		case "":
			s.prompt()
			continue
		}

		s.history = append(s.history, line)
		return line, nil
	}

	if err := s.scanner.Err(); err != nil {
		return "", fmt.Errorf("shell error: %s", err)
	}
	return "", io.EOF
}
