package shell

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func newShellWithInput(input string) (*bytes.Buffer, io.ReadCloser, *Shell) {
	var (
		out = new(bytes.Buffer)
		in  = ioutil.NopCloser(bytes.NewReader([]byte(input)))
		s   = &Shell{
			stdout:  out,
			scanner: bufio.NewScanner(in),
		}
	)
	return out, in, s
}

func TestNewShell(t *testing.T) {
	s := NewShell()
	if s.stdout != os.Stdout {
		t.Errorf("shell output should be set to os.Stdout")
	}
	if s.isPipe {
		t.Errorf("pipe should be false")
	}
}

func TestPrintPrompt(t *testing.T) {
	out, _, s := newShellWithInput("test\n")

	if _, err := s.ReadLine(); err != nil {
		t.Errorf("read line error: %s", err)
	}
	if !strings.HasPrefix(out.String(), "$") {
		t.Error("invalid prompt character")
	}
}

func TestReadLine(t *testing.T) {
	_, in, s := newShellWithInput("test\n")
	in.Close()

	line, err := s.ReadLine()
	if err != nil {
		t.Errorf("read line error: %s", err)
	}

	if line != "test" {
		t.Errorf("read line - want: %s, got: %s", "test", line)
	}
}

func TestCloseInput(t *testing.T) {
	_, in, s := newShellWithInput("test\n")
	in.Close()

	// read in buffer
	s.ReadLine()
	if _, err := s.ReadLine(); err != io.EOF {
		t.Errorf("expected close but got: %s", err)
	}
}

func TestExitFn(t *testing.T) {
	_, _, s := newShellWithInput("exit\n")
	if _, err := s.ReadLine(); err != io.EOF {
		t.Errorf("expected close but got: %s", err)
	}
}

func TestHisotryFn(t *testing.T) {
	out, _, s := newShellWithInput("1\n2\nhistory\n")

	s.ReadLine()
	s.ReadLine()
	s.ReadLine()

	if line := out.String(); strings.HasSuffix(line, "1\n2\n") {
		t.Errorf("history failed - want: %s, got: %s", "1\n2\n", line)
	}
}
