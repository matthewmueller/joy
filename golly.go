package golly

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/matthewmueller/golly/golang"
	"github.com/pkg/errors"
)

// Compile compiles the package
func Compile(path string) (string, error) {
	p, e := normalize(path)
	if e != nil {
		return "", e
	}

	return golang.CompilePackage(p)
}

// CompileFile compiles a single file
func CompileFile(path string) (string, error) {
	p, e := normalize(path)
	if e != nil {
		return "", e
	}

	return golang.CompilePackage(p)
}

// CompileString a source string
func CompileString(path, source string) (string, error) {
	return "", nil
}

// CallGraph fn
func CallGraph(path string) error {
	p, e := normalize(path)
	if e != nil {
		return e
	}

	golang.CallGraph(p)
	return nil
}

// support relative and absolute paths
// TODO: fix crappy code, there's gotta
// be an easier way to do this
func normalize(pkgpath string) (string, error) {
	if !filepath.IsAbs(pkgpath) {
		cwd, e := os.Getwd()
		if e != nil {
			return "", e
		}
		pkgpath = filepath.Join(cwd, pkgpath)
	}

	path, e := filepath.Rel(os.Getenv("GOPATH")+"/src", pkgpath)
	if e != nil {
		return "", e
	}

	return path, nil
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
