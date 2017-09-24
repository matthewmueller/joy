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
		stmt, e := decl(fset, f, decls[i])
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

func decl(fset *token.FileSet, f *ast.File, decl ast.Decl) (s js.IStatement, err error) {
	switch d := decl.(type) {
	case *ast.FuncDecl:
		return function(fset, f, d)
	case *ast.GenDecl:
		return genDecl(fset, f, d)
	default:
		return nil, unhandled("decl", decl)
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

	// the name of the fn
	fnname := js.CreateIdentifier((*fn.Name).Name)

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

	// no receiver means it's a classless function
	if fn.Recv == nil {
		return js.CreateFunctionDeclaration(
			&fnname,
			params,
			js.CreateFunctionBody(body...),
		), nil
	}

	if len(fn.Recv.List) != 1 {
		return nil, fmt.Errorf("function<recv>: only expected 1 func receiver but got %d", len(fn.Recv.List))
	}

	recv := fn.Recv.List[0]
	x, e := expression(fset, f, fn, recv.Type)
	if e != nil {
		return nil, e
	}

	// {recv}.prototype.{name} = function ({params}) {
	//   {body}
	// }
	return js.CreateExpressionStatement(
		js.CreateAssignmentExpression(
			js.CreateMemberExpression(
				js.CreateMemberExpression(
					x,
					js.CreateIdentifier("prototype"),
					false,
				),
				fnname,
				false,
			),
			js.AssignmentOperator("="),
			js.CreateFunctionExpression(
				nil,
				params,
				js.CreateFunctionBody(body...),
			),
		),
	), nil

}

func genDecl(fset *token.FileSet, f *ast.File, n *ast.GenDecl) (j js.IStatement, err error) {
	switch n.Tok.String() {
	case "type":
		return typeSpec(fset, f, n)
	case "import":
		log.Infof("ignoring import statement")
		return js.CreateEmptyStatement(), nil
	case "var":
		return varSpec(fset, f, n)
	default:
		return nil, fmt.Errorf("genDecl: unhandled token %s", n.Tok.String())
	}

	// // specs := n.Specs
	// for _, spec := range n.Specs {
	// 	switch t := spec.(type) {
	// 	// case *ast.ImportSpec:
	// 	// 	return importSpec(fset, f, t)
	// 	case *ast.TypeSpec:
	// 		// type defs will only have 1 spec
	// 		return typeSpec(fset, f, t)
	// 	default:
	// 		return nil, unhandled("genDecl", spec)
	// 	}
	// }

	// return js.CreateEmptyStatement(), nil
}

// func importSpec(fset *token.FileSet, f *ast.File, n *ast.ImportSpec) (j js.IStatement, err error) {
// 	return nil, nil
// }

func typeSpec(fset *token.FileSet, f *ast.File, n *ast.GenDecl) (j js.IStatement, err error) {
	if len(n.Specs) != 1 {
		return nil, fmt.Errorf("genDecl: expected type to only have 1 spec but it has %d", len(n.Specs))
	}

	s, ok := n.Specs[0].(*ast.TypeSpec)
	if !ok {
		return nil, unhandled("typespec", n.Specs[0])
	}

	switch t := s.Type.(type) {
	case *ast.StructType:
		// we turn these into contructor functions
		// so more functions can get attached to them
		return jsConstructorFunction(fset, f, s.Name, t.Fields)
	default:
		return nil, unhandled("typeSpec", s.Type)
	}
}

func varSpec(fset *token.FileSet, f *ast.File, n *ast.GenDecl) (j js.IStatement, err error) {
	var decls []js.VariableDeclarator

	for _, spec := range n.Specs {
		var id js.IPattern
		var e error
		_, _ = id, e

		switch t := spec.(type) {
		case *ast.ValueSpec:
			if len(t.Names) != len(t.Values) {
				return nil, unhandled("varSpec<unbalanced>", spec)
			}

			// handle balanced destructuring
			for i, ident := range t.Names {
				lh := js.CreateIdentifier(ident.Name)

				rh, e := expression(fset, f, nil, t.Values[i])
				if e != nil {
					return j, e
				}

				decl := js.CreateVariableDeclarator(lh, rh)
				decls = append(decls, decl)
			}

		default:
			return nil, unhandled("varSpec", spec)
		}
	}

	return js.CreateVariableDeclaration("var", decls...), nil
	// return nil, nil
}

func jsConstructorFunction(fset *token.FileSet, f *ast.File, id *ast.Ident, fields *ast.FieldList) (j js.FunctionDeclaration, err error) {
	name := js.CreateIdentifier(id.Name)

	// instance variables
	// TODO: this should actually be: []js.IExpressionStatement
	// but CreateFunctionBody accepts an []interface{} currently
	var ivars []interface{}

	for _, field := range fields.List {
		for _, name := range field.Names {
			// this.$name = o.$name
			ivars = append(ivars, js.CreateExpressionStatement(
				js.CreateAssignmentExpression(
					js.CreateMemberExpression(
						js.CreateThisExpression(),
						js.CreateIdentifier(name.Name),
						false,
					),
					js.AssignmentOperator("="),
					js.CreateMemberExpression(
						js.CreateIdentifier("o"),
						js.CreateIdentifier(name.Name),
						false,
					),
				),
			))
		}
	}

	return js.CreateFunctionDeclaration(
		&name,
		// TODO: make API for this
		[]js.IPattern{js.CreateIdentifier("o")},
		js.CreateFunctionBody(ivars...),
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
	case *ast.DeclStmt:
		return declStmt(fset, f, fn, t)
	default:
		return nil, unhandled("funcStatement", stmt)
	}
}

func declStmt(fset *token.FileSet, f *ast.File, fn *ast.FuncDecl, n *ast.DeclStmt) (j js.IStatement, err error) {
	switch t := n.Decl.(type) {
	case *ast.GenDecl:
		return genDecl(fset, f, t)
	default:
		return nil, unhandled("declStmt", n)
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
		for i, lh := range lhs {
			var id js.IPattern
			var e error

			switch t := lh.(type) {
			case *ast.Ident:
				id = js.CreateIdentifier(t.Name)
			case *ast.SelectorExpr:
				id, e = selectorExpr(fset, f, fn, t)
			default:
				return j, fmt.Errorf("assignStatement: unhandled type: %s", reflect.TypeOf(lh))
			}
			if e != nil {
				return j, e
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

	// a, err := test()
	// => var $a = test(), a = $a[0], err = $a[1]
	// TODO: clean up this is terrible
	if len(lhs) > len(rhs) && len(lhs) > 0 && len(rhs) > 0 {

		t, ok := lhs[0].(*ast.Ident)
		if !ok {
			return j, fmt.Errorf("assignStatement<lhs[0]>: unhandled type: %s", reflect.TypeOf(lhs[0]))
		}
		lh := js.CreateIdentifier("$" + t.Name)

		rh, e := expression(fset, f, fn, rhs[0])
		if e != nil {
			return j, e
		}

		decl := js.CreateVariableDeclarator(lh, rh)
		decls = append(decls, decl)

		for i := 0; i < len(lhs); i++ {
			t, ok := lhs[i].(*ast.Ident)
			if !ok {
				return j, fmt.Errorf("assignStatement<lhs[i]>: unhandled type: %s", reflect.TypeOf(lhs[i]))
			}

			decl := js.CreateVariableDeclarator(
				js.CreateIdentifier(t.Name),
				js.CreateMemberExpression(
					lh,
					js.CreateInt(i),
					true,
				),
			)
			decls = append(decls, decl)
		}

		return js.CreateVariableDeclaration("var", decls...), nil
	}

	return j, fmt.Errorf("assignStatement: more values on right than left, is this possible?")
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

	var args []js.IExpression
	for _, arg := range n.Results {
		a, e := expression(fset, f, fn, arg)
		if e != nil {
			return nil, e
		}
		args = append(args, a)
	}

	// return an array
	if len(n.Results) > 1 {
		return js.CreateReturnStatement(js.CreateArrayExpression(args...)), nil
	}

	// return the value by itself
	return js.CreateReturnStatement(args[0]), nil
}

func callExpression(fset *token.FileSet, f *ast.File, fn *ast.FuncDecl, expr *ast.CallExpr) (j js.IExpression, err error) {
	calleeSrc, e := expressionToString(expr.Fun)
	if e != nil {
		return j, e
	}

	// TODO: make sure js.Raw points to golly/js
	if calleeSrc == "js.Raw" && len(expr.Args) >= 1 {
		argSrc, e := expressionToString(expr.Args[0])
		if e != nil {
			return nil, e
		}
		return js.Parse(argSrc)
	}

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
	case *ast.StarExpr:
		return starExpr(fset, f, fn, t)
	case *ast.UnaryExpr:
		return unaryExpr(fset, f, fn, t)
	case *ast.FuncLit:
		return funcLit(fset, f, fn, t)
	// case *ast.KeyValueExpr:
	// 	return keyValueExpr(fset, f, fn, t)
	case nil:
		return nil, nil
	default:
		return nil, fmt.Errorf("expression(): unhandled type: %s", reflect.TypeOf(expr))
	}
}

func expressionToString(expr ast.Expr) (string, error) {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name, nil
	case *ast.SelectorExpr:
		x, e := expressionToString(t.X)
		if e != nil {
			return "", e
		}
		return x + "." + t.Sel.Name, nil
	case *ast.BasicLit:
		return t.Value, nil
	default:
		return "", nil
	}
}

func funcLit(fset *token.FileSet, f *ast.File, fn *ast.FuncDecl, n *ast.FuncLit) (j js.IExpression, err error) {
	// build argument list
	// var args
	var params []js.IPattern
	if n.Type != nil && n.Type.Params != nil {
		fields := n.Type.Params.List
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
	for _, stmt := range n.Body.List {
		jsStmt, e := funcStatement(fset, f, fn, stmt)
		if e != nil {
			return j, e
		}
		body = append(body, jsStmt)
	}

	return js.CreateFunctionExpression(nil, params, js.CreateFunctionBody(body...)), nil
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
	case "nil":
		return js.CreateNull(), nil
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

func starExpr(fset *token.FileSet, f *ast.File, fn *ast.FuncDecl, n *ast.StarExpr) (j js.IExpression, err error) {
	// for now, we're ignoring the pointer
	x, e := expression(fset, f, fn, n.X)
	if e != nil {
		return nil, e
	}

	return x, nil
}

func unaryExpr(fset *token.FileSet, f *ast.File, fn *ast.FuncDecl, n *ast.UnaryExpr) (j js.IExpression, err error) {
	// for now, we're ignoring the pointer
	x, e := expression(fset, f, fn, n.X)
	if e != nil {
		return nil, e
	}

	return x, nil
}

func basiclit(fset *token.FileSet, f *ast.File, fn *ast.FuncDecl, lit *ast.BasicLit) (j js.IExpression, err error) {
	return js.CreateString(lit.Value), nil
}

func compositeLiteral(fset *token.FileSet, f *ast.File, fn *ast.FuncDecl, n *ast.CompositeLit) (j js.IExpression, err error) {
	switch n.Type.(type) {
	case *ast.Ident:
		return jsNewFunction(fset, f, fn, n)
	case *ast.ArrayType:
		return jsArrayExpression(fset, f, fn, n)
	case *ast.MapType:
		return jsObjectExpression(fset, f, fn, n)
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

func jsNewFunction(fset *token.FileSet, f *ast.File, fn *ast.FuncDecl, n *ast.CompositeLit) (j js.NewExpression, err error) {
	var props []js.Property
	var obj *ast.Object
	_ = obj
	var fnname string

	switch t := n.Type.(type) {
	case *ast.Ident:
		fnname = t.Name
		obj = t.Obj
	default:
		return j, unhandled("jsNewFunction<name>", n.Type)
	}

	for _, elt := range n.Elts {
		var prop js.Property
		var e error

		switch t := elt.(type) {
		case *ast.KeyValueExpr:
			prop, e = keyValueExpr(fset, f, fn, t)
		// case *ast.BasicLit:
		// TODO: handle User{"a"}, by looking up the fields in obj

		default:
			return j, unhandled("jsNewFunction<elts>", elt)
		}
		if e != nil {
			return j, e
		}
		props = append(props, prop)
	}

	return js.CreateNewExpression(
		js.CreateIdentifier(fnname),
		[]js.IExpression{js.CreateObjectExpression(props)},
	), nil
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

func selectorExpr(fset *token.FileSet, f *ast.File, fn *ast.FuncDecl, n *ast.SelectorExpr) (j js.MemberExpression, err error) {
	// (user.phone).number
	x, e := expression(fset, f, fn, n.X)
	if e != nil {
		return j, e
	}

	// user.phone.(number)
	s, e := identifier(fset, f, fn, n.Sel)
	if e != nil {
		return j, e
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
	return fmt.Errorf("%s in %s() is not implemented yet", reflect.TypeOf(n), fn)
}
