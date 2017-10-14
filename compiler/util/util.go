// Package util is a grab bag of utility functions
//
// TODO: go through and organize these better and
// eventually remove this file
package util

import (
	"errors"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

var runtimePkg string

// RuntimePath gets the path of our runtime
func RuntimePath() (string, error) {
	if runtimePkg != "" {
		return runtimePkg, nil
	}

	_, file, _, ok := runtime.Caller(0)
	if !ok {
		return "", errors.New("unable to get the filepath")
	}

	root := path.Dir(path.Dir(path.Dir(file)))
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		return "", errors.New("GOPATH must be set in your path")
	}
	gosrc := path.Join(gopath, "src")
	runtimePath := path.Join(root, "runtime")

	// runtime package
	rt, e := filepath.Rel(gosrc, runtimePath)
	if e != nil {
		return "", e
	}
	runtimePkg = rt

	return runtimePkg, nil
}
