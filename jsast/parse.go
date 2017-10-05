package jsast

import (
	"encoding/json"
	"fmt"

	"github.com/dop251/goja"
	"github.com/matthewmueller/golly/bindata"
	"github.com/pkg/errors"
)

var vm *goja.Runtime
var parse goja.Callable

// this represents all possible token types
// TODO: is there a better way to do this
// without parsing a ton of times?
type node struct {
	Type       string  `json:"type,omitempty"`
	Body       []*node `json:"body,omitempty"`
	Argument   *node   `json:"argument,omitempty"`
	Expression *node   `json:"expression,omitempty"`
	Arguments  []*node `json:"arguments,omitempty"`
	Callee     *node   `json:"callee,omitempty"`
	Object     *node   `json:"object,omitempty"`
	Property   *node   `json:"property,omitempty"`
	Computed   bool    `json:"computed,omitempty"`
	Name       string  `json:"name,omitempty"`
}

// TODO: replace this with a parser written in
// golang at some point. I kind of messed this up
// since I used acorn's AST instead of otto's/goja's.
// Acorn's is really nice though because it enforces
// ES3 and supports outer return statements. Probably
// easier to switch our AST format than write a parser.
func load() {
	acornSrc := bindata.MustAsset("bindata/acorn.js")

	vm := goja.New()
	prg, err := goja.Compile("acorn.js", string(acornSrc), false)
	if err != nil {
		panic(err)
	}

	res, err := vm.RunProgram(prg)
	if err != nil {
		panic(err)
	}

	fn, ok := goja.AssertFunction(res)
	if !ok {
		panic(errors.New("couldn't get the parse function"))
	}

	parse = fn
}

// Parse the source
func Parse(src string) (stmts []IStatement, err error) {
	// lazy-load acorn:
	// - ~23ms load overhead
	// - parse time in the microseconds after that
	if parse == nil {
		load()
	}

	res, err := parse(nil, vm.ToValue(src))
	if err != nil {
		return nil, err
	}

	// export the JSON response to []byte
	var nodes []byte
	if e := vm.ExportTo(res, &nodes); e != nil {
		return nil, e
	}

	// parse those []bytes into a walkable ASTs
	var ast node
	if e := json.Unmarshal(nodes, &ast); e != nil {
		return nil, e
	}

	// walk the ast converting the JS ast nodes into
	// the proper golang structs of JS ast nodes
	// defined in jsast/api.go
	return walk(ast)
}

func walk(program node) (stmts []IStatement, err error) {
	for _, stmt := range program.Body {
		if stmt == nil {
			continue
		}

		s, e := statement(*stmt)
		if e != nil {
			return nil, e
		}
		stmts = append(stmts, s)
	}

	return stmts, nil
}

func statement(n node) (IStatement, error) {
	switch n.Type {
	case "ExpressionStatement":
		return expressionStatement(n)

	case "ReturnStatement":
		return returnStatement(n)
	default:
		return nil, unhandled("statement", n.Type)
	}
}

func expression(n node) (IExpression, error) {
	switch n.Type {
	case "CallExpression":
		return callExpression(n)
	case "MemberExpression":
		return memberExpression(n)
	case "Identifier":
		return identifier(n)
	default:
		return nil, unhandled("expression", n.Type)
	}
}

func expressionStatement(n node) (s ExpressionStatement, err error) {
	x, e := expression(*n.Expression)
	if e != nil {
		return s, e
	}
	return CreateExpressionStatement(x), nil
}

func returnStatement(n node) (s ReturnStatement, err error) {
	x, e := expression(*n.Argument)
	if e != nil {
		return s, e
	}
	return CreateReturnStatement(x), nil
}

func callExpression(n node) (x CallExpression, err error) {
	c, e := expression(*n.Callee)
	if e != nil {
		return x, e
	}

	var args []IExpression
	for _, argument := range n.Arguments {
		a, e := expression(*argument)
		if e != nil {
			return x, e
		}
		args = append(args, a)
	}

	return CreateCallExpression(c, args), nil
}

func memberExpression(n node) (x MemberExpression, err error) {
	o, e := expression(*n.Object)
	if e != nil {
		return x, e
	}

	p, e := expression(*n.Property)
	if e != nil {
		return x, e
	}

	return CreateMemberExpression(o, p, n.Computed), nil
}

func identifier(n node) (Identifier, error) {
	return CreateIdentifier(n.Name), nil
}

func unhandled(fn, kind string) error {
	return fmt.Errorf("unhandled function='%s' type='%s'", fn, kind)
}
