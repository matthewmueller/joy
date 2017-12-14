package setup

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/matthewmueller/joy/internal/bindata"
	"github.com/matthewmueller/joy/internal/paths"
	"github.com/matthewmueller/joy/internal/version"
	"github.com/pkg/errors"
)

// Runtime tests and sets up the right place to pull in runtime
// source files. This will depend on if you're
// working from the $GOPATH or if you've simply
// run the binary
func Runtime() error {
	runtimePath, err := paths.Runtime()
	if err != nil {
		return errors.Wrapf(err, "error getting runtime path")
	}

	if exists, err := exists(runtimePath); exists && err != nil {
		return nil
	}

	return setup(runtimePath)
}

// Stdlib tests and sets up the right place to pull in stdlib
// source files. This will depend on if you're
// working from the $GOPATH or if you've simply
// run the binary
func Stdlib() error {
	stdlibPath, err := paths.Stdlib()
	if err != nil {
		return errors.Wrapf(err, "error getting runtime path")
	}

	if exists, err := exists(stdlibPath); exists && err != nil {
		return nil
	}

	return setup(stdlibPath)
}

// Macro sets up the macro path source files.
// This will depend on if you're working from
// the $GOPATH or if you've simply run the binary
func Macro() error {
	macroPath, err := paths.Macro()
	if err != nil {
		return errors.Wrapf(err, "error getting runtime path")
	}

	if exists, err := exists(macroPath); exists && err != nil {
		return nil
	}

	return setup(macroPath)
}

func exists(dir string) (bool, error) {
	// exists
	if _, err := os.Stat(dir); !os.IsNotExist(err) {
		ver, err := ioutil.ReadFile(path.Join(dir, "VERSION"))
		if err != nil {
			return false, errors.Wrapf(err, "unable to check version")
		}

		// if the versions match, there's nothing to do
		if string(ver) == version.Version {
			return true, nil
		}
	}
	return false, nil
}

func setup(dir string) error {
	root, err := paths.Joy()
	if err != nil {
		return errors.Wrapf(err, "error getting root joy")
	}

	rel, err := filepath.Rel(root, dir)
	if err != nil {
		return errors.Wrapf(err, "error getting relative path")
	}

	for _, name := range bindata.AssetNames() {
		if path.Base(name) == "VERSION" {
			continue
		} else if !strings.HasPrefix(name, rel) {
			continue
		}

		data, err := bindata.Asset(name)
		if err != nil {
			return errors.Wrapf(err, "error getting asset data for %s", name)
		}

		fullpath := path.Join(root, name)
		if err := os.MkdirAll(path.Dir(fullpath), 0755); err != nil {
			return errors.Wrapf(err, "error creating dir")
		}

		if err := ioutil.WriteFile(fullpath, data, 0644); err != nil {
			return errors.Wrapf(err, "error writing asset")
		}
	}

	if err := ioutil.WriteFile(path.Join(dir, "VERSION"), []byte(version.Version), 0644); err != nil {
		return errors.Wrapf(err, "error writing version file")
	}

	return nil
}
