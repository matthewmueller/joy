package translator

import (
	"fmt"
	"go/ast"
	"path"
	"reflect"
	"strings"

	"github.com/apex/log"
	"github.com/fatih/structtag"
	"github.com/matthewmueller/golly/golang/scope"
	"github.com/matthewmueller/golly/jsast"
	"github.com/matthewmueller/golly/types"
	"github.com/pkg/errors"
)

// context struct
type context struct {
	declaration *types.Declaration
}

// Result of the translation
type Result struct {
	Node         jsast.INode
	Exported     bool
	Dependencies []string
}

// Translate the declaration
func Translate(decl *types.Declaration) (node jsast.INode, err error) {
	sp := scope.New(decl.Node)

	ctx := &context{
		declaration: decl,
	}

	switch t := decl.Node.(type) {
	case *ast.FuncDecl:
		node, err = funcDecl(ctx, sp, t)
	case *ast.GenDecl:
		node, err = genDecl(ctx, sp, t)
	default:
		return nil, unhandled("Translate", decl)
	}
	if err != nil {
		return nil, err
	}

	return node, nil
	// return &Result{
	// 	Node:         node,
	// 	Dependencies: unique(ctx.dependencies),
	// 	Exported:     ctx.exported,
	// }, nil
}

func funcDecl(ctx *context, sp *scope.Scope, n *ast.FuncDecl) (jsast.IStatement, error) {
	isAsync := ctx.declaration.Async

	// e.g. func hi()
	if n.Body == nil {
		return jsast.CreateEmptyStatement(), nil
	}

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
		jsStmt, e := funcStatement(ctx, sp, stmt)
		if e != nil {
			return nil, e
		}
		body = append(body, jsStmt)
	}

	fnname := jsast.CreateIdentifier(n.Name.Name)

	// async function
	if n.Recv == nil && isAsync {
		return jsast.CreateAsyncFunction(
			&fnname,
			params,
			jsast.CreateFunctionBody(body...),
		), nil
	}

	// regular function
	if n.Recv == nil && !isAsync {
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
	x, e := expression(ctx, sp, recv.Type)
	if e != nil {
		return nil, e
	}

	if len(recv.Names) > 1 {
		return nil, fmt.Errorf("function<recv>: only expected 1 func receiver name but got %d", len(recv.Names))
	}

	// Links the function receiver to "this",
	// Placing it on top of the function body
	// e.g. var d = this
	//
	// TODO: be smarter here and rename the inner body variables to "this"
	if len(recv.Names) == 1 {
		body = append([]interface{}{jsast.CreateVariableDeclaration(
			"var",
			jsast.CreateVariableDeclarator(
				jsast.CreateIdentifier(recv.Names[0].Name),
				jsast.CreateThisExpression(),
			),
		)}, body...)
	}

	var fn jsast.FunctionExpression
	if isAsync {
		fn = jsast.CreateAsyncFunctionExpression(
			&fnname,
			params,
			jsast.CreateFunctionBody(body...),
		)
	} else {
		fn = jsast.CreateFunctionExpression(
			&fnname,
			params,
			jsast.CreateFunctionBody(body...),
		)
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
			fn,
		),
	), nil
}

func genDecl(ctx *context, sp *scope.Scope, n *ast.GenDecl) (j jsast.IStatement, err error) {
	switch n.Tok.String() {
	case "import":
		return importSpec(ctx, sp, n)
	case "type":
		return typeSpec(ctx, sp, n)
	case "var":
		return varSpec(ctx, sp, n)
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

func typeSpec(ctx *context, sp *scope.Scope, n *ast.GenDecl) (j jsast.IStatement, err error) {
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
	// objectof := ctx.info.ObjectOf(s.Name)
	// typeof := ctx.info.TypeOf(s.Name)
	// if tag != nil && objectof != nil {
	// 	ctx.aliases[objectof.String()] = tag
	// 	ctx.aliases[typeof.String()] = tag
	// 	// TODO: not sure if this is a good idea or not
	// 	// but it's to handle pointer receivers in 1 spot
	// 	ctx.aliases["*"+typeof.String()] = tag
	// }

	var ivars []interface{}
	for _, field := range st.Fields.List {
		for _, name := range field.Names {
			// objectof := ctx.info.ObjectOf(name)
			// // TODO: this doesn't need to be parsed for each field name
			// // though I'm not sure how you can have multiple field names
			// // per struct tag
			// if field.Tag != nil {
			// 	tags, err := structtag.Parse(field.Tag.Value[1 : len(field.Tag.Value)-1])
			// 	if err != nil {
			// 		return j, err
			// 	}
			// 	if tag, err := tags.Get("js"); err == nil {
			// 		ctx.aliases[objectof.String()] = tag
			// 	}
			// }

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

func importSpec(ctx *context, sp *scope.Scope, n *ast.GenDecl) (j jsast.IStatement, err error) {
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

func varSpec(ctx *context, sp *scope.Scope, n *ast.GenDecl) (j jsast.IStatement, err error) {
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
					r, e := expression(ctx, sp, t.Values[i])
					if e != nil {
						return j, e
					}
					rh = r
				} else {
					r, e := expression(ctx, sp, t.Type)
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

// func mainFunc(pkg *loader.PackageInfo,  f *ast.File) (j jsast.IStatement, err error) {
// 	// e.g. func main()
// 	if fn.Body == nil {
// 		return jsast.CreateEmptyStatement(), nil
// 	}

// 	// TODO: []IStatement{} instead of interface{}
// 	var body []interface{}
// 	for _, stmt := range fn.Body.List {
// 		jsStmt, e := funcStatement(ctx, sp, stmt)
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

func funcStatement(ctx *context, sp *scope.Scope, n ast.Stmt) (j jsast.IStatement, err error) {
	switch t := n.(type) {
	case *ast.ExprStmt:
		return exprStatement(ctx, sp, t)
	case *ast.IfStmt:
		return ifStmt(ctx, sp, t)
	case *ast.AssignStmt:
		return assignStatement(ctx, sp, t)
	case *ast.ReturnStmt:
		return returnStmt(ctx, sp, t)
	case *ast.RangeStmt:
		return rangeStmt(ctx, sp, t)
	case *ast.ForStmt:
		return forStmt(ctx, sp, t)
	case *ast.DeclStmt:
		return declStmt(ctx, sp, t)
	case *ast.GoStmt:
		return goStmt(ctx, sp, t)
	case *ast.SendStmt:
		return sendStmt(ctx, sp, t)
	default:
		return nil, unhandled("funcStatement", n)
	}
}

func goStmt(ctx *context, sp *scope.Scope, n *ast.GoStmt) (j jsast.IStatement, err error) {
	// isAsync := ctx.declaration.Async

	var args []jsast.IExpression
	for _, arg := range n.Call.Args {
		x, e := expression(ctx, sp, arg)
		if e != nil {
			return nil, e
		}
		args = append(args, x)
	}

	switch t := n.Call.Fun.(type) {
	case *ast.FuncLit:
		// build argument list
		// var args
		var params []jsast.IPattern
		if t.Type != nil && t.Type.Params != nil {
			fields := t.Type.Params.List
			for _, field := range fields {
				// names because: (a, b string, c int) = [[a, b], c]
				for _, name := range field.Names {
					id := jsast.CreateIdentifier(name.Name)
					params = append(params, id)
				}
			}
		}

		var body []interface{}
		for _, stmt := range t.Body.List {
			s, e := statement(ctx, sp, stmt)
			if e != nil {
				return nil, e
			}
			body = append(body, s)
		}

		return jsast.CreateExpressionStatement(
			jsast.CreateCallExpression(
				jsast.CreateAsyncFunctionExpression(
					nil,
					params,
					jsast.CreateFunctionBody(body...),
				),
				args,
			),
		), nil

	// case *ast.FuncDecl:
	// 	log.Infof("go func decl!")
	default:
		return nil, unhandled("goStmt", n.Call)
	}

	// c, e := callExpression(ctx, sp, n.Call)
	// if e != nil {
	// 	return nil, e
	// }

	// ast.
	// log.Infof("Scope %+v", f.Scope.Lookup(n))

	// return jsast.CreateExpressionStatement(c), nil
	// return nil, nil
}

func sendStmt(ctx *context, sp *scope.Scope, n *ast.SendStmt) (j jsast.IStatement, err error) {
	ch, e := expression(ctx, sp, n.Chan)
	if e != nil {
		return nil, e
	}

	val, e := expression(ctx, sp, n.Value)
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

func rangeStmt(ctx *context, sp *scope.Scope, n *ast.RangeStmt) (j jsast.IStatement, err error) {
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

	rh, e := expression(ctx, sp, asn.Rhs[0])
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
		v, e := statement(ctx, sp, stmt)
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

	// a, e := assignStatement(ctx, sp, asn)
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

	// init, e := expression(ctx, sp, n.Key)
	// if e != nil {
	// 	return nil, errors.Wrap(e, "forStmt")
	// }
	// log.Infof("init %s", init)
	// cond, e := expression(ctx, sp, n.Cond)
	// if e != nil {
	// 	return nil, errors.Wrap(e, "forStmt")
	// }

	// post, e := statement(ctx, sp, n.Post)
	// if e != nil {
	// 	return nil, errors.Wrap(e, "forStmt")
	// }

	// body, e := blockStmt(ctx, sp, n.Body)
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

func declStmt(ctx *context, sp *scope.Scope, n *ast.DeclStmt) (j jsast.IStatement, err error) {
	switch t := n.Decl.(type) {
	case *ast.GenDecl:
		return genDecl(ctx, sp, t)
	default:
		return nil, unhandled("declStmt", n)
	}
}

func exprStatement(ctx *context, sp *scope.Scope, expr *ast.ExprStmt) (j jsast.IExpressionStatement, err error) {
	switch t := expr.X.(type) {
	case *ast.CallExpr:
		x, e := callExpression(ctx, sp, t)
		if e != nil {
			return nil, e
		}
		return jsast.CreateExpressionStatement(x), nil
	default:
		return nil, fmt.Errorf("exprStatement: unhandled type: %s", reflect.TypeOf(expr))
	}
}

func ifStmt(ctx *context, sp *scope.Scope, n *ast.IfStmt) (j jsast.IStatement, err error) {
	var e error

	// condition: if [(...)] { ... } else { ... }
	var test jsast.IExpression
	switch t := n.Cond.(type) {
	case *ast.BinaryExpr:
		test, e = binaryExpression(ctx, sp, t)
	case *ast.Ident:
		test, e = identifier(ctx, sp, t)
	default:
		return nil, unhandled("ifStmt<cond>", n.Cond)
	}
	if e != nil {
		return nil, e
	}

	// body : if (...) [{ ... }] else { ... }
	body, e := blockStmt(ctx, sp, n.Body)
	if e != nil {
		return nil, e
	}

	// else: if (...) { ... } else [{ ... }]
	elseBlock := n.Else
	var alt jsast.IStatement
	switch t := elseBlock.(type) {
	case *ast.BlockStmt:
		alt, e = blockStmt(ctx, sp, t)
	case *ast.ExprStmt:
		alt, e = exprStatement(ctx, sp, t)
	case *ast.IfStmt:
		alt, e = ifStmt(ctx, sp, t)
	case *ast.AssignStmt:
		alt, e = assignStatement(ctx, sp, t)
	case *ast.ReturnStmt:
		alt, e = returnStmt(ctx, sp, t)
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

func branchStmt(ctx *context, sp *scope.Scope, n *ast.BranchStmt) (j jsast.IStatement, err error) {
	switch n.Tok.String() {
	case "break":
		return jsast.CreateBreakStatement(nil), nil
	default:
		return nil, fmt.Errorf("unhandled branchStmt: %s", n.Tok.String())
	}
}

func forStmt(ctx *context, sp *scope.Scope, n *ast.ForStmt) (j jsast.IStatement, err error) {

	init, e := statement(ctx, sp, n.Init)
	if e != nil {
		return nil, errors.Wrap(e, "forStmt")
	}

	cond, e := expression(ctx, sp, n.Cond)
	if e != nil {
		return nil, errors.Wrap(e, "forStmt")
	}

	post, e := statement(ctx, sp, n.Post)
	if e != nil {
		return nil, errors.Wrap(e, "forStmt")
	}

	body, e := blockStmt(ctx, sp, n.Body)
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

func statement(ctx *context, sp *scope.Scope, n ast.Stmt) (j jsast.IStatement, err error) {
	switch t := n.(type) {
	case nil:
		return nil, nil
	case *ast.AssignStmt:
		return assignStatement(ctx, sp, t)
	case *ast.IncDecStmt:
		return incDecStmt(ctx, sp, t)
	case *ast.ExprStmt:
		return exprStatement(ctx, sp, t)
	case *ast.IfStmt:
		return ifStmt(ctx, sp, t)
	case *ast.BranchStmt:
		return branchStmt(ctx, sp, t)
	case *ast.ReturnStmt:
		return returnStmt(ctx, sp, t)
	case *ast.SendStmt:
		return sendStmt(ctx, sp, t)
	default:
		return nil, unhandled("statement", n)
	}
}

func blockStmt(ctx *context, sp *scope.Scope, n *ast.BlockStmt) (j jsast.IBlockStatement, err error) {
	var stmts []jsast.IStatement

	for _, stmt := range n.List {
		v, e := statement(ctx, sp, stmt)
		if e != nil {
			return nil, errors.Wrap(e, "blockStmt")
		}
		stmts = append(stmts, v)
	}

	return jsast.CreateBlockStatement(stmts...), nil
}

func assignStatement(ctx *context, sp *scope.Scope, as *ast.AssignStmt) (j jsast.IStatement, err error) {
	// TODO: these are separate, but very similar functions.
	// the reason they're separate is because the JS AST's are
	// different. It'd be good to come up with a way to consolidate
	// this logic though because this variable conversion is a bit tricky
	switch as.Tok.String() {
	case "=":
		return jsAssignStmt(ctx, sp, as)
	case ":=":
		return jsVariableDecl(ctx, sp, as)
	default:
		return nil, fmt.Errorf("unhandled assignStatement<tok>: %s", as.Tok.String())
	}
}

func jsAssignStmt(ctx *context, sp *scope.Scope, as *ast.AssignStmt) (j jsast.IStatement, err error) {
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
			l, e := expression(ctx, sp, lh)
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

			l, e := expression(ctx, sp, lh)
			if e != nil {
				return nil, e
			}

			r, e := expression(ctx, sp, rhs[i])
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

		r, e := expression(ctx, sp, rhs[0])
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

			x, e := expression(ctx, sp, l)
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

func jsVariableDecl(ctx *context, sp *scope.Scope, as *ast.AssignStmt) (j jsast.IStatement, err error) {
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

			r, e := expression(ctx, sp, rhs[i])
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

		r, e := expression(ctx, sp, rhs[0])
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

func incDecStmt(ctx *context, sp *scope.Scope, n *ast.IncDecStmt) (j jsast.IStatement, err error) {
	var op jsast.UpdateOperator

	x, e := expression(ctx, sp, n.X)
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

func returnStmt(ctx *context, sp *scope.Scope, n *ast.ReturnStmt) (j jsast.IStatement, err error) {
	// no return values
	if len(n.Results) == 0 {
		return jsast.CreateReturnStatement(nil), nil
	}

	var args []jsast.IExpression
	for _, arg := range n.Results {
		a, e := expression(ctx, sp, arg)
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

func callExpression(ctx *context, sp *scope.Scope, n *ast.CallExpr) (j jsast.IExpression, err error) {
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
	if expr, e := checkBuiltin(ctx, sp, n); expr != nil || e != nil {
		return expr, e
	}

	callee, e := expression(ctx, sp, n.Fun)
	if e != nil {
		return j, e
	}

	var args []jsast.IExpression
	for _, arg := range n.Args {
		v, e := expression(ctx, sp, arg)
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

func expression(ctx *context, sp *scope.Scope, expr ast.Expr) (j jsast.IExpression, err error) {
	switch t := expr.(type) {
	case *ast.Ident:
		return identifier(ctx, sp, t)
	case *ast.BasicLit:
		return basiclit(ctx, sp, t)
	case *ast.CallExpr:
		return callExpression(ctx, sp, t)
	case *ast.BinaryExpr:
		return binaryExpression(ctx, sp, t)
	case *ast.CompositeLit:
		return compositeLiteral(ctx, sp, t)
	case *ast.SelectorExpr:
		return selectorExpr(ctx, sp, t)
	case *ast.IndexExpr:
		return indexExpr(ctx, sp, t)
	case *ast.StarExpr:
		return starExpr(ctx, sp, t)
	case *ast.UnaryExpr:
		return unaryExpr(ctx, sp, t)
	case *ast.FuncLit:
		return funcLit(ctx, sp, t)
	case *ast.ArrayType:
		return arrayType(ctx, sp, t)
	case *ast.ChanType:
		return chanType(ctx, sp, t)
	// case *ast.KeyValueExpr:
	// 	return keyValueExpr(ctx, sp, t)
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

func funcLit(ctx *context, sp *scope.Scope, n *ast.FuncLit) (j jsast.IExpression, err error) {

	// log.Infof("anonymous func %+v", ctx.info.(n.Type))

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
		jsStmt, e := funcStatement(ctx, sp, stmt)
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
func binaryExpression(ctx *context, sp *scope.Scope, b *ast.BinaryExpr) (j jsast.IExpression, err error) {
	x, e := expression(ctx, sp, b.X)
	if e != nil {
		return nil, e
	}
	y, e := expression(ctx, sp, b.Y)
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

func identifier(ctx *context, sp *scope.Scope, n *ast.Ident) (j jsast.IExpression, err error) {
	// check if we depend on any other
	// declarations with this identifier
	// obj := ctx.info.ObjectOf(n)
	// if obj != nil && obj.Pkg() != nil {
	// 	ctx.dependencies = append(ctx.dependencies, obj.String())
	// }

	// switch obj
	// log.Infof("objcect type %s", obj.Type().String())

	// switch t := obj.Type().String() {
	// case *types.Package:
	// 	log.Warnf("func decl!!!!!!!!")
	// // case *ast.GenDecl:
	// // 	log.Warnf("func decl!!!!!!!!")
	// default:
	// 	log.Warnf("neither %s", reflect.TypeOf(obj))
	// }
	// if obj == nil {
	// 	log.Warnf("no object of")
	// } else {
	// 	log.Infof("identifier id: %s", obj.String())
	// }
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

func starExpr(ctx *context, sp *scope.Scope, n *ast.StarExpr) (j jsast.IExpression, err error) {
	// for now, we're ignoring the pointer
	x, e := expression(ctx, sp, n.X)
	if e != nil {
		return nil, e
	}

	return x, nil
}

func unaryExpr(ctx *context, sp *scope.Scope, n *ast.UnaryExpr) (j jsast.IExpression, err error) {
	// for now, we're ignoring the pointer
	x, e := expression(ctx, sp, n.X)
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

func basiclit(ctx *context, sp *scope.Scope, lit *ast.BasicLit) (j jsast.IExpression, err error) {
	return jsast.CreateLiteral(lit.Value), nil
}

func compositeLiteral(ctx *context, sp *scope.Scope, n *ast.CompositeLit) (j jsast.IExpression, err error) {
	switch n.Type.(type) {
	case *ast.Ident, *ast.SelectorExpr:
		return jsNewFunction(ctx, sp, n)
	case *ast.ArrayType:
		return jsArrayExpression(ctx, sp, n)
	case *ast.MapType:
		return jsObjectExpression(ctx, sp, n)
	default:
		return nil, unhandled("compositeLiteral<type>", n.Type)
	}
}

func jsObjectExpression(ctx *context, sp *scope.Scope, n *ast.CompositeLit) (j jsast.ObjectExpression, err error) {
	var props []jsast.Property

	for _, elt := range n.Elts {
		var prop jsast.Property
		var e error

		switch t := elt.(type) {
		case *ast.KeyValueExpr:
			prop, e = keyValueExpr(ctx, sp, t)
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

func jsNewFunction(ctx *context, sp *scope.Scope, n *ast.CompositeLit) (j jsast.IExpression, err error) {
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
		// objectof := ctx.info.ObjectOf(t.Sel)
		// if objectof != nil && ctx.aliases[objectof.String()] != nil && ctx.aliases[objectof.String()].HasOption("global") {
		// 	return jsast.CreateMemberExpression(
		// 		jsast.CreateIdentifier("window"),
		// 		jsast.CreateIdentifier(ctx.aliases[objectof.String()].Name),
		// 		false,
		// 	), nil
		// }

		sel, e := selectorExpr(ctx, sp, t)
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
			prop, e = keyValueExpr(ctx, sp, t)
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

func jsArrayExpression(ctx *context, sp *scope.Scope, n *ast.CompositeLit) (j jsast.ArrayExpression, err error) {
	var elements []jsast.IExpression

	for _, elt := range n.Elts {
		el, e := expression(ctx, sp, elt)
		if e != nil {
			return j, e
		}
		elements = append(elements, el)
	}

	return jsast.CreateArrayExpression(elements...), nil
}

func keyValueExpr(ctx *context, sp *scope.Scope, n *ast.KeyValueExpr) (j jsast.Property, err error) {
	// get the key
	key, e := expression(ctx, sp, n.Key)
	if e != nil {
		return j, e
	}

	// get the value
	value, e := expression(ctx, sp, n.Value)
	if e != nil {
		return j, e
	}

	return jsast.CreateProperty(key, value, "init"), nil
}

func selectorExpr(ctx *context, sp *scope.Scope, n *ast.SelectorExpr) (j jsast.MemberExpression, err error) {
	var x jsast.IExpression
	var s jsast.IExpression

	// first check if we have an alias already
	// typeof := ctx.info.TypeOf(n.X)
	// if typeof != nil && ctx.aliases[typeof.String()] != nil {
	// 	x = jsast.CreateIdentifier(ctx.aliases[typeof.String()].Name)
	// }

	// (user.phone).number
	if x == nil {
		ex, e := expression(ctx, sp, n.X)
		if e != nil {
			return j, e
		}
		x = ex
	}

	// first check if we have an alias already
	// objectof := ctx.info.ObjectOf(n.Sel)
	// if objectof != nil && ctx.aliases[objectof.String()] != nil {
	// 	s = jsast.CreateIdentifier(ctx.aliases[objectof.String()].Name)
	// }

	// user.phone.(number)
	if s == nil {
		se, e := identifier(ctx, sp, n.Sel)
		if e != nil {
			return j, e
		}
		s = se
	}

	member := jsast.CreateMemberExpression(x, s, false)
	return member, nil
}

func indexExpr(ctx *context, sp *scope.Scope, n *ast.IndexExpr) (j jsast.MemberExpression, err error) {
	// (i)[0]
	x, e := expression(ctx, sp, n.X)
	if e != nil {
		return j, e
	}

	// i([0])
	i, e := expression(ctx, sp, n.Index)
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

func checkBuiltin(ctx *context, sp *scope.Scope, n *ast.CallExpr) (jsast.IExpression, error) {
	id, ok := n.Fun.(*ast.Ident)
	if !ok {
		return nil, nil
	}

	switch id.Name {
	case "append":
		return builtinAppend(ctx, sp, n.Args)
	case "len":
	case "copy":
	case "make":
		return expression(ctx, sp, n.Args[0])
	}

	return nil, nil
}

func builtinAppend(ctx *context, sp *scope.Scope, ns []ast.Expr) (jsast.IExpression, error) {
	var els []jsast.IExpression
	for _, n := range ns {
		x, e := expression(ctx, sp, n)
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

func arrayType(ctx *context, sp *scope.Scope, n *ast.ArrayType) (jsast.IExpression, error) {
	return jsast.CreateArrayExpression(), nil
}

func chanType(ctx *context, sp *scope.Scope, n *ast.ChanType) (jsast.IExpression, error) {
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

func unique(s []string) []string {
	unique := make(map[string]bool, len(s))
	us := make([]string, len(unique))
	for _, elem := range s {
		if len(elem) != 0 {
			if !unique[elem] {
				us = append(us, elem)
				unique[elem] = true
			}
		}
	}
	return us
}
