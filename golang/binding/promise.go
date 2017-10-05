package binding

import (
	"go/ast"

	"github.com/apex/log"
	"github.com/matthewmueller/golly/jsast"
)

// Promise fn
func Promise(n ast.Stmt) ([]jsast.IStatement, error) {
	log.Infof("got promise!")
	return nil, nil
}
