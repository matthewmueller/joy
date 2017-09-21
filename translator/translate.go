package translator

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"

	"github.com/apex/log"
	"github.com/matthewmueller/golly/js"
	"github.com/pkg/errors"
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
	case *ast.ForStmt:
		return forStmt(fset, f, fn, t)
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

func forStmt(fset *token.FileSet, f *ast.File, fn *ast.FuncDecl, n *ast.ForStmt) (j js.IStatement, err error) {

	init, e := statement(fset, f, fn, n.Init)
	if e != nil {
		return nil, errors.Wrap(e, "forStmt")
	}

	cond, e := expression(fset, f, fn, n.Cond)
	if e != nil {
		return nil, errors.Wrap(e, "forStmt")
	}

	post, e := statement(fset, f, fn, n.Post)
	if e != nil {
		return nil, errors.Wrap(e, "forStmt")
	}

	body, e := blockStmt(fset, f, fn, n.Body)
	if e != nil {
		return nil, errors.Wrap(e, "forStmt")
	}

	// In Go the post condition is a statement,
	// in JS it's an expression
	//
	// it can also be nil in the case of for { ... }
	var postexpr js.IExpression
	switch t := post.(type) {
	case js.ExpressionStatement:
		postexpr = t.Expression
	case nil:
		postexpr = nil
	default:
		return nil, unhandled("forStmt<post>", post)
	}

	return js.CreateForStatement(
		init,
		cond,
		postexpr,
		body,
	), nil
}

func statement(fset *token.FileSet, f *ast.File, fn *ast.FuncDecl, n ast.Stmt) (j js.IStatement, err error) {
	switch t := n.(type) {
	case nil:
		return nil, nil
	case *ast.AssignStmt:
		return assignStatement(fset, f, fn, t)
	case *ast.IncDecStmt:
		return incDecStmt(fset, f, fn, t)
	case *ast.ExprStmt:
		return exprStatement(fset, f, fn, t)
	default:
		return nil, unhandled("statement", n)
	}
}

func blockStmt(fset *token.FileSet, f *ast.File, fn *ast.FuncDecl, n *ast.BlockStmt) (j js.IBlockStatement, err error) {
	var stmts []js.IStatement

	for _, stmt := range n.List {
		v, e := statement(fset, f, fn, stmt)
		if e != nil {
			return nil, errors.Wrap(e, "blockStmt")
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

func incDecStmt(fset *token.FileSet, f *ast.File, fn *ast.FuncDecl, n *ast.IncDecStmt) (j js.IStatement, err error) {
	var op js.UpdateOperator

	x, e := expression(fset, f, fn, n.X)
	if e != nil {
		return nil, errors.Wrap(e, "incDecStmt")
	}

	log.Infof("n.Tok.String(): %s", n.Tok.String())
	switch n.Tok.String() {
	case "++":
		op = js.UpdateOperator("++")
	case "--":
		op = js.UpdateOperator("--")
	default:
		return nil, unhandled("incDecStmt", n.Tok)
	}

	return js.CreateExpressionStatement(
		js.CreateUpdateExpression(x, op, false),
	), nil
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
		v, e := expression(fset, f, fn, arg)
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
	case *ast.SelectorExpr:
		return selectorExpr(fset, f, fn, t)
	case *ast.IndexExpr:
		return indexExpr(fset, f, fn, t)
	// case *ast.KeyValueExpr:
	// 	return keyValueExpr(fset, f, fn, t)
	case nil:
		return nil, nil
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

func compositeLiteral(fset *token.FileSet, f *ast.File, fn *ast.FuncDecl, n *ast.CompositeLit) (j js.IExpression, err error) {
	switch n.Type.(type) {
	case *ast.Ident:
		return jsObjectExpression(fset, f, fn, n)
	case *ast.ArrayType:
		return jsArrayExpression(fset, f, fn, n)
	default:
		return nil, unhandled("compositeLiteral<type>", n.Type)
	}
}

func jsObjectExpression(fset *token.FileSet, f *ast.File, fn *ast.FuncDecl, n *ast.CompositeLit) (j js.ObjectExpression, err error) {
	var props []js.Property

	for _, elt := range n.Elts {
		var prop js.Property
		var e error

		switch t := elt.(type) {
		case *ast.KeyValueExpr:
			prop, e = keyValueExpr(fset, f, fn, t)
		default:
			return j, unhandled("jsObjectExpression", elt)
		}
		if e != nil {
			return j, e
		}
		props = append(props, prop)
	}

	return js.CreateObjectExpression(props), nil
}

func jsArrayExpression(fset *token.FileSet, f *ast.File, fn *ast.FuncDecl, n *ast.CompositeLit) (j js.ArrayExpression, err error) {
	var elements []js.IExpression

	for _, elt := range n.Elts {
		el, e := expression(fset, f, fn, elt)
		if e != nil {
			return j, e
		}
		elements = append(elements, el)
	}

	return js.CreateArrayExpression(elements...), nil
}

func keyValueExpr(fset *token.FileSet, f *ast.File, fn *ast.FuncDecl, n *ast.KeyValueExpr) (j js.Property, err error) {
	// get the key
	key, e := expression(fset, f, fn, n.Key)
	if e != nil {
		return j, e
	}

	// get the value
	value, e := expression(fset, f, fn, n.Value)
	if e != nil {
		return j, e
	}

	return js.CreateProperty(key, value, "init"), nil
}

func selectorExpr(fset *token.FileSet, f *ast.File, fn *ast.FuncDecl, n *ast.SelectorExpr) (j js.IExpression, err error) {
	// (user.phone).number
	x, e := expression(fset, f, fn, n.X)
	if e != nil {
		return nil, e
	}

	// user.phone.(number)
	s, e := identifier(fset, f, fn, n.Sel)
	if e != nil {
		return nil, e
	}

	return js.CreateMemberExpression(x, s, false), nil
}

func indexExpr(fset *token.FileSet, f *ast.File, fn *ast.FuncDecl, n *ast.IndexExpr) (j js.MemberExpression, err error) {
	// (i)[0]
	x, e := expression(fset, f, fn, n.X)
	if e != nil {
		return j, e
	}

	// i([0])
	i, e := expression(fset, f, fn, n.Index)
	if e != nil {
		return j, e
	}

	return js.CreateMemberExpression(x, i, true), nil
}

func unhandled(fn string, n interface{}) error {
	return fmt.Errorf("%s(): not sure what this type is %s", fn, reflect.TypeOf(n))
}
