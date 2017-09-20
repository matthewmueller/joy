package translator

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"

	"github.com/apex/log"
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

	return js.CreateProgram(stmts...), nil
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

	if fn.Name != nil && fn.Name.Name == "main" {
		log.Infof("handling main differently")
		return mainFunc(fset, f, fn)
	}

	return j, nil
}

func mainFunc(fset *token.FileSet, f *ast.File, fn *ast.FuncDecl) (j js.IStatement, err error) {
	// e.g. func main()
	if fn.Body == nil {
		return js.CreateEmptyStatement(), nil
	}

	// TODO: []IStatement{} instead of interface{}
	var body []interface{}
	for _, stmt := range fn.Body.List {
		jsStmt, e := funcStatement(fset, f, fn, stmt)
		if e != nil {
			return j, e
		}
		body = append(body, jsStmt)
	}

	return js.CreateExpressionStatement(
		js.CreateCallExpression(
			js.CreateFunctionExpression(nil, []js.IPattern{},
				js.CreateFunctionBody(body...),
			),
			[]js.IExpression{},
		),
	), nil
	// return j, nil
}

func funcStatement(fset *token.FileSet, f *ast.File, fn *ast.FuncDecl, stmt ast.Stmt) (j js.IStatement, err error) {
	switch s := stmt.(type) {
	case *ast.ExprStmt:
		return exprStatement(fset, f, fn, s)
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
	default:
		return nil, fmt.Errorf("expression: unhandled type: %s", reflect.TypeOf(expr))
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
