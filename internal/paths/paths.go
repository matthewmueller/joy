package paths

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/pkg/errors"
)

// cached paths
var joyPath string
var runtimePath string
var stdlibPath string
var macroPath string

// Joy is root the directory where Joy stores source,
// state and analytics. This will first check the $GOPATH,
// but fallback to an OS-specific directory.
//
// TODO: should we use go-bindata and just embed
// the runtime and headless chrome?
//
// Based on: https://github.com/sindresorhus/env-paths
func Joy() (string, error) {
	// use cached
	if joyPath != "" {
		return joyPath, nil
	}

	_, file, _, ok := runtime.Caller(0)
	if !ok {
		return "", fmt.Errorf("error getting caller")
	}

	// short-circuit for testing
	// return getPath("joy")

	root := path.Join(file, "..", "..", "..")
	if _, err := os.Stat(root); !os.IsNotExist(err) {
		joyPath = root
		return root, nil
	}

	return getPath("joy")
}

// Runtime path
func Runtime() (string, error) {
	// use cached
	if runtimePath != "" {
		return runtimePath, nil
	}

	root, err := Joy()
	if err != nil {
		return "", errors.Wrapf(err, "error getting runtime")
	}

	runtimePath = path.Join(root, "internal", "runtime")
	return runtimePath, nil
}

// Stdlib path
func Stdlib() (string, error) {
	// use cached
	if stdlibPath != "" {
		return stdlibPath, nil
	}

	root, err := Joy()
	if err != nil {
		return "", errors.Wrapf(err, "error getting stdlib")
	}

	stdlibPath = path.Join(root, "stdlib")
	return stdlibPath, nil
}

// Macro gets the macro path
//
// Note that this depends on $GOPATH
// since this will be imported as
// a library, so it will depend on
// the go toolchain being setup
//
// TODO: really hope this can go away eventually
func Macro() (string, error) {
	// use cached
	if macroPath != "" {
		return macroPath, nil
	}

	root, err := Joy()
	if err != nil {
		return "", errors.Wrapf(err, "error getting stdlib")
	}

	macroPath = path.Join(root, "macro")
	return macroPath, nil

	// gopath := envOr("GOPATH", defaultGOPATH())
	// if gopath == "" {
	// 	return "", errors.New("no gopath set")
	// }

	// fullpath := path.Join(gopath, "src", "github.com", "matthewmueller", "joy", "macro")
	// if _, err := os.Stat(fullpath); os.IsNotExist(err) {
	// 	return "", errors.New("you'll need to run `go get github.com/matthewmueller/joy/macro` to compile code with macros")
	// }

	// // TODO: fragile, fix for forks and things
	// macroPath = "github.com/matthewmueller/joy/macro"
	// return macroPath, nil
}

// // RuntimePath gets the path of our runtime
// func RuntimePath() (string, error) {
// 	if runtimePkg != "" {
// 		return runtimePkg, nil
// 	}

// 	root, e := JoyPath()
// 	if e != nil {
// 		return "", e
// 	}
// 	runtimePath := path.Join(root, "internal", "runtime")

// 	gosrc, e := GoSourcePath()
// 	if e != nil {
// 		return "", e
// 	}

// 	// runtime package
// 	rt, e := filepath.Rel(gosrc, runtimePath)
// 	if e != nil {
// 		return "", e
// 	}
// 	runtimePkg = rt

// 	return runtimePkg, nil
// }

// get the path to the storage
func getPath(paths ...string) (p string, err error) {
	home, err := homedir.Dir()
	if err != nil {
		return p, err
	}

	switch runtime.GOOS {
	case "darwin":
		ps := append([]string{home, "Library", "Preferences"}, paths...)
		return path.Join(ps...), err
	case "linux":
		base := os.Getenv("XDG_CONFIG_HOME")
		if base == "" {
			base = path.Join(home, ".config")
		}
		ps := append([]string{base}, paths...)
		return path.Join(ps...), err
	case "windows":
		appdata := os.Getenv("LOCALAPPDATA")
		if appdata == "" {
			appdata = path.Join(home, "AppData", "Local")
		}
		ps := append([]string{appdata}, paths...)
		ps = append(ps, "Config")
		return path.Join(ps...), err
	default:
		return p, errors.New("store does not yet support " + runtime.GOOS + ". Please open a pull request!")
	}
}

func envOr(name, def string) string {
	s := os.Getenv(name)
	if s == "" {
		return def
	}
	return s
}

// pulled from go source code
func defaultGOPATH() string {
	env := "HOME"
	if runtime.GOOS == "windows" {
		env = "USERPROFILE"
	} else if runtime.GOOS == "plan9" {
		env = "home"
	}
	if home := os.Getenv(env); home != "" {
		def := filepath.Join(home, "go")
		if filepath.Clean(def) == filepath.Clean(runtime.GOROOT()) {
			// Don't set the default GOPATH to GOROOT,
			// as that will trigger warnings from the go tool.
			return ""
		}
		return def
	}
	return ""
}
