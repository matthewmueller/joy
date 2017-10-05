package binding

import (
	"go/ast"

	"github.com/apex/log"
	"github.com/matthewmueller/golly/jsast"
)

// Raw fn
func Raw(n ast.Stmt) ([]jsast.IStatement, error) {
	log.Infof("got raw!")
	return nil, nil
}
