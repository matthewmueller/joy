package loader

import (
	"go/build"
	"go/parser"
	"os"
	"path"
	"path/filepath"

	"github.com/apex/log"
	"github.com/matthewmueller/golly/compiler/util"
	"github.com/matthewmueller/golly/stdlib"
	"github.com/pkg/errors"
	"golang.org/x/tools/go/loader"
)

// Settings struct
type Settings struct {
	Packages []string
}

// Load the packages (Phase I)
func Load(settings *Settings) (program *loader.Program, err error) {
	defer log.Trace("load").Stop(&err)
	var conf loader.Config

	// TODO: will this work everytime?
	gosrc := path.Join(os.Getenv("GOPATH"), "src")

	// add all the packages as imports
	for _, pkgpath := range settings.Packages {
		if filepath.HasPrefix(pkgpath, gosrc) {
			rel, e := filepath.Rel(gosrc, pkgpath)
			if e != nil {
				return nil, e
			}
			pkgpath = rel
		}

		// support both files and packages
		if path.Ext(pkgpath) == ".go" {
			conf.CreateFromFilenames(path.Dir(pkgpath), pkgpath)
		} else {
			conf.Import(pkgpath)
		}
	}

	// import the runtime by default
	// TODO: there's gotta be a better way to do this
	runtimePkg, err := util.RuntimePath()
	if err != nil {
		return nil, err
	}
	conf.Import(runtimePkg)

	// add comments to the AST
	conf.ParserMode = parser.ParseComments

	// replace stdlib packages with our own
	conf.FindPackage = func(ctxt *build.Context, importPath, fromDir string, mode build.ImportMode) (*build.Package, error) {
		alias, e := stdlib.Supports(importPath)
		if e != nil {
			return nil, e
		}

		// not part of the stdlib
		if alias == "" {
			return ctxt.Import(importPath, fromDir, mode)
		}

		// use the alias
		return ctxt.Import(alias, fromDir, mode)
	}

	// load all the packages
	program, err = conf.Load()
	if err != nil {
		return program, errors.Wrap(err, "unable to load the go package")
	}

	return program, nil
}
