package jsast

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type generator interface {
	fmt.Stringer
	generate(ctx *context) string
}

// interface compliance
var _ generator = (*Identifier)(nil)
var _ generator = (*Literal)(nil)
var _ generator = (*RegExpLiteral)(nil)
var _ generator = (*Program)(nil)

// var _ generator = (*Function)(nil)
var _ generator = (*ExpressionStatement)(nil)

// var _ generator = (*Directive)(nil)
var _ generator = (*BlockStatement)(nil)
var _ generator = (*FunctionBody)(nil)
var _ generator = (*EmptyStatement)(nil)

// var _ generator = (*DebuggerStatement)(nil)
// var _ generator = (*WithStatement)(nil)
var _ generator = (*ReturnStatement)(nil)

// var _ generator = (*LabeledStatement)(nil)
var _ generator = (*BreakStatement)(nil)

// var _ generator = (*ContinueStatement)(nil)
var _ generator = (*IfStatement)(nil)

// var _ generator = (*SwitchStatement)(nil)
// var _ generator = (*SwitchCase)(nil)
// var _ generator = (*ThrowStatement)(nil)
// var _ generator = (*TryStatement)(nil)
// var _ generator = (*CatchClause)(nil)
// var _ generator = (*WhileStatement)(nil)
var _ generator = (*ForStatement)(nil)

// var _ generator = (*ForInStatement)(nil)
// var _ generator = (*Declaration)(nil)
var _ generator = (*FunctionDeclaration)(nil)
var _ generator = (*VariableDeclarator)(nil)
var _ generator = (*ThisExpression)(nil)
var _ generator = (*ArrayExpression)(nil)
var _ generator = (*ObjectExpression)(nil)

// var _ generator = (*Property)(nil)
var _ generator = (*FunctionExpression)(nil)

// var _ generator = (*UnaryExpression)(nil)
var _ generator = (*UpdateExpression)(nil)
var _ generator = (*BinaryExpression)(nil)
var _ generator = (*AssignmentExpression)(nil)
var _ generator = (*LogicalExpression)(nil)
var _ generator = (*MemberExpression)(nil)

// var _ generator = (*ConditionalExpression)(nil)
var _ generator = (*CallExpression)(nil)
var _ generator = (*NewExpression)(nil)

var _ generator = (*SequenceExpression)(nil)
var _ generator = (*AwaitExpression)(nil)
var _ generator = (*Raw)(nil)

// var _ fmt.Stringer = (*Pattern)(nil)

type context struct {
	indent int
}

func indent(n int) string {
	return strings.Repeat("\t", n)
}

// Assemble JS from the AST
func Assemble(p Program) (result string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()
	return p.String(), nil
}

func generate(ctx *context, node interface{}) string {
	if node == nil {
		return ""
	}

	// cast to stringer
	generator, ok := node.(generator)
	if !ok {
		n := node.(INode)
		panic(fmt.Errorf("%s || %s does not implement generator", n.Node().Type, reflect.TypeOf(node)))
	}

	return generator.generate(ctx)
}

// String fn
func (n Program) String() string {
	return n.generate(&context{})
}

// Generate function
func (n Program) generate(ctx *context) string {
	var a []string
	for _, child := range n.Body {
		a = append(a, generate(ctx, child))
	}
	return ";" + strings.Join(a, ";\n")
}

// String fn
func (n ExpressionStatement) String() string {
	return n.generate(&context{})
}

// Generate function
func (n ExpressionStatement) generate(ctx *context) string {
	return generate(ctx, n.Expression)
}

// String fn
func (n CallExpression) String() string {
	return n.generate(&context{})
}

// Generate function
func (n CallExpression) generate(ctx *context) string {
	var a []string
	for _, child := range n.Arguments {
		a = append(a, generate(ctx, child))
	}

	c := generate(ctx, n.Callee)

	// call expressions are handled
	// differently if the callee is
	// a function expression
	switch n.Callee.(type) {
	case Identifier, MemberExpression, AwaitExpression, CallExpression:
		return c + "(" + strings.Join(a, ", ") + ")"
	case FunctionExpression:
		return "(" + c + ")(" + strings.Join(a, ", ") + ")"
	default:
		panic(fmt.Errorf("Assembler/CallExpression: unhandled call expression type %s", reflect.TypeOf(n.Callee)))
	}
}

// String fn
func (n FunctionExpression) String() string {
	return n.generate(&context{})
}

// Generate function
func (n FunctionExpression) generate(ctx *context) string {
	var a []string
	for _, child := range n.Params {
		a = append(a, generate(ctx, child))
	}

	body := generate(ctx, n.Body)

	async := ""
	if n.Async {
		async = "async "
	}

	generator := ""
	if n.Generator {
		generator = " * "
	}

	return async + "function" + generator + "(" + strings.Join(a, ", ") + ") " + body
}

// String fn
func (n FunctionDeclaration) String() string {
	return n.generate(&context{})
}

// Generate function
func (n FunctionDeclaration) generate(ctx *context) string {
	var a []string
	for _, child := range n.Params {
		a = append(a, generate(ctx, child))
	}

	body := generate(ctx, n.Body)

	name := ""
	if n.ID != nil {
		n := generate(ctx, *n.ID)
		name = " " + n
	}

	async := ""
	if n.Async {
		async = "async "
	}

	generator := ""
	if n.Generator {
		generator = " *"
	}

	fn := async + "function" + generator + name + " (" + strings.Join(a, ", ") + ") " + body

	return fn
}

// String fn
func (n FunctionBody) String() string {
	return n.generate(&context{})
}

// Generate function
func (n FunctionBody) generate(ctx *context) string {

	var a []string
	ctx.indent++
	for _, child := range n.Body {
		a = append(a, indent(ctx.indent)+generate(ctx, child))
	}
	ctx.indent--

	return "{\n" + strings.Join(a, ";\n") + "\n" + indent(ctx.indent) + "}"
}

// String fn
func (n MemberExpression) String() string {
	return n.generate(&context{})
}

// Generate function
func (n MemberExpression) generate(ctx *context) string {
	obj := generate(ctx, n.Object)
	prop := generate(ctx, n.Property)

	// TODO: there maybe quite a few more
	// cases where we'll want to do this
	switch n.Object.Expression().Type {
	case "LogicalExpression":
		obj = "(" + obj + ")"
	}

	// e.g. hi[world]
	if n.Computed {
		return obj + "[" + prop + "]"
	}

	// hi.world
	return obj + "." + prop
}

// String fn
func (n Identifier) String() string {
	return n.generate(&context{})
}

// Generate function
func (n Identifier) generate(ctx *context) string {
	return n.Name
}

// String fn
func (n Literal) String() string {
	return n.generate(&context{})
}

// TODO: Move thi
// Generate functions into the syntax itself
func (n Literal) generate(ctx *context) string {
	switch t := n.Value.(type) {
	case string:
		return t
	case bool:
		return strconv.FormatBool(t)
	case int:
		return strconv.Itoa(t)
	case float32:
		return strconv.FormatFloat(float64(t), 'f', -1, 32)
	case nil:
		return "null"
	default:
		panic(fmt.Errorf("literal needs to handle %t", reflect.TypeOf(t)))
	}
}

// String fn
func (n RegExpLiteral) String() string {
	return n.generate(&context{})
}

// Generate function
func (n RegExpLiteral) generate(ctx *context) string {
	panic("RegExpLiteral stringer not implemented yet")
}

// String fn
func (n VariableDeclaration) String() string {
	return n.generate(&context{})
}

// Generate function
func (n VariableDeclaration) generate(ctx *context) string {
	var a []string
	for _, child := range n.Declarations {
		a = append(a, generate(ctx, child))
	}
	return n.Kind + " " + strings.Join(a, ", ")
}

// String fn
func (n VariableDeclarator) String() string {
	return n.generate(&context{})
}

// Generate function
func (n VariableDeclarator) generate(ctx *context) string {
	v := generate(ctx, n.ID)
	if n.Init == nil {
		return v
	}

	x := generate(ctx, n.Init)
	return v + " = " + x
}

// String fn
func (n ReturnStatement) String() string {
	return n.generate(&context{})
}

// Generate function
func (n ReturnStatement) generate(ctx *context) string {
	r := generate(ctx, n.Argument)
	return "return " + r + ";"
}

// String fn
func (n ArrayExpression) String() string {
	return n.generate(&context{})
}

// Generate function
func (n ArrayExpression) generate(ctx *context) string {
	var a []string
	for _, child := range n.Elements {
		a = append(a, generate(ctx, child))
	}
	return "[" + strings.Join(a, ", ") + "]"
}

// String fn
func (n BinaryExpression) String() string {
	return n.generate(&context{})
}

// Generate function
func (n BinaryExpression) generate(ctx *context) string {
	l := generate(ctx, n.Left)
	o := generate(ctx, n.Operator)
	r := generate(ctx, n.Right)
	return l + " " + o + " " + r
}

// String fn
func (n BinaryOperator) String() string {
	return n.generate(&context{})
}

// Generate function
func (n BinaryOperator) generate(ctx *context) string {
	return string(n)
}

// String fn
func (n EmptyStatement) String() string {
	return n.generate(&context{})
}

// Generate function
func (n EmptyStatement) generate(ctx *context) string {
	return ""
}

// String fn
func (n ObjectExpression) String() string {
	return n.generate(&context{})
}

// Generate function
func (n ObjectExpression) generate(ctx *context) string {
	if len(n.Properties) == 0 {
		return "{}"
	}

	var props []string
	ctx.indent++
	for _, prop := range n.Properties {
		k := generate(ctx, prop.Key)
		v := generate(ctx, prop.Value)
		props = append(props, indent(ctx.indent)+k+": "+v)
	}
	ctx.indent--
	return "{\n" + strings.Join(props, ",\n") + "\n" + indent(ctx.indent) + "}"
}

// String fn
func (n IfStatement) String() string {
	return n.generate(&context{})
}

// Generate function
func (n IfStatement) generate(ctx *context) string {
	t := generate(ctx, n.Test)
	c := generate(ctx, n.Consequent)

	if n.Alternate == nil {
		return "if (" + t + ") " + c + ""
	}
	a := generate(ctx, n.Alternate)

	return "if (" + t + ") " + c + " else " + a + ""
}

// String fn
func (n BlockStatement) String() string {
	return n.generate(&context{})
}

// Generate function
func (n BlockStatement) generate(ctx *context) string {
	var a []string
	for _, child := range n.Body {
		ctx.indent++
		a = append(a, indent(ctx.indent)+generate(ctx, child))
		ctx.indent--
	}
	return "{\n" + strings.Join(a, "\n") + "\n" + indent(ctx.indent) + "}"
}

// String fn
func (n LogicalExpression) String() string {
	return n.generate(&context{})
}

// Generate function
func (n LogicalExpression) generate(ctx *context) string {
	l := generate(ctx, n.Left)
	r := generate(ctx, n.Right)
	return l + " " + string(n.Operator) + " " + r
}

// String fn
func (n ForStatement) String() string {
	return n.generate(&context{})
}

// Generate function
func (n ForStatement) generate(ctx *context) string {
	var cond []string

	cond = append(cond, generate(ctx, n.Init))
	cond = append(cond, generate(ctx, n.Test))
	cond = append(cond, generate(ctx, n.Update))

	// for body
	body := generate(ctx, n.Body)
	return "for (" + strings.Join(cond, "; ") + ") " + body
}

// String fn
func (n UpdateExpression) String() string {
	return n.generate(&context{})
}

// Generate function
func (n UpdateExpression) generate(ctx *context) string {
	x := generate(ctx, n.Argument)
	return x + string(n.Operator)
}

// String fn
func (n AssignmentExpression) String() string {
	return n.generate(&context{})
}

// Generate function
func (n AssignmentExpression) generate(ctx *context) string {
	l := generate(ctx, n.Left)
	r := generate(ctx, n.Right)
	return l + " " + string(n.Operator) + " " + r
}

// String fn
func (n ThisExpression) String() string {
	return n.generate(&context{})
}

// Generate function
func (n ThisExpression) generate(ctx *context) string {
	return "this"
}

// String fn
func (n NewExpression) String() string {
	return n.generate(&context{})
}

// Generate function
func (n NewExpression) generate(ctx *context) string {
	c := generate(ctx, n.Callee)

	var a []string
	for _, child := range n.Arguments {
		a = append(a, generate(ctx, child))
	}

	return "new " + c + "(" + strings.Join(a, ", ") + ")"
}

// String fn
func (n BreakStatement) String() string {
	return n.generate(&context{})
}

// Generate function
func (n BreakStatement) generate(ctx *context) string {
	return "break"
}

// String fn
func (n SequenceExpression) String() string {
	return n.generate(&context{})
}

// Generate function
func (n SequenceExpression) generate(ctx *context) string {
	var exprs []string
	for _, expr := range n.Expressions {
		exprs = append(exprs, generate(ctx, expr))
	}
	return strings.Join(exprs, ", ")
}

// String fn
func (n ThrowStatement) String() string {
	return n.generate(&context{})
}

// Generate function
func (n ThrowStatement) generate(ctx *context) string {
	return "throw " + generate(ctx, n.Argument)
}

// String fn
func (n AwaitExpression) String() string {
	return n.generate(&context{})
}

// Generate function
func (n AwaitExpression) generate(ctx *context) string {
	return "await " + generate(ctx, n.Argument)
}

// String fn
func (n Raw) String() string {
	return n.generate(&context{})
}

// Generate function
func (n Raw) generate(ctx *context) string {
	return n.Source
}

// String fn
func (n MultiStatement) String() string {
	return n.generate(&context{})
}

// Generate function
func (n MultiStatement) generate(ctx *context) string {
	var a []string
	for _, stmt := range n.Statements {
		a = append(a, generate(ctx, stmt))
	}
	return strings.Join(a, ";\n")
}

// String fn
func (n ForInStatement) String() string {
	return n.generate(&context{})
}

// Generate function
func (n ForInStatement) generate(ctx *context) string {
	left := generate(ctx, n.Left)
	right := generate(ctx, n.Right)
	body := generate(ctx, n.Body)
	return "for (" + left + " in " + right + ") " + body
}

// String fn
// func (n DebuggerStatement) String() string {
// 	return ""
// }

// Generate function
// func (n DebuggerStatement) generate(ctx *context) string {
// 	return "DebuggerStatement"
// }

// Generate function
// func (n WithStatement) generate(ctx *context) string {
// 	return "WithStatement"
// }

// Generate function
// func (n LabeledStatement) generate(ctx *context) string {
// 	return "LabeledStatement"
// }

// Generate function
// func (n ContinueStatement) generate(ctx *context) string {
// 	return "ContinueStatement"
// }

// Generate function
// func (n IfStatement) generate(ctx *context) string {
// 	return "IfStatement"
// }

// Generate function
// func (n SwitchStatement) generate(ctx *context) string {
// 	return "SwitchStatement"
// }

// Generate function
// func (n ThrowStatement) generate(ctx *context) string {
// 	return "ThrowStatement"
// }

// Generate function
// func (n TryStatement) generate(ctx *context) string {
// 	return "TryStatement"
// }

// Generate function
// func (n WhileStatement) generate(ctx *context) string {
// 	return "WhileStatement"
// }

// Generate function
// func (n DoWhileStatement) generate(ctx *context) string {
// 	return "DoWhileStatement"
// }

// Generate function
// func (n Declaration) generate(ctx *context) string {
// 	return "Declaration"
// }
