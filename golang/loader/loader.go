package loader

import (
	"go/build"
	"go/parser"
	"path"
	"path/filepath"

	"github.com/matthewmueller/golly/golang/util"
	"github.com/matthewmueller/golly/stdlib"
	"github.com/pkg/errors"
	"golang.org/x/tools/go/loader"
)

// Load takes full paths to packages and loads them
// e.g. $GOPATH/src/github.com/matthewmueller/golly/
func Load(packages ...string) (program *loader.Program, err error) {
	// defer log.Trace("load").Stop(&err)
	var conf loader.Config

	goSrc, err := util.GoSourcePath()
	if err != nil {
		return nil, err
	}

	// add all the packages as imports
	for _, pkgpath := range packages {
		if filepath.HasPrefix(pkgpath, goSrc) {
			rel, e := filepath.Rel(goSrc, pkgpath)
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
