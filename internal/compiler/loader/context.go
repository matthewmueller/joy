package loader

import (
	"go/build"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

// defaultContext comes from the standard library
func defaultContext() build.Context {
	var c build.Context

	c.GOARCH = envOr("GOARCH", runtime.GOARCH)
	c.GOOS = envOr("GOOS", runtime.GOOS)
	c.GOROOT = path.Clean(runtime.GOROOT())
	c.GOPATH = envOr("GOPATH", defaultGOPATH())
	c.Compiler = runtime.Compiler
	c.ReleaseTags = build.Default.ReleaseTags
	c.CgoEnabled = build.Default.CgoEnabled
	c.BuildTags = build.Default.BuildTags[:]

	return c
}

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

func envOr(name, def string) string {
	s := os.Getenv(name)
	if s == "" {
		return def
	}
	return s
}
