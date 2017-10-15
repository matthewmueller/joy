package index

import (
	"fmt"
	"go/ast"

	"golang.org/x/tools/go/loader"

	"github.com/fatih/structtag"
	"github.com/matthewmueller/golly/compiler/translator"
	"github.com/matthewmueller/golly/compiler/types"
	"github.com/matthewmueller/golly/compiler/util"
	"github.com/matthewmueller/golly/jsast"
)

// Function creates a new function declaration
func (i *Index) Function(info *loader.PackageInfo, n *ast.FuncDecl) error {
	id := info.ObjectOf(n.Name).String()

	fn := &function{
		info:  info,
		index: i,
		id:    id,
		node:  n,
	}

	// add to the index
	i.decls[id] = fn

	return nil
}

var _ types.Declaration = (*function)(nil)

// function declaration type
type function struct {
	info  *loader.PackageInfo
	index *Index
	id    string
	node  *ast.FuncDecl

	name     string
	omit     bool
	exported bool
	async    bool

	// Just names for now because
	// we only need the names right
	// now for js.Rewrite.
	params []string

	// This is nonnil when there is a comment above
	// the declaration
	// e.g. // js:"..."
	tag *structtag.Tag

	// Rewrite *Rewrite
	dependencies []types.Declaration
}

func (f *function) process() {

}

// ID string
func (f *function) ID() string {
	return f.id
}

// Name of the declaration
func (f *function) Name() string {
	return f.node.Name.Name
}

// Exported fn
// TODO
func (f *function) Exported() bool {
	return true
}

// Path to the declaration
func (f *function) Path() string {
	return f.info.Pkg.Path()
}

// Dependencies fn
func (f *function) Dependencies() (deps []types.Declaration, err error) {
	ast.Inspect(f.node, func(n ast.Node) bool {
		switch t := n.(type) {
		case *ast.FuncDecl:
			tag, e := util.JSTag(t.Doc)
			if e != nil {
				err = e
				return false
			} else if tag != nil {
				f.tag = tag
				if tag.HasOption("global") {
					f.exported = false
				}
			}
		}
		return true
	})
	if err != nil {
		return deps, err
	}

	return f.dependencies, nil
}

// Translate fn
func (f *function) Translate() (jsast.IStatement, error) {
	isAsync := f.async
	n := f.node

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
		stmt, e := translator.Statement(f.index, stmt)
		if e != nil {
			return nil, e
		}
		body = append(body, stmt)
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
	recvDecl := ctx.index.FindByNode(ctx.info, recv.Type)
	if recvDecl != nil {
		// remove prototypes where the class is global
		if recvDecl.JSTag != nil && recvDecl.JSTag.HasOption("global") {
			return jsast.CreateEmptyStatement(), nil
		}
	}

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
