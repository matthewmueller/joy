package loader

import (
	"fmt"
	"go/build"
	"go/parser"
	"path"
	"path/filepath"
	"strings"

	"github.com/asaskevich/govalidator"

	"github.com/matthewmueller/joy/internal/std"
	"github.com/pkg/errors"
	"golang.org/x/tools/go/loader"
)

// Config struct
type Config struct {
	Packages []string `valid:"required"`
	JoyPath  string   `valid:"required"`
}

// Load takes full paths to packages and loads them
// e.g. $GOPATH/src/github.com/matthewmueller/joy/
func Load(cfg *Config) (program *loader.Program, err error) {
	if valid, err := govalidator.ValidateStruct(cfg); !valid || err != nil {
		return program, errors.Wrapf(err, "invalid loader config")
	}

	var conf loader.Config

	// paths
	gosrc := path.Join(build.Default.GOPATH, "src")
	stdPath := path.Join(cfg.JoyPath, "stdlib")
	macroPath := path.Join(cfg.JoyPath, "macro")
	runtimePath := path.Join(cfg.JoyPath, "internal", "runtime")

	// add comments to the AST
	conf.ParserMode = parser.ParseComments

	// We implement a custom package resolution here,
	// because we need to inject our own stdlib and
	// macro system here.
	//
	// We also want to make certain things just work
	// without having Go setup. Code that uses
	// Joy's runtime (e.g. for channel support)
	// macros, and stdlib shouldn't need Go setup.
	//
	// This is similar to how Go works without $GOPATH
	// setup.
	conf.FindPackage = func(ctxt *build.Context, importPath, fromDir string, mode build.ImportMode) (*build.Package, error) {
		// handle stdlib
		if std.In(importPath) {
			if !std.Supported(importPath) {
				return nil, fmt.Errorf("%s is not supported yet by Joy", importPath)
			}

			rel, err := filepath.Rel(gosrc, path.Join(stdPath, importPath))
			if err != nil {
				return nil, errors.Wrapf(err, "error getting relative path for the stdlib")
			}

			return ctxt.Import(rel, gosrc, mode)
		}

		// know where to find joy/macro, so we don't need to Go
		// installed just to compile basic programs
		if strings.HasSuffix(importPath, "joy/macro") {
			rel, err := filepath.Rel(gosrc, macroPath)
			if err != nil {
				return nil, errors.Wrapf(err, "error getting relative path for the stdlib")
			}
			return ctxt.Import(rel, gosrc, mode)
		}

		// the default loader's importer errors out on absolute
		// paths, so we'll take the relative path in that case
		// TODO: we can probably add this later, but the way
		// our paths get written right now, they have to be
		// consistent in the import path
		// if importPath[0] == filepath.Separator {
		// 	rel, err := filepath.Rel(gosrc, importPath)
		// 	if err != nil {
		// 		return nil, errors.Wrapf(err, "error getting relative import path")
		// 	}
		// 	return ctxt.Import(rel, gosrc, mode)
		// }

		// import the regular way
		return ctxt.Import(importPath, gosrc, mode)
	}

	// include the runtime by default
	rel, err := filepath.Rel(gosrc, runtimePath)
	if err != nil {
		return nil, errors.Wrapf(err, "error getting relative path for the runtime")
	}
	conf.Import(rel)

	// import all the packages
	for _, pkgPath := range cfg.Packages {
		rel, err := filepath.Rel(gosrc, pkgPath)
		if err != nil {
			return nil, errors.Wrapf(err, "not relative to $GOPATH")
		}
		conf.Import(rel)
	}

	// load all the packages
	program, err = conf.Load()
	if err != nil {
		return program, errors.Wrap(err, "unable to load the go package")
	}

	return program, nil
}
