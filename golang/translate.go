package golang

import (
	"fmt"
	"go/ast"
	"path"
	"reflect"
	"strings"
	"unicode"

	"golang.org/x/tools/go/loader"

	"github.com/apex/log"
	"github.com/matthewmueller/golly/js"
	"github.com/pkg/errors"
)

func translatePackage(info *loader.PackageInfo) (j js.IExpression, err error) {
	var files []js.IStatement

	// translate each file into an array of statements
	for _, file := range info.Files {
		stmts, e := translateFile(info, file)
		if e != nil {
			return j, errors.Wrapf(e, "error translating %s", file.Name.Name)
		}
		files = append(files, stmts...)
	}

	// find all the exported fields
	exports := []string{}
	pkgscope := info.Pkg.Scope().Names()
	for _, name := range pkgscope {
		if name == "main" || unicode.IsUpper(rune(name[0])) {
			exports = append(exports, name)
		}
	}

	// create a return statement for the exported fieldss
	var props []js.Property
	for _, export := range exports {
		props = append(props, js.CreateProperty(
			js.CreateIdentifier(export),
			js.CreateIdentifier(export),
			"init",
		))
	}
	exportobj := js.CreateReturnStatement(
		js.CreateObjectExpression(props),
	)

	var body []interface{}
	for _, stmt := range append(files, exportobj) {
		body = append(body, stmt)
	}

	pkgfn := js.CreateCallExpression(
		js.CreateFunctionExpression(nil, []js.IPattern{},
			js.CreateFunctionBody(body...),
		),
		[]js.IExpression{},
	)

	return pkgfn, nil
}

// Translate from a Go AST to a Javascript AST
func translateFile(pkg *loader.PackageInfo, f *ast.File) (j []js.IStatement, err error) {
	// ast.SortImports(pkg, f)
	return fileToProgram(pkg, f)
	// return j, nil
}

func fileToProgram(pkg *loader.PackageInfo, f *ast.File) (j []js.IStatement, err error) {
	var stmts []js.IStatement
	for i := 0; i < len(f.Decls); i++ {
		stmt, e := decl(pkg, f, f.Decls[i])
		if e != nil {
			return j, e
		}
		stmts = append(stmts, stmt)
	}
	return stmts, nil
	// call main()
	// get the public functions, variables and types
	// exports, err := getExports(decls)
	// if err != nil {
	// 	return nil, errors.Wrap(err, "unable to get exports")
	// }

	// range over the file's scope objects
	// var exports []string
	// for name := range f.Scope.Objects {
	// 	if name == "main" || unicode.IsUpper(rune(name[0])) {
	// 		exports = append(exports, name)
	// 	}
	// }

	// pretty.Println(f.Scope.Objects)
	// // get the public variables
	// exports := pkg.Scopes[f].Parent().Names()
	// f.Scope.Objects
	// var props []js.Property
	// for _, export := range exports {
	// 	props = append(props, js.CreateProperty(
	// 		js.CreateIdentifier(export),
	// 		js.CreateIdentifier(export),
	// 		"init",
	// 	))
	// }
	// exportobj := js.CreateReturnStatement(js.CreateObjectExpression(props))

	// log.Infof("exports %+v", exports)
	// for i := 0; i < l; i++ {
	// 	log.Infof("decl %s", decls[i])

	// }
	// callmain := js.CreateExpressionStatement(
	// 	js.CreateCallExpression(
	// 		js.CreateIdentifier("main"),
	// 		[]js.IExpression{},
	// 	),
	// )

	// return js.CreateCallExpression(
	// 	js.CreateFunctionExpression(nil, []js.IPattern{},
	// 		js.CreateFunctionBody(append(stmts, exportobj)...),
	// 	),
	// 	[]js.IExpression{},
	// ), nil
}

// TODO: there's probably a better way to do this
// func getExports(decls []ast.Decl) (exports []string, err error) {
// 	for _, n := range decls {
// 		switch t := n.(type) {
// 		case *ast.FuncDecl:
// 			// check if uppercase
// 			if t.Name == nil || len(t.Name.Name) == 0 || !unicode.IsUpper(rune(t.Name.Name[0])) {
// 				continue
// 			}
// 			exports = append(exports, t.Name.Name)
// 		case *ast.GenDecl:
// 			switch t.Tok.String() {
// 			case "var":
// 				// case ""
// 			}
// 		default:
// 			return nil, unhandled("getExports", n)
// 		}
// 	}

// 	return exports, nil
// }

func decl(pkg *loader.PackageInfo, f *ast.File, n ast.Decl) (s js.IStatement, err error) {
	switch d := n.(type) {
	case *ast.FuncDecl:
		return function(pkg, f, d)
	case *ast.GenDecl:
		return genDecl(pkg, f, d)
	default:
		return nil, unhandled("decl", n)
	}
}

func function(pkg *loader.PackageInfo, f *ast.File, n *ast.FuncDecl) (j js.IStatement, err error) {
	// e.g. func hi()
	if n.Body == nil {
		return js.CreateEmptyStatement(), nil
	}

	// anonymous fns
	if n.Name == nil {
		return j, fmt.Errorf("function: anon functions not yet supported")
	}

	// the name of the fn
	fnname := js.CreateIdentifier((*n.Name).Name)

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
		jsStmt, e := funcStatement(pkg, f, n, stmt)
		if e != nil {
			return j, e
		}
		body = append(body, jsStmt)
	}

	// no receiver means it's a classless function
	if n.Recv == nil {
		return js.CreateFunctionDeclaration(
			&fnname,
			params,
			js.CreateFunctionBody(body...),
		), nil
	}

	if len(n.Recv.List) != 1 {
		return nil, fmt.Errorf("function<recv>: only expected 1 func receiver but got %d", len(n.Recv.List))
	}

	recv := n.Recv.List[0]
	x, e := expression(pkg, f, n, recv.Type)
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

func genDecl(pkg *loader.PackageInfo, f *ast.File, n *ast.GenDecl) (j js.IStatement, err error) {
	switch n.Tok.String() {
	case "import":
		return importSpec(pkg, f, n)
	case "type":
		return typeSpec(pkg, f, n)
	case "var":
		return varSpec(pkg, f, n)
	default:
		return nil, fmt.Errorf("genDecl: unhandled token %s", n.Tok.String())
	}

	// // specs := n.Specs
	// for _, spec := range n.Specs {
	// 	switch t := spec.(type) {
	// 	// case *ast.ImportSpec:
	// 	// 	return importSpec(pkg, f, t)
	// 	case *ast.TypeSpec:
	// 		// type defs will only have 1 spec
	// 		return typeSpec(pkg, f, t)
	// 	default:
	// 		return nil, unhandled("genDecl", spec)
	// 	}
	// }

	// return js.CreateEmptyStatement(), nil
}

// func importSpec(pkg *loader.PackageInfo,  f *ast.File, n *ast.ImportSpec) (j js.IStatement, err error) {
// 	return nil, nil
// }

func typeSpec(pkg *loader.PackageInfo, f *ast.File, n *ast.GenDecl) (j js.IStatement, err error) {
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
		return jsConstructorFunction(pkg, f, s.Name, t.Fields)
	default:
		return nil, unhandled("typeSpec", s.Type)
	}
}

func importSpec(pkg *loader.PackageInfo, f *ast.File, n *ast.GenDecl) (j js.IStatement, err error) {
	var decls []js.VariableDeclarator

	for _, spec := range n.Specs {
		switch t := spec.(type) {
		case *ast.ImportSpec:
			var lh js.Identifier
			p := t.Path.Value

			// import dep "pkg/dep"
			if t.Name != nil {
				lh = js.CreateIdentifier(t.Name.Name)
			} else if p != "" {
				lh = js.CreateIdentifier(path.Base(strings.Trim(p, `"`)))
			} else {
				return nil, unhandled("importSpec<empty>", spec)
			}

			rh := js.CreateMemberExpression(
				js.CreateIdentifier("pkg"),
				js.CreateString(t.Path.Value),
				true,
			)

			decl := js.CreateVariableDeclarator(lh, rh)
			decls = append(decls, decl)
		default:
			return nil, unhandled("importSpec", spec)
		}
	}

	return js.CreateVariableDeclaration("var", decls...), nil
}

func varSpec(pkg *loader.PackageInfo, f *ast.File, n *ast.GenDecl) (j js.IStatement, err error) {
	var decls []js.VariableDeclarator

	for _, spec := range n.Specs {
		switch t := spec.(type) {
		case *ast.ValueSpec:
			if len(t.Names) != len(t.Values) {
				return nil, unhandled("varSpec<unbalanced>", spec)
			}

			// handle balanced destructuring
			for i, ident := range t.Names {
				lh := js.CreateIdentifier(ident.Name)

				rh, e := expression(pkg, f, nil, t.Values[i])
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

func jsConstructorFunction(pkg *loader.PackageInfo, f *ast.File, id *ast.Ident, fields *ast.FieldList) (j js.FunctionDeclaration, err error) {
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

// func mainFunc(pkg *loader.PackageInfo,  f *ast.File, fn *ast.FuncDecl) (j js.IStatement, err error) {
// 	// e.g. func main()
// 	if fn.Body == nil {
// 		return js.CreateEmptyStatement(), nil
// 	}

// 	// TODO: []IStatement{} instead of interface{}
// 	var body []interface{}
// 	for _, stmt := range fn.Body.List {
// 		jsStmt, e := funcStatement(pkg, f, fn, stmt)
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

func funcStatement(pkg *loader.PackageInfo, f *ast.File, fn *ast.FuncDecl, stmt ast.Stmt) (j js.IStatement, err error) {
	switch t := stmt.(type) {
	case *ast.ExprStmt:
		return exprStatement(pkg, f, fn, t)
	case *ast.IfStmt:
		return ifStmt(pkg, f, fn, t)
	case *ast.AssignStmt:
		return assignStatement(pkg, f, fn, t)
	case *ast.ReturnStmt:
		return returnStatement(pkg, f, fn, t)
	case *ast.ForStmt:
		return forStmt(pkg, f, fn, t)
	case *ast.DeclStmt:
		return declStmt(pkg, f, fn, t)
	default:
		return nil, unhandled("funcStatement", stmt)
	}
}

func declStmt(pkg *loader.PackageInfo, f *ast.File, fn *ast.FuncDecl, n *ast.DeclStmt) (j js.IStatement, err error) {
	switch t := n.Decl.(type) {
	case *ast.GenDecl:
		return genDecl(pkg, f, t)
	default:
		return nil, unhandled("declStmt", n)
	}
}

func exprStatement(pkg *loader.PackageInfo, f *ast.File, fn *ast.FuncDecl, expr *ast.ExprStmt) (j js.IExpressionStatement, err error) {
	switch t := expr.X.(type) {
	case *ast.CallExpr:
		x, e := callExpression(pkg, f, fn, t)
		if e != nil {
			return nil, e
		}
		return js.CreateExpressionStatement(x), nil
	default:
		return nil, fmt.Errorf("exprStatement: unhandled type: %s", reflect.TypeOf(expr))
	}
}

func ifStmt(pkg *loader.PackageInfo, f *ast.File, fn *ast.FuncDecl, n *ast.IfStmt) (j js.IStatement, err error) {
	var e error

	// condition: if [(...)] { ... } else { ... }
	var test js.IExpression
	switch t := n.Cond.(type) {
	case *ast.BinaryExpr:
		test, e = binaryExpression(pkg, f, fn, t)
	case *ast.Ident:
		test, e = identifier(pkg, f, fn, t)
	default:
		return nil, unhandled("ifStmt<cond>", n.Cond)
	}
	if e != nil {
		return nil, e
	}

	// body : if (...) [{ ... }] else { ... }
	body, e := blockStmt(pkg, f, fn, n.Body)
	if e != nil {
		return nil, e
	}

	// else: if (...) { ... } else [{ ... }]
	elseBlock := n.Else
	var alt js.IStatement
	switch t := elseBlock.(type) {
	case *ast.BlockStmt:
		alt, e = blockStmt(pkg, f, fn, t)
	case *ast.ExprStmt:
		alt, e = exprStatement(pkg, f, fn, t)
	case *ast.IfStmt:
		alt, e = ifStmt(pkg, f, fn, t)
	case *ast.AssignStmt:
		alt, e = assignStatement(pkg, f, fn, t)
	case *ast.ReturnStmt:
		alt, e = returnStatement(pkg, f, fn, t)
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

func forStmt(pkg *loader.PackageInfo, f *ast.File, fn *ast.FuncDecl, n *ast.ForStmt) (j js.IStatement, err error) {

	init, e := statement(pkg, f, fn, n.Init)
	if e != nil {
		return nil, errors.Wrap(e, "forStmt")
	}

	cond, e := expression(pkg, f, fn, n.Cond)
	if e != nil {
		return nil, errors.Wrap(e, "forStmt")
	}

	post, e := statement(pkg, f, fn, n.Post)
	if e != nil {
		return nil, errors.Wrap(e, "forStmt")
	}

	body, e := blockStmt(pkg, f, fn, n.Body)
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

func statement(pkg *loader.PackageInfo, f *ast.File, fn *ast.FuncDecl, n ast.Stmt) (j js.IStatement, err error) {
	switch t := n.(type) {
	case nil:
		return nil, nil
	case *ast.AssignStmt:
		return assignStatement(pkg, f, fn, t)
	case *ast.IncDecStmt:
		return incDecStmt(pkg, f, fn, t)
	case *ast.ExprStmt:
		return exprStatement(pkg, f, fn, t)
	default:
		return nil, unhandled("statement", n)
	}
}

func blockStmt(pkg *loader.PackageInfo, f *ast.File, fn *ast.FuncDecl, n *ast.BlockStmt) (j js.IBlockStatement, err error) {
	var stmts []js.IStatement

	for _, stmt := range n.List {
		v, e := statement(pkg, f, fn, stmt)
		if e != nil {
			return nil, errors.Wrap(e, "blockStmt")
		}
		stmts = append(stmts, v)
	}

	return js.CreateBlockStatement(stmts...), nil
}

func assignStatement(pkg *loader.PackageInfo, f *ast.File, fn *ast.FuncDecl, as *ast.AssignStmt) (j js.IStatement, err error) {
	lhs := as.Lhs
	rhs := as.Rhs

	// handle balanced destructuring
	var decls []js.VariableDeclarator
	if len(lhs) == len(rhs) {
		for i, lh := range lhs {
			var id js.IPattern
			var e error

			rh, e := expression(pkg, f, fn, rhs[i])
			if e != nil {
				return j, e
			}

			switch t := lh.(type) {
			case *ast.Ident:
				id = js.CreateIdentifier(t.Name)
			// exit early if it's a member expression
			// e.g. document.test = 'hi'
			// TODO: cleanup
			case *ast.SelectorExpr:
				memberExpr, e := selectorExpr(pkg, f, fn, t)
				if e != nil {
					return j, e
				}
				return js.CreateExpressionStatement(
					js.CreateAssignmentExpression(
						memberExpr,
						js.AssignmentOperator("="),
						rh,
					),
				), nil
			default:
				return j, fmt.Errorf("assignStatement: unhandled type: %s", reflect.TypeOf(lh))
			}
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

		rh, e := expression(pkg, f, fn, rhs[0])
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

func incDecStmt(pkg *loader.PackageInfo, f *ast.File, fn *ast.FuncDecl, n *ast.IncDecStmt) (j js.IStatement, err error) {
	var op js.UpdateOperator

	x, e := expression(pkg, f, fn, n.X)
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

func returnStatement(pkg *loader.PackageInfo, f *ast.File, fn *ast.FuncDecl, n *ast.ReturnStmt) (j js.IStatement, err error) {
	// no return values
	if len(n.Results) == 0 {
		return js.CreateReturnStatement(nil), nil
	}

	var args []js.IExpression
	for _, arg := range n.Results {
		a, e := expression(pkg, f, fn, arg)
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

func callExpression(pkg *loader.PackageInfo, f *ast.File, fn *ast.FuncDecl, expr *ast.CallExpr) (j js.IExpression, err error) {
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

	callee, e := expression(pkg, f, fn, expr.Fun)
	if e != nil {
		return j, e
	}

	var args []js.IExpression
	for _, arg := range expr.Args {
		v, e := expression(pkg, f, fn, arg)
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

func expression(pkg *loader.PackageInfo, f *ast.File, fn *ast.FuncDecl, expr ast.Expr) (j js.IExpression, err error) {
	switch t := expr.(type) {
	case *ast.Ident:
		return identifier(pkg, f, fn, t)
	case *ast.BasicLit:
		return basiclit(pkg, f, fn, t)
	case *ast.CallExpr:
		return callExpression(pkg, f, fn, t)
	case *ast.BinaryExpr:
		return binaryExpression(pkg, f, fn, t)
	case *ast.CompositeLit:
		return compositeLiteral(pkg, f, fn, t)
	case *ast.SelectorExpr:
		return selectorExpr(pkg, f, fn, t)
	case *ast.IndexExpr:
		return indexExpr(pkg, f, fn, t)
	case *ast.StarExpr:
		return starExpr(pkg, f, fn, t)
	case *ast.UnaryExpr:
		return unaryExpr(pkg, f, fn, t)
	case *ast.FuncLit:
		return funcLit(pkg, f, fn, t)
	// case *ast.KeyValueExpr:
	// 	return keyValueExpr(pkg, f, fn, t)
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

func funcLit(pkg *loader.PackageInfo, f *ast.File, fn *ast.FuncDecl, n *ast.FuncLit) (j js.IExpression, err error) {
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
		jsStmt, e := funcStatement(pkg, f, fn, stmt)
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
func binaryExpression(pkg *loader.PackageInfo, f *ast.File, fn *ast.FuncDecl, b *ast.BinaryExpr) (j js.IExpression, err error) {
	x, e := expression(pkg, f, fn, b.X)
	if e != nil {
		return nil, e
	}
	y, e := expression(pkg, f, fn, b.Y)
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

func identifier(pkg *loader.PackageInfo, f *ast.File, fn *ast.FuncDecl, ident *ast.Ident) (j js.IExpression, err error) {
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

func starExpr(pkg *loader.PackageInfo, f *ast.File, fn *ast.FuncDecl, n *ast.StarExpr) (j js.IExpression, err error) {
	// for now, we're ignoring the pointer
	x, e := expression(pkg, f, fn, n.X)
	if e != nil {
		return nil, e
	}

	return x, nil
}

func unaryExpr(pkg *loader.PackageInfo, f *ast.File, fn *ast.FuncDecl, n *ast.UnaryExpr) (j js.IExpression, err error) {
	// for now, we're ignoring the pointer
	x, e := expression(pkg, f, fn, n.X)
	if e != nil {
		return nil, e
	}

	return x, nil
}

func basiclit(pkg *loader.PackageInfo, f *ast.File, fn *ast.FuncDecl, lit *ast.BasicLit) (j js.IExpression, err error) {
	return js.CreateString(lit.Value), nil
}

func compositeLiteral(pkg *loader.PackageInfo, f *ast.File, fn *ast.FuncDecl, n *ast.CompositeLit) (j js.IExpression, err error) {
	switch n.Type.(type) {
	case *ast.Ident, *ast.SelectorExpr:
		return jsNewFunction(pkg, f, fn, n)
	case *ast.ArrayType:
		return jsArrayExpression(pkg, f, fn, n)
	case *ast.MapType:
		return jsObjectExpression(pkg, f, fn, n)
	default:
		return nil, unhandled("compositeLiteral<type>", n.Type)
	}
}

func jsObjectExpression(pkg *loader.PackageInfo, f *ast.File, fn *ast.FuncDecl, n *ast.CompositeLit) (j js.ObjectExpression, err error) {
	var props []js.Property

	for _, elt := range n.Elts {
		var prop js.Property
		var e error

		switch t := elt.(type) {
		case *ast.KeyValueExpr:
			prop, e = keyValueExpr(pkg, f, fn, t)
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

func jsNewFunction(pkg *loader.PackageInfo, f *ast.File, fn *ast.FuncDecl, n *ast.CompositeLit) (j js.NewExpression, err error) {
	var fnname js.IExpression
	var props []js.Property

	switch t := n.Type.(type) {
	// e.g. Document{}
	case *ast.Ident:
		fnname = js.CreateIdentifier(t.Name)
	// e.g. dom.Document{}
	case *ast.SelectorExpr:
		sel, e := selectorExpr(pkg, f, fn, t)
		if e != nil {
			return j, e
		}
		fnname = sel
	default:
		return j, unhandled("jsNewFunction<name>", n.Type)
	}

	for _, elt := range n.Elts {
		var prop js.Property
		var e error

		switch t := elt.(type) {
		case *ast.KeyValueExpr:
			prop, e = keyValueExpr(pkg, f, fn, t)
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
		fnname,
		[]js.IExpression{js.CreateObjectExpression(props)},
	), nil
}

func jsArrayExpression(pkg *loader.PackageInfo, f *ast.File, fn *ast.FuncDecl, n *ast.CompositeLit) (j js.ArrayExpression, err error) {
	var elements []js.IExpression

	for _, elt := range n.Elts {
		el, e := expression(pkg, f, fn, elt)
		if e != nil {
			return j, e
		}
		elements = append(elements, el)
	}

	return js.CreateArrayExpression(elements...), nil
}

func keyValueExpr(pkg *loader.PackageInfo, f *ast.File, fn *ast.FuncDecl, n *ast.KeyValueExpr) (j js.Property, err error) {
	// get the key
	key, e := expression(pkg, f, fn, n.Key)
	if e != nil {
		return j, e
	}

	// get the value
	value, e := expression(pkg, f, fn, n.Value)
	if e != nil {
		return j, e
	}

	return js.CreateProperty(key, value, "init"), nil
}

func selectorExpr(pkg *loader.PackageInfo, f *ast.File, fn *ast.FuncDecl, n *ast.SelectorExpr) (j js.MemberExpression, err error) {
	// (user.phone).number
	x, e := expression(pkg, f, fn, n.X)
	if e != nil {
		return j, e
	}

	// user.phone.(number)
	s, e := identifier(pkg, f, fn, n.Sel)
	if e != nil {
		return j, e
	}

	return js.CreateMemberExpression(x, s, false), nil
}

func indexExpr(pkg *loader.PackageInfo, f *ast.File, fn *ast.FuncDecl, n *ast.IndexExpr) (j js.MemberExpression, err error) {
	// (i)[0]
	x, e := expression(pkg, f, fn, n.X)
	if e != nil {
		return j, e
	}

	// i([0])
	i, e := expression(pkg, f, fn, n.Index)
	if e != nil {
		return j, e
	}

	return js.CreateMemberExpression(x, i, true), nil
}

func unhandled(fn string, n interface{}) error {
	return fmt.Errorf("%s in %s() is not implemented yet", reflect.TypeOf(n), fn)
}
