package golly

import (
	"fmt"
	"go/ast"
	"go/build"
	"go/parser"
	"go/token"
	"io/ioutil"

	"golang.org/x/tools/go/loader"

	"github.com/apex/log"

	"github.com/matthewmueller/golly/golang"
	"github.com/pkg/errors"
)

// Compile the package, file or source
func Compile(path string, source string) (string, error) {
	if e := compilePackage(path); e != nil {
		return "", e
	}
	return "", nil
}

// compilePackage compiles a package by it's path
func compilePackage(packagePath string) error {
	var conf loader.Config
	conf.Import(packagePath)

	// load each of the imports
	conf.FindPackage = func(ctx *build.Context, importPath, fromDir string, mode build.ImportMode) (*build.Package, error) {
		fmt.Println("importing", importPath)
		return ctx.Import(importPath, fromDir, mode)
	}

	// load the package
	pkgs, err := conf.Load()
	if err != nil {
		return errors.Wrap(err, "unable to load the go package")
	}

	// translate each file to their Javascript counterpart
	for pkg, info := range pkgs.AllPackages {
		log.Infof("building... %s", pkg.Name())
		for _, file := range info.Files {
			log.Infof("translating... %s", file.Name.Name)
			ast, e := golang.Translate(info, file)
			if e != nil {
				return errors.Wrapf(e, "error translating %s", file.Name.Name)
			}
			log.Infof("translated: %s\n%s", file.Name.Name, ast)
		}
	}

	return nil
}

// // Compile source
// func Compile(src string) (string, error) {
// 	return compile("stdin.js", src)
// }

// // CompileFile fn
// func CompileFile(file string) (string, error) {
// 	src, err := ioutil.ReadFile(file)
// 	if err != nil {
// 		return "", errors.Wrap(err, "couldn't read file")
// 	}

// 	return compile(file, string(src))
// }

// // Compile source
// func compile(file, src string) (string, error) {
// 	// Parse Go source into a Go AST
// 	fset := token.NewFileSet() // positions are relative to fset
// 	f, err := parser.ParseFile(fset, file, []byte(src), 0)
// 	if err != nil {
// 		return "", errors.Wrap(err, "couldn't parse Go source to Go AST")
// 	}

// 	// Translate Go AST => JS AST
// 	program, err := golang.Translate(fset, f)
// 	if err != nil {
// 		return "", errors.Wrap(err, "couldn't translate Go AST to JS AST")
// 	}

// 	// Assemble JS source from JS AST
// 	javascript := js.Assemble(program)

// 	return javascript, nil
// }

// PrintAST fn
func PrintAST(file string) error {
	gosrc, err := ioutil.ReadFile(file)
	if err != nil {
		return errors.Wrap(err, "couldn't read file")
	}

	// Parse Go source into a Go AST
	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, file, gosrc, 0)
	if err != nil {
		return errors.Wrap(err, "couldn't parse Go source to Go AST")
	}

	ast.Print(fset, f)
	return nil
}
