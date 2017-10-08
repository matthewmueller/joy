package binding

import (
	"go/ast"

	"github.com/matthewmueller/golly/jsast"
)

// Bind expands any statements we find based
// against any of the predefined rules. It's
// top-level because these bindings can span
// multiple lines (statements). Later on,
// we can probably refactor the translator
// to support returning an array of statements
// or expressions and then merging it in where
// it makes sense.
//
// TODO: performance. this is the 3rd traversal of the AST.
// we can speed this up with scope (being able to reach backwards)
// or marking nodes during the inspection phase.
// TODO: better fn name? This was a hard comment to write haha
func Bind(n ast.Stmt) ([]interface{}, error) {
	stmts, err := Promise(n)
	if err != nil {
		return nil, err
	} else if stmts != nil {
		return stmts, nil
	}

	stmts, err = Raw(n)
	if err != nil {
		return nil, err
	} else if stmts != nil {
		return stmts, nil
	}

	return nil, nil
}

func expandJSExpression(n ast.Stmt) ([]interface{}, error) {
	// ast.Print(nil, n)

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
	// log.Infof("hi!")

	x, ok := sel.X.(*ast.Ident)
	if !ok {
		return nil, nil
	}

	if x.Name != "js" {
		return nil, nil
	}

	switch sel.Sel.Name {
	case "Raw":
		Raw(n)
	case "Promise":
		Promise(n)
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
	// start := time.Now()
	stmts, e := jsast.Parse(src)
	if e != nil {
		return nil, e
	}
	// log.Infof("took %s", time.Since(start))

	var ret []interface{}
	for _, stmt := range stmts {
		ret = append(ret, stmt)
	}

	return ret, nil
}
