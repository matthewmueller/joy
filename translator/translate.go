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
	// if fn.Recv != nil {
	// 	fn.Recv.List
	// }
	// for _, field := range fn.Recv {

	// }

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
	switch s := stmt.(type) {
	case *ast.ExprStmt:
		return exprStatement(fset, f, fn, s)
	case *ast.AssignStmt:
		return assignStatement(fset, f, fn, s)
	case *ast.ReturnStmt:
		return returnStatement(fset, f, fn, s)
	default:
		return nil, fmt.Errorf("funcStatement: unhandled type %s", reflect.TypeOf(stmt))
	}
}

func exprStatement(fset *token.FileSet, f *ast.File, fn *ast.FuncDecl, expr *ast.ExprStmt) (j js.IStatement, err error) {
	switch s := expr.X.(type) {
	case *ast.CallExpr:
		x, e := callExpression(fset, f, fn, s)
		if e != nil {
			return nil, e
		}
		return js.CreateExpressionStatement(x), nil
	default:
		return nil, fmt.Errorf("exprStatement: unhandled type: %s", reflect.TypeOf(expr))
	}
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

func callExpression(fset *token.FileSet, f *ast.File, fn *ast.FuncDecl, expr *ast.CallExpr) (j js.IExpression, err error) {
	callee, e := expression(fset, f, fn, expr.Fun)
	if e != nil {
		return nil, e
	}

	var args []js.IExpression
	for _, arg := range expr.Args {
		a, e := expression(fset, f, fn, arg)
		if e != nil {
			return nil, e
		}
		args = append(args, a)
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
	default:
		return nil, fmt.Errorf("expression: unhandled type: %s", reflect.TypeOf(expr))
	}
}

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
		return nil, fmt.Errorf("binaryExpression: not sure how to handle this operator: %s", b.Op.String())
	}

	op := js.BinaryOperator(b.Op.String())

	return js.CreateBinaryExpression(x, op, y), nil
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
