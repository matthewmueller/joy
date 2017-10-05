package binding

import (
	"go/ast"
	"time"

	"github.com/apex/log"
	"github.com/matthewmueller/golly/jsast"
)

// Raw fn
func Raw(n ast.Stmt) ([]interface{}, error) {
	xs, ok := n.(*ast.ExprStmt)
	if !ok {
		return nil, nil
	}

	cx, ok := xs.X.(*ast.CallExpr)
	if !ok {
		return nil, nil
	}

	sel, ok := cx.Fun.(*ast.SelectorExpr)
	if !ok {
		return nil, nil
	}

	x, ok := sel.X.(*ast.Ident)
	if !ok {
		return nil, nil
	}

	if x.Name != "js" || sel.Sel.Name != "Raw" {
		return nil, nil
	}

	if len(cx.Args) == 0 {
		return nil, nil
	}

	lit, ok := cx.Args[0].(*ast.BasicLit)
	if !ok {
		return nil, nil
	}

	// TODO: ensure this is point to the correct js.Raw
	// id := info.ObjectOf(x)
	// _ = id

	src := lit.Value[1 : len(lit.Value)-1]
	start := time.Now()
	stmts, e := jsast.Parse(src)
	if e != nil {
		return nil, e
	}
	log.Infof("took %s", time.Since(start))

	var ret []interface{}
	for _, stmt := range stmts {
		ret = append(ret, stmt)
	}

	return ret, nil
}
