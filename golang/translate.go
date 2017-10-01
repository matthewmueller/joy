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
	"github.com/fatih/structtag"
	"github.com/matthewmueller/golly/jsast"
	"github.com/pkg/errors"
)

// TODO: should be within struct, not global
var aliases = map[string]*structtag.Tag{}

func translatePackage(info *loader.PackageInfo) (j jsast.IExpression, err error) {
	var files []jsast.IStatement

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
	pkgscope := info.Pkg.Scope()
	for _, name := range pkgscope.Names() {
		obj := pkgscope.Lookup(name)

		// don't export globals
		if obj != nil &&
			obj.Type() != nil &&
			aliases[obj.Type().String()] != nil &&
			aliases[obj.Type().String()].HasOption("global") {
			continue
		}

		if name == "main" || unicode.IsUpper(rune(name[0])) {
			exports = append(exports, name)
		}
	}

	var exportobj jsast.IStatement

	// create a return statement for the exported fields
	if len(exports) > 0 {
		var props []jsast.Property
		for _, export := range exports {
			props = append(props, jsast.CreateProperty(
				jsast.CreateIdentifier(export),
				jsast.CreateIdentifier(export),
				"init",
			))
		}
		exportobj = jsast.CreateReturnStatement(
			jsast.CreateObjectExpression(props),
		)
	} else {
		exportobj = jsast.CreateEmptyStatement()
	}

	var body []interface{}
	for _, stmt := range append(files, exportobj) {
		body = append(body, stmt)
	}

	// creates a self-invoking function containing our package
	// e.g. (function() { ... })()
	pkgfn := jsast.CreateCallExpression(
		jsast.CreateFunctionExpression(nil, []jsast.IPattern{},
			jsast.CreateFunctionBody(body...),
		),
		[]jsast.IExpression{},
	)

	return pkgfn, nil
}

func translateFile(pkg *loader.PackageInfo, f *ast.File) (j []jsast.IStatement, err error) {
	var stmts []jsast.IStatement
	for i := 0; i < len(f.Decls); i++ {
		stmt, e := decl(pkg, f, f.Decls[i])
		if e != nil {
			return j, e
		}
		stmts = append(stmts, stmt)
	}
	return stmts, nil
}

func decl(pkg *loader.PackageInfo, f *ast.File, n ast.Decl) (s jsast.IStatement, err error) {
	switch d := n.(type) {
	case *ast.FuncDecl:
		return funcDecl(pkg, f, d)
	case *ast.GenDecl:
		return genDecl(pkg, f, d)
	default:
		return nil, unhandled("decl", n)
	}
}

func funcDecl(pkg *loader.PackageInfo, f *ast.File, n *ast.FuncDecl) (j jsast.IStatement, err error) {
	// e.g. func hi()
	if n.Body == nil {
		return jsast.CreateEmptyStatement(), nil
	}

	// get js:"..." tag on top of function if there is one
	tag, e := getCommentTag(n.Doc)
	if e != nil {
		return j, e
	}

	// get the fullname of this function
	obj := pkg.ObjectOf(n.Name)
	if obj == nil {
		log.Warnf("no object of")
	} else {
		log.Infof(obj.String())
	}

	// potentially rename functions
	if tag != nil && obj != nil {
		aliases[obj.String()] = tag
	}

	// get the name of the function
	fnname := jsast.CreateIdentifier((*n.Name).Name)

	// build argument list
	// var args
	var params []jsast.IPattern
	if n.Type != nil && n.Type.Params != nil {
		fields := n.Type.Params.List
		for _, field := range fields {
			// names because: (a, b string, c int) = [[a, b], c]
			for _, name := range field.Names {
				id := jsast.CreateIdentifier(name.Name)
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
		return jsast.CreateFunction(
			&fnname,
			params,
			jsast.CreateFunctionBody(body...),
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

	// This is to determine if the function should be excluded
	// from the final output because it's already a global DOM node.
	//
	// Lookup the receiver's type, note that this has nothing to do
	// with whether or not there's a comment on the function
	// itself.
	//
	// TODO: This code may panic, make more robust
	if len(recv.Names) > 0 {
		typeof := pkg.TypeOf(recv.Names[0])
		if typeof != nil && aliases[typeof.String()] != nil && aliases[typeof.String()].HasOption("global") {
			return jsast.CreateEmptyStatement(), nil
		}
	}

	// {recv}.prototype.{name} = function ({params}) {
	//   {body}
	// }
	return jsast.CreateExpressionStatement(
		jsast.CreateAssignmentExpression(
			jsast.CreateMemberExpression(
				jsast.CreateMemberExpression(
					x,
					jsast.CreateIdentifier("prototype"),
					false,
				),
				fnname,
				false,
			),
			jsast.AssignmentOperator("="),
			jsast.CreateFunctionExpression(
				nil,
				params,
				jsast.CreateFunctionBody(body...),
			),
		),
	), nil

}

func genDecl(pkg *loader.PackageInfo, f *ast.File, n *ast.GenDecl) (j jsast.IStatement, err error) {
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

	// return jsast.CreateEmptyStatement(), nil
}

// func importSpec(pkg *loader.PackageInfo,  f *ast.File, n *ast.ImportSpec) (j jsast.IStatement, err error) {
// 	return nil, nil
// }

func typeSpec(pkg *loader.PackageInfo, f *ast.File, n *ast.GenDecl) (j jsast.IStatement, err error) {
	if len(n.Specs) != 1 {
		return nil, fmt.Errorf("genDecl: expected type to only have 1 spec but it has %d", len(n.Specs))
	}

	s, ok := n.Specs[0].(*ast.TypeSpec)
	if !ok {
		return nil, unhandled("typeSpec<TypeSpec>", n.Specs[0])
	}

	var st *ast.StructType
	switch t := s.Type.(type) {
	case *ast.StructType:
		st = t
	case *ast.InterfaceType:
		log.Warnf("ignoring interface type. TODO: not sure if these would always be excluded from JS")
		return jsast.CreateEmptyStatement(), nil
	default:
		return nil, unhandled("typeSpec<StructType>", s.Type)
	}

	tag, e := getCommentTag(n.Doc)
	if e != nil {
		return nil, e
	}

	// store the tag for later renaming
	objectof := pkg.ObjectOf(s.Name)
	typeof := pkg.TypeOf(s.Name)
	if tag != nil && objectof != nil {
		aliases[objectof.String()] = tag
		aliases[typeof.String()] = tag
		// TODO: not sure if this is a good idea or not
		// but it's to handle pointer receivers in 1 spot
		aliases["*"+typeof.String()] = tag
	}

	var ivars []interface{}
	for _, field := range st.Fields.List {
		for _, name := range field.Names {
			objectof := pkg.ObjectOf(name)
			// TODO: this doesn't need to be parsed for each field name
			// though I'm not sure how you can have multiple field names
			// per struct tag
			if field.Tag != nil {
				tags, err := structtag.Parse(field.Tag.Value[1 : len(field.Tag.Value)-1])
				if err != nil {
					return j, err
				}
				if tag, err := tags.Get("js"); err == nil {
					aliases[objectof.String()] = tag
				}
			}

			// this.$name = o.$name
			ivars = append(ivars, jsast.CreateExpressionStatement(
				jsast.CreateAssignmentExpression(
					jsast.CreateMemberExpression(
						jsast.CreateThisExpression(),
						jsast.CreateIdentifier(name.Name),
						false,
					),
					jsast.AssignmentOperator("="),
					jsast.CreateMemberExpression(
						jsast.CreateIdentifier("o"),
						jsast.CreateIdentifier(name.Name),
						false,
					),
				),
			))
		}
	}

	if tag != nil && tag.HasOption("global") {
		return jsast.CreateEmptyStatement(), nil
	}

	ident := jsast.CreateIdentifier(s.Name.Name)
	return jsast.CreateFunction(
		&ident,
		// TODO: make API for this
		[]jsast.IPattern{jsast.CreateIdentifier("o")},
		jsast.CreateFunctionBody(ivars...),
	), nil
}

func importSpec(pkg *loader.PackageInfo, f *ast.File, n *ast.GenDecl) (j jsast.IStatement, err error) {
	var decls []jsast.VariableDeclarator

	for _, spec := range n.Specs {
		switch t := spec.(type) {
		case *ast.ImportSpec:
			var lh jsast.Identifier
			p := t.Path.Value

			// import dep "pkg/dep"
			if t.Name != nil {
				lh = jsast.CreateIdentifier(t.Name.Name)
			} else if p != "" {
				lh = jsast.CreateIdentifier(path.Base(strings.Trim(p, `"`)))
			} else {
				return nil, unhandled("importSpec<empty>", spec)
			}

			rh := jsast.CreateMemberExpression(
				jsast.CreateIdentifier("pkg"),
				jsast.CreateLiteral(t.Path.Value),
				true,
			)

			decl := jsast.CreateVariableDeclarator(lh, rh)
			decls = append(decls, decl)
		default:
			return nil, unhandled("importSpec", spec)
		}
	}

	return jsast.CreateVariableDeclaration("var", decls...), nil
}

func varSpec(pkg *loader.PackageInfo, f *ast.File, n *ast.GenDecl) (j jsast.IStatement, err error) {
	var decls []jsast.VariableDeclarator

	for _, spec := range n.Specs {
		switch t := spec.(type) {
		case *ast.ValueSpec:
			lval := len(t.Values)

			// handle balanced destructuring
			for i, ident := range t.Names {
				lh := jsast.CreateIdentifier(ident.Name)

				var rh jsast.IExpression
				if i < lval {
					r, e := expression(pkg, f, nil, t.Values[i])
					if e != nil {
						return j, e
					}
					rh = r
				} else {
					r, e := expression(pkg, f, nil, t.Type)
					if e != nil {
						return nil, e
					}
					rh = r
				}

				decl := jsast.CreateVariableDeclarator(lh, rh)
				decls = append(decls, decl)
			}

		default:
			return nil, unhandled("varSpec", spec)
		}
	}

	return jsast.CreateVariableDeclaration("var", decls...), nil
	// return nil, nil
}

// func mainFunc(pkg *loader.PackageInfo,  f *ast.File, fn *ast.FuncDecl) (j jsast.IStatement, err error) {
// 	// e.g. func main()
// 	if fn.Body == nil {
// 		return jsast.CreateEmptyStatement(), nil
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

// 	return jsast.CreateExpressionStatement(
// 		jsast.CreateCallExpression(
// 			jsast.CreateFunctionExpression(nil, []jsast.IPattern{},
// 				jsast.CreateFunctionBody(body...),
// 			),
// 			[]jsast.IExpression{},
// 		),
// 	), nil
// 	// return j, nil
// }

func funcStatement(pkg *loader.PackageInfo, f *ast.File, fn *ast.FuncDecl, stmt ast.Stmt) (j jsast.IStatement, err error) {
	switch t := stmt.(type) {
	case *ast.ExprStmt:
		return exprStatement(pkg, f, fn, t)
	case *ast.IfStmt:
		return ifStmt(pkg, f, fn, t)
	case *ast.AssignStmt:
		return assignStatement(pkg, f, fn, t)
	case *ast.ReturnStmt:
		return returnStmt(pkg, f, fn, t)
	case *ast.RangeStmt:
		return rangeStmt(pkg, f, fn, t)
	case *ast.ForStmt:
		return forStmt(pkg, f, fn, t)
	case *ast.DeclStmt:
		return declStmt(pkg, f, fn, t)
	case *ast.GoStmt:
		return goStmt(pkg, f, fn, t)
	case *ast.SendStmt:
		return sendStmt(pkg, f, fn, t)
	default:
		return nil, unhandled("funcStatement", stmt)
	}
}

func goStmt(pkg *loader.PackageInfo, f *ast.File, fn *ast.FuncDecl, n *ast.GoStmt) (j jsast.IStatement, err error) {
	c, e := callExpression(pkg, f, fn, n.Call)
	if e != nil {
		return nil, e
	}

	// ast.
	// log.Infof("Scope %+v", f.Scope.Lookup(n))

	return jsast.CreateExpressionStatement(c), nil
	// return nil, nil
}

func sendStmt(pkg *loader.PackageInfo, f *ast.File, fn *ast.FuncDecl, n *ast.SendStmt) (j jsast.IStatement, err error) {
	ch, e := expression(pkg, f, fn, n.Chan)
	if e != nil {
		return nil, e
	}

	val, e := expression(pkg, f, fn, n.Value)
	if e != nil {
		return nil, e
	}

	return jsast.CreateExpressionStatement(
		jsast.CreateAwaitExpression(
			jsast.CreateCallExpression(
				jsast.CreateMemberExpression(
					ch,
					jsast.CreateIdentifier("Send"),
					false,
				),
				[]jsast.IExpression{val},
			),
		),
	), nil
}

func rangeStmt(pkg *loader.PackageInfo, f *ast.File, fn *ast.FuncDecl, n *ast.RangeStmt) (j jsast.IStatement, err error) {
	id, ok := n.Key.(*ast.Ident)
	if !ok {
		return nil, unhandled("rangeStmt<ident>", n.Key)
	}

	if id.Obj == nil {
		return nil, unhandled("rangeStmt<obj>", id.Obj)
	}

	asn, ok := id.Obj.Decl.(*ast.AssignStmt)
	if !ok {
		return nil, unhandled("rangeStmt<decl>", id.Obj)
	}

	if len(asn.Lhs) == 0 {
		return nil, fmt.Errorf("rangeStmt: didn't expect len(lhs) == 0")
	}

	// create the condition
	if len(asn.Rhs) == 0 {
		return nil, fmt.Errorf("rangeStmt: didn't expect len(rhs) == 0")
	}

	rh, e := expression(pkg, f, fn, asn.Rhs[0])
	if e != nil {
		return nil, e
	}

	var inits []jsast.VariableDeclarator
	idx, ok := asn.Lhs[0].(*ast.Ident)
	if !ok {
		return nil, unhandled("rangeStmt<idx>", asn.Lhs[0])
	}
	inits = append(inits, jsast.CreateVariableDeclarator(
		jsast.CreateIdentifier(idx.Name),
		jsast.CreateInt(0),
	))

	var val *ast.Ident
	if len(asn.Lhs) >= 2 {
		val, ok = asn.Lhs[1].(*ast.Ident)
		if !ok {
			return nil, unhandled("rangeStmt<val>", asn.Lhs[1])
		}
		inits = append(inits, jsast.CreateVariableDeclarator(
			jsast.CreateIdentifier(val.Name),
			nil,
		))
	}

	init := jsast.CreateVariableDeclaration("var", inits...)

	cond := jsast.CreateBinaryExpression(
		jsast.CreateIdentifier(idx.Name),
		jsast.BinaryOperator("<"),
		jsast.CreateMemberExpression(
			rh,
			jsast.CreateIdentifier("length"),
			false,
		),
	)

	postexpr := jsast.CreateUpdateExpression(
		jsast.CreateIdentifier(idx.Name),
		jsast.UpdateOperator("++"),
		false,
	)

	// build the body
	var stmts []jsast.IStatement
	if val != nil {
		stmts = append(stmts, jsast.CreateVariableDeclaration(
			"var",
			jsast.CreateVariableDeclarator(
				jsast.CreateIdentifier(val.Name),
				jsast.CreateMemberExpression(
					rh,
					jsast.CreateIdentifier(idx.Name),
					true,
				),
			),
		))
	}
	for _, stmt := range n.Body.List {
		v, e := statement(pkg, f, fn, stmt)
		if e != nil {
			return nil, errors.Wrap(e, "rangeStmt<body>")
		}
		stmts = append(stmts, v)
	}
	body := jsast.CreateBlockStatement(stmts...)

	// TODO:
	// Range expression                              1st value          2nd value
	// [x] array or slice  a  [n]E, *[n]E, or []E    index    i  int    a[i]       E
	// [ ] string          s  string type            index    i  int    see below  rune
	// [ ] map             m  map[K]V                key      k  K      m[k]       V
	// [ ] channel         c  chan E, <-chan E       element  e  E
	kind, e := getObjectType(asn.Rhs[0])
	if e != nil {
		return nil, e
	}

	// ast.
	// ast.Print(nil, asn.Rhs[0])

	switch kind.(type) {
	case *ast.ArrayType:
		return jsast.CreateForStatement(
			init,
			cond,
			postexpr,
			body,
		), nil
	default:
		return nil, unhandled("rangeStmt<rhs.obj.type>", id.Obj)
	}

	// switch asn.Rhs[0].

	// // parse lefthand side
	// lhs := asn.Lhs
	// llhs := len(asn.Lhs)

	// if llhs == 0 {
	// 	return nil, fmt.Errorf("rangeStmt: didn't expect len(lhs) == 0")
	// }

	// var inits []jsast.VariableDeclarator

	// idx, ok := lhs[0].(*ast.Ident)
	// if !ok {
	// 	return nil, unhandled("rangeStmt<idx>", lhs[0])
	// }
	// inits = append(inits, jsast.CreateVariableDeclarator(
	// 	jsast.CreateIdentifier(idx.Name),
	// 	jsast.CreateInt(0),
	// ))

	// // handle the value if there is one
	// var val *ast.Ident
	// if llhs == 2 {
	// 	val, ok = lhs[1].(*ast.Ident)
	// 	if !ok {
	// 		return nil, unhandled("rangeStmt<val>", lhs[1])
	// 	}
	// 	inits = append(inits, jsast.CreateVariableDeclarator(
	// 		jsast.CreateIdentifier(val.Name),
	// 		nil,
	// 	))
	// }

	// ux, ok := asn.Rhs[0].(*ast.UnaryExpr)
	// if !ok {
	// 	return nil, unhandled("rangeStmt<rhs[0]>", asn.Rhs[0])
	// }

	// var rh jsast.IExpression
	// switch t := ux.X.(type) {
	// case *ast.Ident:
	// 	rh = jsast.CreateMemberExpression(
	// 		jsast.CreateIdentifier(t.Name),
	// 		jsast.CreateIdentifier("length"),
	// 		false,
	// 	)
	// default:
	// 	return nil, unhandled("rangeStmt<rh>", ux.X)
	// }

	// for _, lhs := range asn.Lhs {
	// 	switch t := lhs.(type) {
	// 	case *ast.Ident:
	// 		decls = append(decls, jsast.CreateVariableDeclarator(
	// 			jsast.CreateIdentifier(t.Name),
	// 			nil,
	// 		))
	// 	default:
	// 		return nil, unhandled("rangeStmt<lhs>", lhs)
	// 	}
	// }
	// init := jsast.CreateVariableDeclaration("var", decls...)
	// _ = init

	// stmt, expr, e := VariableHandler(id.Obj.Decl)
	// if e != nil {
	// 	return nil, e
	// }

	// log.Infof("stmt %s", stmt)
	// log.Infof("expr %s", expr)

	// variables
	// pairs := variablePairs(asn)

	// for _, asn := range asn.Lhs {
	// 	switch {
	// 		case *ast.Ident;

	// 	}
	// }

	// ast.Print(nil, asn)
	// log.Infof("asn %s", asn)
	// _ = asn
	// if
	// lhs := asn.Lhs
	// llhs := len(lhs)

	// for _,  := range s {

	// }
	// if llhs == 2 {
	// 	jsVariableDecl()
	// } else if llhs == 1 {

	// }
	// if asn.
	// asn.Lhs

	// a, e := assignStatement(pkg, f, fn, asn)
	// if e != nil {
	// 	return j, e
	// }
	// log.Infof("%s", a)

	// obj.

	// switch t := n.Key.(type) {
	// case *ast.Ident:
	// 	t.
	// }
	// ast.Print(nil, n.Key)
	// switch n.Key {
	// 	case
	// }

	// init, e := expression(pkg, f, fn, n.Key)
	// if e != nil {
	// 	return nil, errors.Wrap(e, "forStmt")
	// }
	// log.Infof("init %s", init)
	// cond, e := expression(pkg, f, fn, n.Cond)
	// if e != nil {
	// 	return nil, errors.Wrap(e, "forStmt")
	// }

	// post, e := statement(pkg, f, fn, n.Post)
	// if e != nil {
	// 	return nil, errors.Wrap(e, "forStmt")
	// }

	// body, e := blockStmt(pkg, f, fn, n.Body)
	// if e != nil {
	// 	return nil, errors.Wrap(e, "rangeStmt")
	// }
	// _ = body
	// In Go the post condition is a statement,
	// in JS it's an expression
	//
	// it can also be nil in the case of for { ... }
	// var postexpr jsast.IExpression
	// switch t := post.(type) {
	// case jsast.ExpressionStatement:
	// 	postexpr = t.Expression
	// case nil:
	// 	postexpr = nil
	// default:
	// 	return nil, unhandled("forStmt<post>", post)
	// }

	// return jsast.CreateEmptyStatement(), nil

	// return jsast.CreateForStatement(
	// 	init,
	// 	cond,
	// 	postexpr,
	// 	body,
	// ), nil
}

func declStmt(pkg *loader.PackageInfo, f *ast.File, fn *ast.FuncDecl, n *ast.DeclStmt) (j jsast.IStatement, err error) {
	switch t := n.Decl.(type) {
	case *ast.GenDecl:
		return genDecl(pkg, f, t)
	default:
		return nil, unhandled("declStmt", n)
	}
}

func exprStatement(pkg *loader.PackageInfo, f *ast.File, fn *ast.FuncDecl, expr *ast.ExprStmt) (j jsast.IExpressionStatement, err error) {
	switch t := expr.X.(type) {
	case *ast.CallExpr:
		x, e := callExpression(pkg, f, fn, t)
		if e != nil {
			return nil, e
		}
		return jsast.CreateExpressionStatement(x), nil
	default:
		return nil, fmt.Errorf("exprStatement: unhandled type: %s", reflect.TypeOf(expr))
	}
}

func ifStmt(pkg *loader.PackageInfo, f *ast.File, fn *ast.FuncDecl, n *ast.IfStmt) (j jsast.IStatement, err error) {
	var e error

	// condition: if [(...)] { ... } else { ... }
	var test jsast.IExpression
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
	var alt jsast.IStatement
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
		alt, e = returnStmt(pkg, f, fn, t)
	case nil:
		alt = jsast.CreateEmptyStatement()
	default:
		return nil, unhandled("ifStmt<else>", elseBlock)
	}
	if e != nil {
		return nil, e
	}

	return jsast.CreateIfStatement(
		test,
		body,
		alt,
	), nil
}

func branchStmt(pkg *loader.PackageInfo, f *ast.File, fn *ast.FuncDecl, n *ast.BranchStmt) (j jsast.IStatement, err error) {
	switch n.Tok.String() {
	case "break":
		return jsast.CreateBreakStatement(nil), nil
	default:
		return nil, fmt.Errorf("unhandled branchStmt: %s", n.Tok.String())
	}
}

func forStmt(pkg *loader.PackageInfo, f *ast.File, fn *ast.FuncDecl, n *ast.ForStmt) (j jsast.IStatement, err error) {

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
	var postexpr jsast.IExpression
	switch t := post.(type) {
	case jsast.ExpressionStatement:
		postexpr = t.Expression
	case nil:
		postexpr = nil
	default:
		return nil, unhandled("forStmt<post>", post)
	}

	return jsast.CreateForStatement(
		init,
		cond,
		postexpr,
		body,
	), nil
}

func statement(pkg *loader.PackageInfo, f *ast.File, fn *ast.FuncDecl, n ast.Stmt) (j jsast.IStatement, err error) {
	switch t := n.(type) {
	case nil:
		return nil, nil
	case *ast.AssignStmt:
		return assignStatement(pkg, f, fn, t)
	case *ast.IncDecStmt:
		return incDecStmt(pkg, f, fn, t)
	case *ast.ExprStmt:
		return exprStatement(pkg, f, fn, t)
	case *ast.IfStmt:
		return ifStmt(pkg, f, fn, t)
	case *ast.BranchStmt:
		return branchStmt(pkg, f, fn, t)
	case *ast.ReturnStmt:
		return returnStmt(pkg, f, fn, t)
	default:
		return nil, unhandled("statement", n)
	}
}

func blockStmt(pkg *loader.PackageInfo, f *ast.File, fn *ast.FuncDecl, n *ast.BlockStmt) (j jsast.IBlockStatement, err error) {
	var stmts []jsast.IStatement

	for _, stmt := range n.List {
		v, e := statement(pkg, f, fn, stmt)
		if e != nil {
			return nil, errors.Wrap(e, "blockStmt")
		}
		stmts = append(stmts, v)
	}

	return jsast.CreateBlockStatement(stmts...), nil
}

func assignStatement(pkg *loader.PackageInfo, f *ast.File, fn *ast.FuncDecl, as *ast.AssignStmt) (j jsast.IStatement, err error) {
	// TODO: these are separate, but very similar functions.
	// the reason they're separate is because the JS AST's are
	// different. It'd be good to come up with a way to consolidate
	// this logic though because this variable conversion is a bit tricky
	switch as.Tok.String() {
	case "=":
		return jsAssignStmt(pkg, f, fn, as)
	case ":=":
		return jsVariableDecl(pkg, f, fn, as)
	default:
		return nil, fmt.Errorf("unhandled assignStatement<tok>: %s", as.Tok.String())
	}
}

func jsAssignStmt(pkg *loader.PackageInfo, f *ast.File, fn *ast.FuncDecl, as *ast.AssignStmt) (j jsast.IStatement, err error) {
	var exprs []jsast.IExpression
	lhs := as.Lhs
	rhs := as.Rhs
	llhs := len(lhs)
	lrhs := len(rhs)

	// ensure we're not in an invalid state
	if llhs != lrhs && lrhs > 1 {
		return nil, fmt.Errorf("invalid golang assignment (AFAIK)")
	}

	// nothing on right side
	if lrhs == 0 {
		for _, lh := range lhs {
			l, e := expression(pkg, f, fn, lh)
			if e != nil {
				return nil, e
			}
			exprs = append(exprs, l)
		}
	}

	// balanced on both sides
	if llhs == lrhs {
		for i, lh := range lhs {
			if isUnderscoreVariable(lh) {
				continue
			}

			l, e := expression(pkg, f, fn, lh)
			if e != nil {
				return nil, e
			}

			r, e := expression(pkg, f, fn, rhs[i])
			if e != nil {
				return nil, e
			}

			exprs = append(exprs, jsast.CreateAssignmentExpression(
				l,
				jsast.AssignmentOperator("="),
				r,
			))
		}
	}

	// unbalanced
	if llhs != lrhs {
		var lname string

		if isUnderscoreVariable(lhs[0]) {
			return nil, unhandled("jsAssignStmt<underscore>", lhs[0])
		}

		switch t := lhs[0].(type) {
		case *ast.Ident:
			lname = "$" + t.Name
		// case *ast.SelectorExpr:
		// return unhandled("jsAssignStmt<selectorExpr>", t)
		default:
			return nil, unhandled("jsAssignStmt<unbalanced>", lhs[0])
		}

		r, e := expression(pkg, f, fn, rhs[0])
		if e != nil {
			return nil, e
		}

		exprs = append(exprs, jsast.CreateAssignmentExpression(
			jsast.CreateIdentifier(lname),
			jsast.AssignmentOperator("="),
			r,
		))

		for i, l := range lhs {
			if isUnderscoreVariable(lhs[0]) {
				continue
			}

			x, e := expression(pkg, f, fn, l)
			if e != nil {
				return nil, e
			}

			exprs = append(exprs, jsast.CreateAssignmentExpression(
				x,
				jsast.AssignmentOperator("="),
				jsast.CreateMemberExpression(
					jsast.CreateIdentifier(lname),
					jsast.CreateInt(i),
					true,
				),
			))
		}
	}

	return jsast.CreateExpressionStatement(
		jsast.CreateSequenceExpression(exprs...),
	), nil
}

func jsVariableDecl(pkg *loader.PackageInfo, f *ast.File, fn *ast.FuncDecl, as *ast.AssignStmt) (j jsast.IStatement, err error) {
	var stmts []jsast.VariableDeclarator

	lhs := as.Lhs
	rhs := as.Rhs
	llhs := len(lhs)
	lrhs := len(rhs)

	// ensure we're not in an invalid state
	if llhs != lrhs && lrhs > 1 {
		return nil, fmt.Errorf("invalid golang assignment (AFAIK)")
	}

	// nothing on right side
	if lrhs == 0 {
		for _, lh := range lhs {
			l, ok := lh.(*ast.Ident)
			if !ok {
				return nil, fmt.Errorf("jsVariableDecl<zero>: unhandled type: %s", reflect.TypeOf(lh))
			}

			stmts = append(stmts, jsast.CreateVariableDeclarator(
				jsast.CreateIdentifier(l.Name),
				nil,
			))
		}
	}

	// balanced on both sides
	if llhs == lrhs {
		for i, lh := range lhs {
			l, ok := lh.(*ast.Ident)
			if !ok {
				return nil, fmt.Errorf("jsVariableDecl<balanced>: unhandled type: %s", reflect.TypeOf(lh))
			}

			r, e := expression(pkg, f, fn, rhs[i])
			if e != nil {
				return nil, e
			}

			stmts = append(stmts, jsast.CreateVariableDeclarator(
				jsast.CreateIdentifier(l.Name),
				r,
			))
		}
	}

	// unbalanced
	if llhs != lrhs {
		l, ok := lhs[0].(*ast.Ident)
		if !ok {
			return nil, fmt.Errorf("jsVariableDecl<unbalanced>: unhandled type: %s", reflect.TypeOf(lhs[0]))
		}
		lname := "$" + l.Name

		r, e := expression(pkg, f, fn, rhs[0])
		if e != nil {
			return nil, e
		}

		stmts = append(stmts, jsast.CreateVariableDeclarator(
			jsast.CreateIdentifier(lname),
			r,
		))

		for i, l := range lhs {
			x, ok := l.(*ast.Ident)
			if !ok {
				return nil, fmt.Errorf("jsVariableDecl<unbalanced>: unhandled type: %s", reflect.TypeOf(x))
			}

			stmts = append(stmts, jsast.CreateVariableDeclarator(
				jsast.CreateIdentifier(x.Name),
				jsast.CreateMemberExpression(
					jsast.CreateIdentifier(lname),
					jsast.CreateInt(i),
					true,
				),
			))
		}
	}

	return jsast.CreateVariableDeclaration("var", stmts...), nil
}

func incDecStmt(pkg *loader.PackageInfo, f *ast.File, fn *ast.FuncDecl, n *ast.IncDecStmt) (j jsast.IStatement, err error) {
	var op jsast.UpdateOperator

	x, e := expression(pkg, f, fn, n.X)
	if e != nil {
		return nil, errors.Wrap(e, "incDecStmt")
	}

	switch n.Tok.String() {
	case "++":
		op = jsast.UpdateOperator("++")
	case "--":
		op = jsast.UpdateOperator("--")
	default:
		return nil, unhandled("incDecStmt", n.Tok)
	}

	return jsast.CreateExpressionStatement(
		jsast.CreateUpdateExpression(x, op, false),
	), nil
}

func returnStmt(pkg *loader.PackageInfo, f *ast.File, fn *ast.FuncDecl, n *ast.ReturnStmt) (j jsast.IStatement, err error) {
	// no return values
	if len(n.Results) == 0 {
		return jsast.CreateReturnStatement(nil), nil
	}

	var args []jsast.IExpression
	for _, arg := range n.Results {
		a, e := expression(pkg, f, fn, arg)
		if e != nil {
			return nil, e
		}
		args = append(args, a)
	}

	// return an array
	if len(n.Results) > 1 {
		return jsast.CreateReturnStatement(jsast.CreateArrayExpression(args...)), nil
	}

	// return the value by itself
	return jsast.CreateReturnStatement(args[0]), nil
}

func callExpression(pkg *loader.PackageInfo, f *ast.File, fn *ast.FuncDecl, n *ast.CallExpr) (j jsast.IExpression, err error) {
	// calleeSrc, e := expressionToString(expr.Fun)
	// if e != nil {
	// 	return j, e
	// }

	// // TODO: make sure jsast.Raw points to golly/js
	// if calleeSrc == "jsast.Raw" && len(expr.Args) >= 1 {
	// 	argSrc, e := expressionToString(expr.Args[0])
	// 	if e != nil {
	// 		return nil, e
	// 	}
	// 	return jsast.Parse(argSrc)
	// }

	// create an expression for built-in golang functions like append
	// TODO: better name for what this does
	if expr, e := checkBuiltin(pkg, f, fn, n); expr != nil || e != nil {
		return expr, e
	}

	callee, e := expression(pkg, f, fn, n.Fun)
	if e != nil {
		return j, e
	}

	var args []jsast.IExpression
	for _, arg := range n.Args {
		v, e := expression(pkg, f, fn, arg)
		if e != nil {
			return j, e
		}
		args = append(args, v)
	}

	return jsast.CreateCallExpression(
		callee,
		args,
	), nil
}

func expression(pkg *loader.PackageInfo, f *ast.File, fn *ast.FuncDecl, expr ast.Expr) (j jsast.IExpression, err error) {
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
	case *ast.ArrayType:
		return arrayType(pkg, f, fn, t)
	case *ast.ChanType:
		return chanType(pkg, f, fn, t)
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

func funcLit(pkg *loader.PackageInfo, f *ast.File, fn *ast.FuncDecl, n *ast.FuncLit) (j jsast.IExpression, err error) {

	// log.Infof("anonymous func %+v", pkg.(n.Type))

	// build argument list
	// var args
	var params []jsast.IPattern
	if n.Type != nil && n.Type.Params != nil {
		fields := n.Type.Params.List
		for _, field := range fields {
			// names because: (a, b string, c int) = [[a, b], c]
			for _, name := range field.Names {
				id := jsast.CreateIdentifier(name.Name)
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

	return jsast.CreateFunctionExpression(nil, params, jsast.CreateFunctionBody(body...)), nil
}

// binary expressions in Go can be either:
//    Binaryexpression || LogicalExpression
// in jsast.
func binaryExpression(pkg *loader.PackageInfo, f *ast.File, fn *ast.FuncDecl, b *ast.BinaryExpr) (j jsast.IExpression, err error) {
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
		return jsast.CreateLogicalExpression(x, jsast.LogicalOperator(op), y), nil
	// TODO: prune. should be only values that are possible in Go
	case "==", "!=", "===", "!==",
		"<", "<=", ">", ">=", "<<",
		">>", ">>>", "+", "-", "*",
		"/", "%", "|", "^", "&",
		"in", "instanceof":
		return jsast.CreateBinaryExpression(x, jsast.BinaryOperator(op), y), nil
	default:
		return nil, unhandled("binaryExpression<unknown op>", op)
	}
}

func identifier(pkg *loader.PackageInfo, f *ast.File, fn *ast.FuncDecl, n *ast.Ident) (j jsast.IExpression, err error) {
	// a, e := getObjectType(n)
	// if e != nil {
	// 	return nil, e
	// }
	obj := pkg.ObjectOf(n)
	if obj == nil {
		log.Warnf("no object of")
	} else {
		log.Infof(obj.String())
	}
	// fmt.Println(n.Obj)
	// fmt.Printf("%+v\n", a)

	// TODO: lookup table
	switch n.Name {
	case "nil":
		return jsast.CreateNull(), nil
	case "println":
		return jsast.CreateMemberExpression(
			jsast.CreateIdentifier("console"),
			jsast.CreateIdentifier("log"),
			false,
		), nil
	default:
		return jsast.CreateIdentifier(n.Name), nil
	}
}

func starExpr(pkg *loader.PackageInfo, f *ast.File, fn *ast.FuncDecl, n *ast.StarExpr) (j jsast.IExpression, err error) {
	// for now, we're ignoring the pointer
	x, e := expression(pkg, f, fn, n.X)
	if e != nil {
		return nil, e
	}

	return x, nil
}

func unaryExpr(pkg *loader.PackageInfo, f *ast.File, fn *ast.FuncDecl, n *ast.UnaryExpr) (j jsast.IExpression, err error) {
	// for now, we're ignoring the pointer
	x, e := expression(pkg, f, fn, n.X)
	if e != nil {
		return nil, e
	}

	switch n.Op.String() {
	case "<-":
		return jsast.CreateAwaitExpression(
			jsast.CreateCallExpression(
				jsast.CreateMemberExpression(
					x,
					jsast.CreateIdentifier("Recv"),
					false,
				),
				[]jsast.IExpression{},
			),
		), nil
	default:
		return x, nil
	}
}

func basiclit(pkg *loader.PackageInfo, f *ast.File, fn *ast.FuncDecl, lit *ast.BasicLit) (j jsast.IExpression, err error) {
	return jsast.CreateLiteral(lit.Value), nil
}

func compositeLiteral(pkg *loader.PackageInfo, f *ast.File, fn *ast.FuncDecl, n *ast.CompositeLit) (j jsast.IExpression, err error) {
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

func jsObjectExpression(pkg *loader.PackageInfo, f *ast.File, fn *ast.FuncDecl, n *ast.CompositeLit) (j jsast.ObjectExpression, err error) {
	var props []jsast.Property

	for _, elt := range n.Elts {
		var prop jsast.Property
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

	return jsast.CreateObjectExpression(props), nil
}

func jsNewFunction(pkg *loader.PackageInfo, f *ast.File, fn *ast.FuncDecl, n *ast.CompositeLit) (j jsast.IExpression, err error) {
	var fnname jsast.IExpression
	var props []jsast.Property

	switch t := n.Type.(type) {
	// e.g. Document{}
	case *ast.Ident:
		fnname = jsast.CreateIdentifier(t.Name)
	// e.g. dom.Document{}
	case *ast.SelectorExpr:
		// this will convert `dom.Document{}` into
		// var document = window.document
		// where the js tag is `js:"document,global"`
		objectof := pkg.ObjectOf(t.Sel)
		if objectof != nil && aliases[objectof.String()] != nil && aliases[objectof.String()].HasOption("global") {
			return jsast.CreateMemberExpression(
				jsast.CreateIdentifier("window"),
				jsast.CreateIdentifier(aliases[objectof.String()].Name),
				false,
			), nil
		}

		sel, e := selectorExpr(pkg, f, fn, t)
		if e != nil {
			return j, e
		}
		fnname = sel
	default:
		return j, unhandled("jsNewFunction<name>", n.Type)
	}

	for _, elt := range n.Elts {
		var prop jsast.Property
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

	return jsast.CreateNewExpression(
		fnname,
		[]jsast.IExpression{jsast.CreateObjectExpression(props)},
	), nil
}

func jsArrayExpression(pkg *loader.PackageInfo, f *ast.File, fn *ast.FuncDecl, n *ast.CompositeLit) (j jsast.ArrayExpression, err error) {
	var elements []jsast.IExpression

	for _, elt := range n.Elts {
		el, e := expression(pkg, f, fn, elt)
		if e != nil {
			return j, e
		}
		elements = append(elements, el)
	}

	return jsast.CreateArrayExpression(elements...), nil
}

func keyValueExpr(pkg *loader.PackageInfo, f *ast.File, fn *ast.FuncDecl, n *ast.KeyValueExpr) (j jsast.Property, err error) {
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

	return jsast.CreateProperty(key, value, "init"), nil
}

func selectorExpr(pkg *loader.PackageInfo, f *ast.File, fn *ast.FuncDecl, n *ast.SelectorExpr) (j jsast.MemberExpression, err error) {
	var x jsast.IExpression
	var s jsast.IExpression

	// first check if we have an alias already
	// typeof := pkg.TypeOf(n.X)
	// if typeof != nil && aliases[typeof.String()] != nil {
	// 	x = jsast.CreateIdentifier(aliases[typeof.String()].Name)
	// }

	// (user.phone).number
	if x == nil {
		ex, e := expression(pkg, f, fn, n.X)
		if e != nil {
			return j, e
		}
		x = ex
	}

	// first check if we have an alias already
	objectof := pkg.ObjectOf(n.Sel)
	if objectof != nil && aliases[objectof.String()] != nil {
		s = jsast.CreateIdentifier(aliases[objectof.String()].Name)
	}

	// user.phone.(number)
	if s == nil {
		se, e := identifier(pkg, f, fn, n.Sel)
		if e != nil {
			return j, e
		}
		s = se
	}

	member := jsast.CreateMemberExpression(x, s, false)
	return member, nil
}

func indexExpr(pkg *loader.PackageInfo, f *ast.File, fn *ast.FuncDecl, n *ast.IndexExpr) (j jsast.MemberExpression, err error) {
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

	return jsast.CreateMemberExpression(x, i, true), nil
}

func unhandled(fn string, n interface{}) error {
	return fmt.Errorf("%s in %s() is not implemented yet", reflect.TypeOf(n), fn)
}

func getCommentTag(n *ast.CommentGroup) (*structtag.Tag, error) {
	if n == nil {
		return nil, nil
	}

	comments := n.List
	for _, comment := range comments {
		if !strings.Contains(comment.Text, "js:") {
			continue
		}

		tags, err := structtag.Parse(comment.Text[2:])
		if err != nil {
			return nil, err
		}

		jstag, err := tags.Get("js")
		if err != nil {
			return nil, err
		}

		return jstag, nil
	}

	return nil, nil
}

func checkBuiltin(pkg *loader.PackageInfo, f *ast.File, fn *ast.FuncDecl, n *ast.CallExpr) (jsast.IExpression, error) {
	id, ok := n.Fun.(*ast.Ident)
	if !ok {
		return nil, nil
	}

	switch id.Name {
	case "append":
		return builtinAppend(pkg, f, fn, n.Args)
	case "len":
	case "copy":
	case "make":
		return expression(pkg, f, fn, n.Args[0])
	}

	return nil, nil
}

func builtinAppend(pkg *loader.PackageInfo, f *ast.File, fn *ast.FuncDecl, ns []ast.Expr) (jsast.IExpression, error) {
	var els []jsast.IExpression
	for _, n := range ns {
		x, e := expression(pkg, f, fn, n)
		if e != nil {
			return nil, e
		}
		els = append(els, x)
	}

	if len(els) == 1 {
		return els[0], nil
	}

	return jsast.CreateCallExpression(
		jsast.CreateMemberExpression(
			els[0],
			jsast.CreateIdentifier("concat"),
			false,
		),
		els[1:],
	), nil
}

func isUnderscoreVariable(expr ast.Expr) bool {
	switch t := expr.(type) {
	case *ast.Ident:
		if t.Name == "_" {
			return true
		}
	}
	return false
}

func arrayType(pkg *loader.PackageInfo, f *ast.File, fn *ast.FuncDecl, n *ast.ArrayType) (jsast.IExpression, error) {
	return jsast.CreateArrayExpression(), nil
}

func chanType(pkg *loader.PackageInfo, f *ast.File, fn *ast.FuncDecl, n *ast.ChanType) (jsast.IExpression, error) {
	return jsast.CreateNewExpression(
		jsast.CreateIdentifier("Channel"),
		[]jsast.IExpression{},
	), nil
}

// get the identifier when possible
func getObjectType(expr ast.Expr) (ast.Expr, error) {
	var id *ast.Ident
	switch t := expr.(type) {
	case *ast.UnaryExpr:
		return getObjectType(t.X)
	case *ast.SelectorExpr:
		return getObjectType(t.Sel)
	case *ast.Ident:
		id = t
	default:
		return nil, unhandled("getObjectType<expr>", expr)
	}

	obj := id.Obj
	if obj == nil {
		return nil, fmt.Errorf("getObjectType object nil")
	}

	switch t := obj.Decl.(type) {
	case *ast.ValueSpec:
		if t.Type == nil {
			return nil, fmt.Errorf("getObjectType type nil")
		}
		return t.Type, nil
	default:
		return nil, unhandled("getObjectType<decl>", expr)
	}
}
