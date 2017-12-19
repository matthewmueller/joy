// Package mains returns absolute file paths to the packages that are "main"
// packages, from the list of packages given. The list of packages can
// include relative paths, the special "..." Go keyword, etc.
package mains

import (
	"bytes"
	"fmt"
	"go/build"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	"github.com/apex/log"
	"github.com/pkg/errors"
)

// Find mains in a variety of ways
// returning an array of absolute
// paths.
//
// Mostly based on:
// https://github.com/mitchellh/gox/blob/c9740af9c6574448fd48eb30a71f964014c7a837/go.go#L123
func Find(packages ...string) (mains []string, err error) {
	cwd, err := os.Getwd()
	if err != nil {
		return mains, err
	}

	gosrc := path.Join(build.Default.GOPATH, "src")

	for i, pkg := range packages {
		// no prefix
		if pkg[0] != '.' && !filepath.IsAbs(pkg) {
			packages[i] = "." + string(filepath.Separator) + pkg
			continue
		}

		// absolute
		if filepath.IsAbs(pkg) {
			rel, err := filepath.Rel(cwd, pkg)
			if err != nil {
				return mains, errors.Wrapf(err, "unable to get relative path")
			}

			if rel[0] != '.' {
				packages[i] = "." + string(filepath.Separator) + rel
			} else {
				packages[i] = rel
			}

			continue
		}
	}

	goCmd, err := exec.LookPath("go")
	if err != nil {
		return nil, err
	}

	args := make([]string, 0, len(packages)+3)
	args = append(args, "list", "-f", "{{.Name}}|{{.ImportPath}}")
	args = append(args, packages...)
	output, err := execGo(goCmd, nil, "", args...)
	if err != nil {
		return nil, err
	}

	results := make([]string, 0, len(output))
	for _, line := range strings.Split(output, "\n") {
		if line == "" {
			continue
		}

		parts := strings.SplitN(line, "|", 2)
		if len(parts) != 2 {
			log.Warnf("Bad line reading packages: %s", line)
			continue
		}

		if parts[0] == "main" {
			if parts[1] == "command-line-arguments" {
				// TODO: i don't think this will work for multiple files
				for _, pkg := range packages {
					results = append(results, path.Join(cwd, filepath.Dir(pkg)))
				}
				continue
			}

			if strings.HasPrefix(parts[1], "_/") {
				results = append(results, strings.TrimPrefix(parts[1], "_"))
				continue
			}

			// TODO: this needs testing, i think the tests never
			// reach this because they're in ./testdata
			results = append(results, path.Join(gosrc, parts[1]))
		}
	}

	return results, nil
}

// Execute go in a separate processs
func execGo(GoCmd string, env []string, dir string, args ...string) (string, error) {
	var stderr, stdout bytes.Buffer
	cmd := exec.Command(GoCmd, args...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if env != nil {
		cmd.Env = env
	}
	if dir != "" {
		cmd.Dir = dir
	}
	if err := cmd.Run(); err != nil {
		err = fmt.Errorf("%s\nStderr: %s", err, stderr.String())
		return "", err
	}
	return stdout.String(), nil
}
