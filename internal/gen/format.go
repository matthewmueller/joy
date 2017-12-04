package gen

import (
	"bytes"
	"io"
	"io/ioutil"
	"os/exec"

	"github.com/pkg/errors"
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

// FormatAll formats by a path overwriting the original file
func FormatAll(dir string) (err error) {
	cmd := exec.Command("goimports", "-w", dir)

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}
	defer stderr.Close()

	if e := cmd.Start(); e != nil {
		return errors.Wrapf(err, "error running goimports")
	}

	formattingError, err := ioutil.ReadAll(stderr)
	if err != nil {
		return errors.Wrapf(err, "stderr error: %s", formattingError)
	}

	if err := cmd.Wait(); err != nil {
		return errors.Wrapf(err, "error running goimports\n%s", formattingError)
	}

	return nil
}
