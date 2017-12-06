// Package mains returns absolute file paths to the packages that are "main"
// packages, from the list of packages given. The list of packages can
// include relative paths, the special "..." Go keyword, etc.
package mains

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	"github.com/apex/log"
	"github.com/matthewmueller/joy/internal/compiler/util"
	"github.com/pkg/errors"
)

// Find mains in a variety of ways
// returning an array of absolute
// paths.
//
// Mostly based on:
// https://github.com/mitchellh/gox/blob/c9740af9c6574448fd48eb30a71f964014c7a837/go.go#L123
func Find(packages []string) ([]string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	goSrc, err := util.GoSourcePath()
	if err != nil {
		return nil, err
	}

	for i, pkg := range packages {
		isVariadic := strings.HasSuffix(pkg, "...")
		pkg = strings.TrimRight(pkg, "/.")

		if filepath.Ext(pkg) != "" {
			pkg = path.Dir(pkg)
		}

		if pkg == "." {
			packages[i] = pkg
			continue
		}

		// absolute => relative
		if pkg[0] == filepath.Separator {
			rel, err := filepath.Rel(pwd, pkg)
			if err != nil {
				return nil, errors.Wrapf(err, "couldn't get relative path")
			}
			pkg = rel
		}

		// relative
		if pkg[0] == '.' && pkg[1] == filepath.Separator {
			if isVariadic {
				pkg += "/..."
			}
			packages[i] = pkg
			continue
		}

		pkg = "." + string(filepath.Separator) + pkg
		if isVariadic {
			pkg += "/..."
		}
		packages[i] = pkg
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
			// TODO: not sure if this is reliable
			// but it's for when you pass a filepath
			// to go list, it returns command-line-arguments
			if parts[1] == "command-line-arguments" {
				// TODO: i don't think this will work for multiple files
				for _, pkg := range packages {
					results = append(results, pkg)
				}
				return results, nil
			}

			fullpath := strings.TrimPrefix(parts[1], "_")
			if fullpath[0] != '/' {
				fullpath = path.Join(goSrc, fullpath)
			}

			results = append(results, fullpath)
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
