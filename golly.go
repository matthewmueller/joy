package golly

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"

	"github.com/matthewmueller/golly/translator"

	"github.com/matthewmueller/golly/js"
	"github.com/pkg/errors"
)

// CompileFile fn
func CompileFile(file string) (string, error) {
	gosrc, err := ioutil.ReadFile(file)
	if err != nil {
		return "", errors.Wrap(err, "couldn't read file")
	}

	// Parse Go source into a Go AST
	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, file, gosrc, 0)
	if err != nil {
		return "", errors.Wrap(err, "couldn't parse Go source to Go AST")
	}

	// Translate Go AST => JS AST
	program, err := translator.Translate(fset, f)
	if err != nil {
		return "", errors.Wrap(err, "couldn't translate Go AST to JS AST")
	}

	// Assemble JS source from JS AST
	src, err := js.Assemble(program)
	if err != nil {
		return "", errors.Wrap(err, "couldn't assemble JS AST to javascript source")
	}

	return src, nil
}

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
