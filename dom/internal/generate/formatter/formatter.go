package formatter

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"os/exec"
)

// Format the output using goimports
func Format(input string) (output string, err error) {
	cmd := exec.Command("goimports")
	stdin, err := cmd.StdinPipe()

	if err != nil {
		return output, err
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return output, err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return output, err
	}

	reader := bytes.NewBufferString(input)

	if e := cmd.Start(); e != nil {
		return output, e
	}

	io.Copy(stdin, reader)
	stdin.Close()

	formatted, err := ioutil.ReadAll(stdout)
	if err != nil {
		return output, err
	}

	formattingError, err := ioutil.ReadAll(stderr)
	if err != nil {
		return output, err
	}

	stderr.Close()
	stdout.Close()

	if e := cmd.Wait(); e != nil {
		return output, errors.New(string(formattingError))
	}

	return string(formatted), nil
}
