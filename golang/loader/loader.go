package loader

import (
	"go/build"
	"go/parser"
	"path"
	"path/filepath"

	"github.com/apex/log"
	"github.com/matthewmueller/golly/golang/util"
	"github.com/matthewmueller/golly/stdlib"
	"github.com/pkg/errors"
	"golang.org/x/tools/go/loader"
)

// Load takes full paths to packages and loads them
// e.g. $GOPATH/src/github.com/matthewmueller/golly/
func Load(packages ...string) (program *loader.Program, err error) {
	defer log.Trace("load").Stop(&err)
	var conf loader.Config

	goSrc, err := util.GoSourcePath()
	if err != nil {
		return nil, err
	}

	// add all the packages as imports
	for _, fullpath := range packages {
		packagePath, e := filepath.Rel(goSrc, fullpath)
		if e != nil {
			return nil, e
		}

		// support both files and packages
		if path.Ext(fullpath) == ".go" {
			log.Debugf("path=%s filename=%s", path.Dir(packagePath), fullpath)
			conf.CreateFromFilenames(path.Dir(packagePath), fullpath)
		} else {
			log.Debugf("pkgpath=%s", packagePath)
			conf.Import(packagePath)
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
