package variable

import (
	"fmt"
	"go/ast"
	"reflect"

	"github.com/matthewmueller/golly/jsast"
	"github.com/pkg/errors"
)

// TODO: this is not in use yet. I'm not sure how best to integrate
// it at the moment, since variables can be initialized in a multitude
// of ways. The API here may change to just structs containing the data
// so we can turn it into a JS AST however we see fit. Once this is in
// place we should be able to take care of variables in every way that
// Go can.

// Handle function
//
// Unfortunately, this process has lots of interleaving parts to it.
// That being said, this code is pretty crappy and there's lots that
// could be improved.
func Handle(n interface{}) (jsast.IStatement, jsast.IExpression, error) {
	switch t := n.(type) {
	case *ast.GenDecl:
		return genDeclPairs(t)
	case *ast.AssignStmt:
		return assignStmtPairs(t)
	default:
		return nil, nil, unhandled("variablePairs", t)
	}
}

func genDeclPairs(n *ast.GenDecl) (jsast.IStatement, jsast.IExpression, error) {
	var decls []jsast.VariableDeclarator

	// ast.Inspect
	for _, spec := range n.Specs {
		switch t := spec.(type) {
		case *ast.ValueSpec:
			vars, e := valueSpecHandler(t)
			if e != nil {
				return nil, nil, e
			}
			decls = append(decls, vars...)
		}
	}

	if len(decls) == 0 {
		return nil, nil, nil
	}

	return jsast.CreateVariableDeclaration("var", decls...), nil, nil
}

func valueSpecHandler(n *ast.ValueSpec) (decls []jsast.VariableDeclarator, e error) {
	lv := len(n.Values)

	for i, name := range n.Names {
		var val jsast.IExpression

		if i < lv {
			val, e = getValue(n.Values[i])
		} else {
			val, e = getEmpty(n.Type)
		}
		if e != nil {
			return decls, e
		}

		decls = append(decls, jsast.CreateVariableDeclarator(
			jsast.CreateIdentifier(name.Name),
			val,
		))
	}

	return decls, nil
}

func assignStmtPairs(n *ast.AssignStmt) (jsast.IStatement, jsast.IExpression, error) {
	var decls []jsast.VariableDeclarator
	var exprs []jsast.IExpression
	op := n.Tok.String()
	lrh := len(n.Rhs)
	llh := len(n.Lhs)
	destructured := lrh > 0 && lrh < llh

	if len(n.Lhs) == 0 {
		return nil, nil, errors.New("invalid assumption that assignment statements are always len(n.Lhs) > 0")
	}

	nth, e := firstNonBlank(n.Lhs)
	if e != nil {
		return nil, nil, e
	} else if nth < 0 {
		for _, rh := range n.Rhs {
			x, e := getValue(rh)
			if e != nil {
				return nil, nil, e
			} else if x != nil {
				exprs = append(exprs, x)
			}
		}
		return nil, jsast.CreateSequenceExpression(exprs...), nil
	}

	for i, l := range n.Lhs {
		var r ast.Expr
		if i < lrh {
			r = n.Rhs[i]
		}

		var d *jsast.VariableDeclarator
		var x jsast.IExpression
		var e error

		switch op {
		case ":=":
			if destructured {
				if i == 0 {
					d, e = assignStmtVars(n.Lhs[nth], r, "$", "", i)
					if e != nil {
						return nil, nil, e
					} else if d != nil {
						decls = append(decls, *d)
					}
				}
				d, e = assignStmtVars(l, n.Lhs[nth], "", "$", i)
			} else {
				d, e = assignStmtVars(l, r, "", "", i)
			}
			if d != nil {
				decls = append(decls, *d)
			}
		case "=":
			if destructured {
				if i == 0 {
					d, e = assignStmtVars(n.Lhs[nth], r, "$", "", i)
					if e != nil {
						return nil, nil, e
					} else if d != nil {
						decls = append(decls, *d)
					}
				}
				x, e = assignStmtExprs(l, n.Lhs[nth], "", "$", i)
			} else {
				x, e = assignStmtExprs(l, r, "", "", i)
			}
			if x != nil {
				exprs = append(exprs, x)
			}
		}
		if e != nil {
			return nil, nil, e
		}

	}

	if len(decls) > 0 && len(exprs) > 0 {
		return jsast.CreateVariableDeclaration("var", decls...),
			jsast.CreateSequenceExpression(exprs...),
			nil
	}

	if len(decls) > 0 {
		return jsast.CreateVariableDeclaration("var", decls...),
			nil,
			nil
	}

	if len(exprs) > 0 {
		return nil,
			jsast.CreateSequenceExpression(exprs...),
			nil
	}

	return nil, nil, nil
}

func assignStmtVars(l ast.Expr, r ast.Expr, prefixLeft, prefixRight string, nth int) (*jsast.VariableDeclarator, error) {
	jl, e := getPattern(l, prefixLeft)
	if e != nil {
		return nil, e
	} else if jl == nil {
		return nil, nil
	}

	var jr jsast.IExpression
	if prefixRight == "" {
		jr, e = getValue(r)
	} else {
		jr, e = getArrayValue(r, prefixRight, nth)
	}
	if e != nil {
		return nil, e
	}

	decl := jsast.CreateVariableDeclarator(jl, jr)
	return &decl, nil
}

func assignStmtExprs(l ast.Expr, r ast.Expr, prefixLeft, prefixRight string, nth int) (jsast.IExpression, error) {
	jl, e := getExpression(l, prefixLeft)
	if e != nil {
		return nil, e
	} else if jl == nil {
		return nil, nil
	}

	var jr jsast.IExpression
	if prefixRight == "" {
		jr, e = getValue(r)
	} else {
		jr, e = getArrayValue(r, prefixRight, nth)
	}
	if e != nil {
		return nil, e
	}

	return jsast.CreateAssignmentExpression(
		jl,
		jsast.AssignmentOperator("="),
		jr,
	), nil
}

func firstNonBlank(ns []ast.Expr) (int, error) {
	for i, n := range ns {
		l, e := getPatternName(n)
		if e != nil {
			return 0, e
		}

		if l != "_" {
			return i, nil
		}
	}

	return -1, nil
}

func getPattern(n ast.Expr, prefix string) (jsast.IPattern, error) {
	switch t := n.(type) {
	case *ast.Ident:
		if t.Name == "_" {
			return nil, nil
		}
		return jsast.CreateIdentifier(prefix + t.Name), nil
	case *ast.SelectorExpr:
		return getPattern(t.X, prefix)
	default:
		return nil, unhandled("getPattern()", n)
	}
}

func getExpression(n ast.Expr, prefix string) (jsast.IExpression, error) {
	switch t := n.(type) {
	case *ast.Ident:
		if t.Name == "_" {
			return nil, nil
		}
		return jsast.CreateIdentifier(prefix + t.Name), nil
	case *ast.SelectorExpr:
		x, e := getExpression(t.X, prefix)
		if e != nil {
			return nil, e
		} else if x == nil {
			return nil, nil
		}

		return jsast.CreateMemberExpression(
			x,
			jsast.CreateIdentifier(t.Sel.Name),
			false,
		), nil
	default:
		return nil, unhandled("pattern", n)
	}
}

func getPatternName(n ast.Expr) (string, error) {
	switch t := n.(type) {
	case *ast.Ident:
		return t.Name, nil
	case *ast.SelectorExpr:
		return getPatternName(t.X)
	default:
		return "", unhandled("patternName", n)
	}
}

func getValue(n ast.Expr) (jsast.IExpression, error) {
	// ast.Print(nil, n)
	switch t := n.(type) {
	case *ast.BasicLit:
		return jsast.CreateLiteral(t.Value), nil
	case *ast.Ident:
		return jsast.CreateIdentifier(t.Name), nil
	case *ast.SelectorExpr:
		x, e := getValue(t.X)
		if e != nil {
			return nil, e
		}
		return jsast.CreateMemberExpression(
			x,
			jsast.CreateIdentifier(t.Sel.Name),
			false,
		), nil
	case *ast.UnaryExpr:
		return getValue(t.X)
	case nil:
		return jsast.CreateNull(), nil
	default:
		return nil, unhandled("getValue", n)
	}
}

func getArrayValue(n ast.Expr, prefix string, nth int) (jsast.IExpression, error) {
	// ast.Print(nil, n)
	switch t := n.(type) {
	case *ast.BasicLit:
		return jsast.CreateMemberExpression(
			jsast.CreateLiteral(prefix+t.Value),
			jsast.CreateInt(nth),
			true,
		), nil
	case *ast.Ident:
		return jsast.CreateMemberExpression(
			jsast.CreateIdentifier(prefix+t.Name),
			jsast.CreateInt(nth),
			true,
		), nil
	case *ast.SelectorExpr:
		return getArrayValue(t.X, prefix, nth)
	case nil:
		return jsast.CreateNull(), nil
	default:
		return nil, unhandled("getArrayValue", n)
	}
}

func getEmpty(n ast.Expr) (jsast.IExpression, error) {
	// ast.Print(nil, n)
	switch t := n.(type) {
	case *ast.Ident:
		return getEmptyIdent(t)
	case *ast.ArrayType:
		return jsast.CreateArrayExpression(), nil
	default:
		return nil, unhandled("getEmpty", n)
	}
}

func getEmptyIdent(n *ast.Ident) (jsast.IExpression, error) {
	switch n.Name {
	case "int", "int8", "int16", "int32", "int64",
		"uint8", "uint16", "uint32", "uint64":
		return jsast.CreateInt(0), nil
	case "string":
		return jsast.CreateString(""), nil
	case "float", "float32", "float64":
		return jsast.CreateFloat(0), nil
	default:
		return nil, fmt.Errorf("unhandled: getEmptyIdent '%s'", n.Name)
	}
}

func unhandled(fn string, n interface{}) error {
	return fmt.Errorf("%s in %s() is not implemented yet", reflect.TypeOf(n), fn)
}
