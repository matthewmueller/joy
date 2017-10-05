package variable

import (
	"fmt"
	"go/ast"
	"reflect"

	js "github.com/matthewmueller/golly/jsast"
	"github.com/pkg/errors"
)

// Handle function
//
// Unfortunately, this process has lots of interweaving parts to it.
// That being said, this code is pretty crappy and there's lots that
// could be improved.
//
// TODO: move me
func Handle(n interface{}) (js.IStatement, js.IExpression, error) {
	switch t := n.(type) {
	case *ast.GenDecl:
		return genDeclPairs(t)
	case *ast.AssignStmt:
		return assignStmtPairs(t)
	default:
		return nil, nil, unhandled("variablePairs", t)
	}
}

func genDeclPairs(n *ast.GenDecl) (js.IStatement, js.IExpression, error) {
	var decls []js.VariableDeclarator

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

	return js.CreateVariableDeclaration("var", decls...), nil, nil
}

func valueSpecHandler(n *ast.ValueSpec) (decls []js.VariableDeclarator, e error) {
	lv := len(n.Values)

	for i, name := range n.Names {
		var val js.IExpression

		if i < lv {
			val, e = getValue(n.Values[i])
		} else {
			val, e = getEmpty(n.Type)
		}
		if e != nil {
			return decls, e
		}

		decls = append(decls, js.CreateVariableDeclarator(
			js.CreateIdentifier(name.Name),
			val,
		))
	}

	return decls, nil
}

func assignStmtPairs(n *ast.AssignStmt) (js.IStatement, js.IExpression, error) {
	var decls []js.VariableDeclarator
	var exprs []js.IExpression
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
		return nil, js.CreateSequenceExpression(exprs...), nil
	}

	for i, l := range n.Lhs {
		var r ast.Expr
		if i < lrh {
			r = n.Rhs[i]
		}

		var d *js.VariableDeclarator
		var x js.IExpression
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
		return js.CreateVariableDeclaration("var", decls...),
			js.CreateSequenceExpression(exprs...),
			nil
	}

	if len(decls) > 0 {
		return js.CreateVariableDeclaration("var", decls...),
			nil,
			nil
	}

	if len(exprs) > 0 {
		return nil,
			js.CreateSequenceExpression(exprs...),
			nil
	}

	return nil, nil, nil
}

func assignStmtVars(l ast.Expr, r ast.Expr, prefixLeft, prefixRight string, nth int) (*js.VariableDeclarator, error) {
	jl, e := getPattern(l, prefixLeft)
	if e != nil {
		return nil, e
	} else if jl == nil {
		return nil, nil
	}

	var jr js.IExpression
	if prefixRight == "" {
		jr, e = getValue(r)
	} else {
		jr, e = getArrayValue(r, prefixRight, nth)
	}
	if e != nil {
		return nil, e
	}

	decl := js.CreateVariableDeclarator(jl, jr)
	return &decl, nil
}

func assignStmtExprs(l ast.Expr, r ast.Expr, prefixLeft, prefixRight string, nth int) (js.IExpression, error) {
	jl, e := getExpression(l, prefixLeft)
	if e != nil {
		return nil, e
	} else if jl == nil {
		return nil, nil
	}

	var jr js.IExpression
	if prefixRight == "" {
		jr, e = getValue(r)
	} else {
		jr, e = getArrayValue(r, prefixRight, nth)
	}
	if e != nil {
		return nil, e
	}

	return js.CreateAssignmentExpression(
		jl,
		js.AssignmentOperator("="),
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

func getPattern(n ast.Expr, prefix string) (js.IPattern, error) {
	switch t := n.(type) {
	case *ast.Ident:
		if t.Name == "_" {
			return nil, nil
		}
		return js.CreateIdentifier(prefix + t.Name), nil
	case *ast.SelectorExpr:
		return getPattern(t.X, prefix)
	default:
		return nil, unhandled("getPattern()", n)
	}
}

func getExpression(n ast.Expr, prefix string) (js.IExpression, error) {
	switch t := n.(type) {
	case *ast.Ident:
		if t.Name == "_" {
			return nil, nil
		}
		return js.CreateIdentifier(prefix + t.Name), nil
	case *ast.SelectorExpr:
		x, e := getExpression(t.X, prefix)
		if e != nil {
			return nil, e
		} else if x == nil {
			return nil, nil
		}

		return js.CreateMemberExpression(
			x,
			js.CreateIdentifier(t.Sel.Name),
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

func getValue(n ast.Expr) (js.IExpression, error) {
	// ast.Print(nil, n)
	switch t := n.(type) {
	case *ast.BasicLit:
		return js.CreateLiteral(t.Value), nil
	case *ast.Ident:
		return js.CreateIdentifier(t.Name), nil
	case *ast.SelectorExpr:
		x, e := getValue(t.X)
		if e != nil {
			return nil, e
		}
		return js.CreateMemberExpression(
			x,
			js.CreateIdentifier(t.Sel.Name),
			false,
		), nil
	case *ast.UnaryExpr:
		return getValue(t.X)
	case nil:
		return js.CreateNull(), nil
	default:
		return nil, unhandled("getValue", n)
	}
}

func getArrayValue(n ast.Expr, prefix string, nth int) (js.IExpression, error) {
	// ast.Print(nil, n)
	switch t := n.(type) {
	case *ast.BasicLit:
		return js.CreateMemberExpression(
			js.CreateLiteral(prefix+t.Value),
			js.CreateInt(nth),
			true,
		), nil
	case *ast.Ident:
		return js.CreateMemberExpression(
			js.CreateIdentifier(prefix+t.Name),
			js.CreateInt(nth),
			true,
		), nil
	case *ast.SelectorExpr:
		return getArrayValue(t.X, prefix, nth)
	case nil:
		return js.CreateNull(), nil
	default:
		return nil, unhandled("getArrayValue", n)
	}
}

func getEmpty(n ast.Expr) (js.IExpression, error) {
	// ast.Print(nil, n)
	switch t := n.(type) {
	case *ast.Ident:
		return getEmptyIdent(t)
	case *ast.ArrayType:
		return js.CreateArrayExpression(), nil
	default:
		return nil, unhandled("getEmpty", n)
	}
}

func getEmptyIdent(n *ast.Ident) (js.IExpression, error) {
	switch n.Name {
	case "int", "int8", "int16", "int32", "int64",
		"uint8", "uint16", "uint32", "uint64":
		return js.CreateInt(0), nil
	case "string":
		return js.CreateString(""), nil
	case "float", "float32", "float64":
		return js.CreateFloat(0), nil
	default:
		return nil, fmt.Errorf("unhandled: getEmptyIdent '%s'", n.Name)
	}
}

func unhandled(fn string, n interface{}) error {
	return fmt.Errorf("%s in %s() is not implemented yet", reflect.TypeOf(n), fn)
}
