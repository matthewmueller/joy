package translator

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"

	"github.com/matthewmueller/golly/js"
)

// Translate from a Go AST to a Javascript AST
func Translate(fset *token.FileSet, f *ast.File) (p js.Program, err error) {
	ast.SortImports(fset, f)
	return fileToProgram(fset, f)
}

func fileToProgram(fset *token.FileSet, f *ast.File) (p js.Program, err error) {
	decls := f.Decls
	l := len(decls)

	// TODO: can i use additional interfaces to tighten this up?
	stmts := []interface{}{}
	for i := 0; i < l; i++ {
		stmt, e := declToStatement(fset, f, decls[i])
		if e != nil {
			return p, e
		}
		stmts = append(stmts, stmt)
	}

	// call main()
	callmain := js.CreateExpressionStatement(
		js.CreateCallExpression(
			js.CreateIdentifier("main"),
			[]js.IExpression{},
		),
	)

	// self-executing
	wrap := js.CreateExpressionStatement(
		js.CreateCallExpression(
			js.CreateFunctionExpression(nil, []js.IPattern{},
				js.CreateFunctionBody(append(stmts, callmain)...),
			),
			[]js.IExpression{},
		),
	)

	return js.CreateProgram(wrap), nil
}

func declToStatement(fset *token.FileSet, f *ast.File, decl ast.Decl) (s js.IStatement, err error) {
	switch d := decl.(type) {
	case *ast.FuncDecl:
		return function(fset, f, d)
	default:
		return js.CreateEmptyStatement(), nil
	}
}

func function(fset *token.FileSet, f *ast.File, fn *ast.FuncDecl) (j js.IStatement, err error) {
	// e.g. func hi()
	if fn.Body == nil {
		return js.CreateEmptyStatement(), nil
	}

	// anonymous fns
	if fn.Name == nil {
		return j, fmt.Errorf("function: anon functions not yet supported")
	}

	// create the params
	name := js.CreateIdentifier((*fn.Name).Name)

	// build argument list
	// var args
	var params []js.IPattern
	if fn.Type != nil && fn.Type.Params != nil {
		fields := fn.Type.Params.List
		for _, field := range fields {
			// names because: (a, b string, c int) = [[a, b], c]
			for _, name := range field.Names {
				id := js.CreateIdentifier(name.Name)
				params = append(params, id)
			}
		}
	}

	// create the body
	var body []interface{}
	for _, stmt := range fn.Body.List {
		jsStmt, e := funcStatement(fset, f, fn, stmt)
		if e != nil {
			return j, e
		}
		body = append(body, jsStmt)
	}

	return js.CreateFunctionDeclaration(
		&name,
		params,
		js.CreateFunctionBody(body...),
	), nil
}

// func mainFunc(fset *token.FileSet, f *ast.File, fn *ast.FuncDecl) (j js.IStatement, err error) {
// 	// e.g. func main()
// 	if fn.Body == nil {
// 		return js.CreateEmptyStatement(), nil
// 	}

// 	// TODO: []IStatement{} instead of interface{}
// 	var body []interface{}
// 	for _, stmt := range fn.Body.List {
// 		jsStmt, e := funcStatement(fset, f, fn, stmt)
// 		if e != nil {
// 			return j, e
// 		}
// 		body = append(body, jsStmt)
// 	}

// 	return js.CreateExpressionStatement(
// 		js.CreateCallExpression(
// 			js.CreateFunctionExpression(nil, []js.IPattern{},
// 				js.CreateFunctionBody(body...),
// 			),
// 			[]js.IExpression{},
// 		),
// 	), nil
// 	// return j, nil
// }

func funcStatement(fset *token.FileSet, f *ast.File, fn *ast.FuncDecl, stmt ast.Stmt) (j js.IStatement, err error) {
	switch t := stmt.(type) {
	case *ast.ExprStmt:
		return exprStatement(fset, f, fn, t)
	case *ast.IfStmt:
		return ifStmt(fset, f, fn, t)
	case *ast.AssignStmt:
		return assignStatement(fset, f, fn, t)
	case *ast.ReturnStmt:
		return returnStatement(fset, f, fn, t)
	default:
		return nil, unhandled("funcStatement", stmt)
	}
}

func exprStatement(fset *token.FileSet, f *ast.File, fn *ast.FuncDecl, expr *ast.ExprStmt) (j js.IExpressionStatement, err error) {
	switch t := expr.X.(type) {
	case *ast.CallExpr:
		x, e := callExpression(fset, f, fn, t)
		if e != nil {
			return nil, e
		}
		return js.CreateExpressionStatement(x), nil
	default:
		return nil, fmt.Errorf("exprStatement: unhandled type: %s", reflect.TypeOf(expr))
	}
}

func ifStmt(fset *token.FileSet, f *ast.File, fn *ast.FuncDecl, n *ast.IfStmt) (j js.IStatement, err error) {
	var e error

	// var consequent js.IBlockStatement
	// var alternate js.IBlockStatement

	// condition: if [(...)] { ... } else { ... }
	var test js.IExpression
	switch t := n.Cond.(type) {
	case *ast.BinaryExpr:
		test, e = binaryExpression(fset, f, fn, t)
	case *ast.Ident:
		test, e = identifier(fset, f, fn, t)
	default:
		return nil, unhandled("ifStmt<cond>", n.Cond)
	}
	if e != nil {
		return nil, e
	}

	// body : if (...) [{ ... }] else { ... }
	body, e := blockStmt(fset, f, fn, n.Body)
	if e != nil {
		return nil, e
	}

	// else: if (...) { ... } else [{ ... }]
	elseBlock := n.Else
	var alt js.IStatement
	switch t := elseBlock.(type) {
	case *ast.BlockStmt:
		alt, e = blockStmt(fset, f, fn, t)
	case *ast.ExprStmt:
		alt, e = exprStatement(fset, f, fn, t)
	case *ast.IfStmt:
		alt, e = ifStmt(fset, f, fn, t)
	case *ast.AssignStmt:
		alt, e = assignStatement(fset, f, fn, t)
	case *ast.ReturnStmt:
		alt, e = returnStatement(fset, f, fn, t)
	case nil:
		alt = js.CreateEmptyStatement()
	default:
		return nil, unhandled("ifStmt<else>", elseBlock)
	}
	if e != nil {
		return nil, e
	}

	return js.CreateIfStatement(
		test,
		body,
		alt,
	), nil
}

func blockStmt(fset *token.FileSet, f *ast.File, fn *ast.FuncDecl, n *ast.BlockStmt) (j js.IBlockStatement, err error) {
	var stmts []js.IStatement

	for _, stmt := range n.List {
		var v js.IStatement
		var e error
		switch t := stmt.(type) {
		case *ast.ExprStmt:
			v, e = exprStatement(fset, f, fn, t)
		case *ast.IfStmt:
			v, e = ifStmt(fset, f, fn, t)
		case *ast.AssignStmt:
			v, e = assignStatement(fset, f, fn, t)
		case *ast.ReturnStmt:
			v, e = returnStatement(fset, f, fn, t)
		default:
			return nil, unhandled("ifStmt<body>", stmt)
		}
		if e != nil {
			return nil, e
		}
		stmts = append(stmts, v)
	}

	return js.CreateBlockStatement(stmts...), nil
}

func assignStatement(fset *token.FileSet, f *ast.File, fn *ast.FuncDecl, as *ast.AssignStmt) (j js.IStatement, err error) {
	lhs := as.Lhs
	rhs := as.Rhs

	// handle balanced destructuring
	var decls []js.VariableDeclarator
	if len(lhs) == len(rhs) {
		for i, expr := range lhs {
			var id js.IPattern

			switch t := expr.(type) {
			case *ast.Ident:
				id = js.CreateIdentifier(t.Name)
			default:
				return j, fmt.Errorf("assignStatement: unhandled type: %s", reflect.TypeOf(expr))
			}

			rh, e := expression(fset, f, fn, rhs[i])
			if e != nil {
				return j, e
			}

			decl := js.CreateVariableDeclarator(id, rh)
			decls = append(decls, decl)
		}

		return js.CreateVariableDeclaration("var", decls...), nil
	}

	return j, fmt.Errorf("assignStatement: unbalanced variables not implemented yet")
}

func returnStatement(fset *token.FileSet, f *ast.File, fn *ast.FuncDecl, n *ast.ReturnStmt) (j js.IStatement, err error) {
	// no return values
	if len(n.Results) == 0 {
		return js.CreateReturnStatement(nil), nil
	}

	if len(n.Results) > 1 {
		return nil, fmt.Errorf("returnStatement: multiple return values not implement yet")
	}

	var args []js.IExpression
	for _, arg := range n.Results {
		a, e := expression(fset, f, fn, arg)
		if e != nil {
			return nil, e
		}
		args = append(args, a)
	}

	// return an array (TODO: be smarter later)
	return js.CreateReturnStatement(args[0]), nil
}

func callExpression(fset *token.FileSet, f *ast.File, fn *ast.FuncDecl, expr *ast.CallExpr) (j js.CallExpression, err error) {
	callee, e := expression(fset, f, fn, expr.Fun)
	if e != nil {
		return j, e
	}

	var args []js.IExpression
	for _, arg := range expr.Args {
		var v js.IExpression
		var e error
		switch t := arg.(type) {
		case *ast.BasicLit:
			v, e = basiclit(fset, f, fn, t)
		case *ast.Ident:
			v, e = identifier(fset, f, fn, t)
		case *ast.SelectorExpr:
			v, e = selectorExpr(fset, f, fn, t)
		default:
			return j, unhandled("callExpression", arg)
		}
		if e != nil {
			return j, e
		}
		args = append(args, v)
	}

	return js.CreateCallExpression(
		callee,
		args,
	), nil
}

// TODO: move into the respective functions, the translation is not 1:1
func expression(fset *token.FileSet, f *ast.File, fn *ast.FuncDecl, expr ast.Expr) (j js.IExpression, err error) {
	switch t := expr.(type) {
	case *ast.Ident:
		return identifier(fset, f, fn, t)
	case *ast.BasicLit:
		return basiclit(fset, f, fn, t)
	case *ast.CallExpr:
		return callExpression(fset, f, fn, t)
	case *ast.BinaryExpr:
		return binaryExpression(fset, f, fn, t)
	case *ast.CompositeLit:
		return compositeLiteral(fset, f, fn, t)
	default:
		return nil, fmt.Errorf("expression(): unhandled type: %s", reflect.TypeOf(expr))
	}
}

// binary expressions in Go can be either:
//    Binaryexpression || LogicalExpression
// in JS.
func binaryExpression(fset *token.FileSet, f *ast.File, fn *ast.FuncDecl, b *ast.BinaryExpr) (j js.IExpression, err error) {
	x, e := expression(fset, f, fn, b.X)
	if e != nil {
		return nil, e
	}
	y, e := expression(fset, f, fn, b.Y)
	if e != nil {
		return nil, e
	}

	if !b.Op.IsOperator() {
		return nil, unhandled("binaryExpression<not op>", b.Op)
	}

	op := b.Op.String()
	switch op {
	case "||", "&&":
		return js.CreateLogicalExpression(x, js.LogicalOperator(op), y), nil
	// TODO: prune. should be only values that are possible in Go
	case "==", "!=", "===", "!==",
		"<", "<=", ">", ">=", "<<",
		">>", ">>>", "+", "-", "*",
		"/", "%", "|", "^", "&",
		"in", "instanceof":
		return js.CreateBinaryExpression(x, js.BinaryOperator(op), y), nil
	default:
		return nil, unhandled("binaryExpression<unknown op>", op)
	}
}

func identifier(fset *token.FileSet, f *ast.File, fn *ast.FuncDecl, ident *ast.Ident) (j js.IExpression, err error) {
	// TODO: lookup table
	switch ident.Name {
	case "println":
		return js.CreateMemberExpression(
			js.CreateIdentifier("console"),
			js.CreateIdentifier("log"),
			false,
		), nil
	default:
		return js.CreateIdentifier(ident.Name), nil
	}
}

func basiclit(fset *token.FileSet, f *ast.File, fn *ast.FuncDecl, lit *ast.BasicLit) (j js.IExpression, err error) {
	return js.CreateString(lit.Value), nil
}

func compositeLiteral(fset *token.FileSet, f *ast.File, fn *ast.FuncDecl, n *ast.CompositeLit) (j js.ObjectExpression, err error) {
	var props []js.Property

	for _, elt := range n.Elts {
		var prop js.Property
		switch t := elt.(type) {
		case *ast.KeyValueExpr:
			p, e := keyValueExpr(fset, f, fn, t)
			if e != nil {
				return j, e
			}
			prop = p
		case *ast.CompositeLit:
			// k, e := t.Type
			// t.Type
			// t.Elts[]
			// v, e := compositeLiteral(fset, f, fn, t)
			// if e != nil {
			// 	return j, e
			// }

			// prop = v
		default:
			return j, fmt.Errorf("compositeLiteral(): not sure what this type is %s", reflect.TypeOf(elt))
		}

		props = append(props, prop)
	}

	return js.CreateObjectExpression(props), nil
}

func keyValueExpr(fset *token.FileSet, f *ast.File, fn *ast.FuncDecl, n *ast.KeyValueExpr) (j js.Property, err error) {
	var key interface{}
	var value js.IExpression

	// get the key
	switch t := n.Key.(type) {
	case *ast.Ident:
		k, e := identifier(fset, f, fn, t)
		if e != nil {
			return j, e
		}
		key = k
	default:
		return j, unhandled("keyValueExpr<key>", n.Key)
	}

	// get the value
	switch t := n.Value.(type) {
	case *ast.BasicLit:
		v, e := basiclit(fset, f, fn, t)
		if e != nil {
			return j, e
		}
		value = v
	case *ast.CompositeLit:
		v, e := compositeLiteral(fset, f, fn, t)
		if e != nil {
			return j, e
		}
		value = v
	case *ast.Ident:
		v, e := identifier(fset, f, fn, t)
		if e != nil {
			return j, e
		}
		value = v
	default:
		return j, unhandled("keyValueExpr<value>", n.Value)
	}

	return js.CreateProperty(key, value, "init"), nil
}

func selectorExpr(fset *token.FileSet, f *ast.File, fn *ast.FuncDecl, n *ast.SelectorExpr) (j js.IExpression, err error) {
	var lhs js.IExpression
	switch t := n.X.(type) {
	// (user.phone).number
	case *ast.SelectorExpr:
		x, e := selectorExpr(fset, f, fn, t)
		if e != nil {
			return nil, e
		}
		lhs = x
	// (user).phone
	case *ast.Ident:
		x, e := identifier(fset, f, fn, t)
		if e != nil {
			return nil, e
		}
		lhs = x
	default:
		return nil, unhandled("selectorExpr", n.X)
	}

	// user.phone.(number)
	s, e := identifier(fset, f, fn, n.Sel)
	if e != nil {
		return nil, e
	}

	return js.CreateMemberExpression(lhs, s, false), nil
}

func unhandled(fn string, n interface{}) error {
	return fmt.Errorf("%s(): not sure what this type is %s", fn, reflect.TypeOf(n))
}
