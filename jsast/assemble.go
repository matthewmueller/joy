package jsast

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// interface compliance
var _ fmt.Stringer = (*Identifier)(nil)
var _ fmt.Stringer = (*Literal)(nil)
var _ fmt.Stringer = (*RegExpLiteral)(nil)
var _ fmt.Stringer = (*Program)(nil)

// var _ fmt.Stringer = (*Function)(nil)
var _ fmt.Stringer = (*ExpressionStatement)(nil)

// var _ fmt.Stringer = (*Directive)(nil)
var _ fmt.Stringer = (*BlockStatement)(nil)
var _ fmt.Stringer = (*FunctionBody)(nil)
var _ fmt.Stringer = (*EmptyStatement)(nil)

// var _ fmt.Stringer = (*DebuggerStatement)(nil)
// var _ fmt.Stringer = (*WithStatement)(nil)
var _ fmt.Stringer = (*ReturnStatement)(nil)

// var _ fmt.Stringer = (*LabeledStatement)(nil)
var _ fmt.Stringer = (*BreakStatement)(nil)

// var _ fmt.Stringer = (*ContinueStatement)(nil)
var _ fmt.Stringer = (*IfStatement)(nil)

// var _ fmt.Stringer = (*SwitchStatement)(nil)
// var _ fmt.Stringer = (*SwitchCase)(nil)
// var _ fmt.Stringer = (*ThrowStatement)(nil)
// var _ fmt.Stringer = (*TryStatement)(nil)
// var _ fmt.Stringer = (*CatchClause)(nil)
// var _ fmt.Stringer = (*WhileStatement)(nil)
var _ fmt.Stringer = (*ForStatement)(nil)

// var _ fmt.Stringer = (*ForInStatement)(nil)
// var _ fmt.Stringer = (*Declaration)(nil)
var _ fmt.Stringer = (*FunctionDeclaration)(nil)
var _ fmt.Stringer = (*VariableDeclarator)(nil)
var _ fmt.Stringer = (*ThisExpression)(nil)
var _ fmt.Stringer = (*ArrayExpression)(nil)
var _ fmt.Stringer = (*ObjectExpression)(nil)

// var _ fmt.Stringer = (*Property)(nil)
var _ fmt.Stringer = (*FunctionExpression)(nil)

// var _ fmt.Stringer = (*UnaryExpression)(nil)
var _ fmt.Stringer = (*UpdateExpression)(nil)
var _ fmt.Stringer = (*BinaryExpression)(nil)
var _ fmt.Stringer = (*AssignmentExpression)(nil)
var _ fmt.Stringer = (*LogicalExpression)(nil)
var _ fmt.Stringer = (*MemberExpression)(nil)

// var _ fmt.Stringer = (*ConditionalExpression)(nil)
var _ fmt.Stringer = (*CallExpression)(nil)
var _ fmt.Stringer = (*NewExpression)(nil)

var _ fmt.Stringer = (*SequenceExpression)(nil)
var _ fmt.Stringer = (*AwaitExpression)(nil)

// var _ fmt.Stringer = (*Pattern)(nil)

// Assemble JS from the AST
func Assemble(node interface{}) string {
	// pretty.Println(node)
	return stringify(node)
}

func stringify(node interface{}) string {
	if node == nil {
		return ""
	}

	// cast to stringer
	stringer, ok := node.(fmt.Stringer)
	if !ok {
		n := node.(INode)
		panic(fmt.Errorf("%s || %s does not implement stringer", n.Node().Type, reflect.TypeOf(node)))
	}

	return stringer.String()
}

func (n Program) String() string {
	var a []string
	for _, child := range n.Body {
		a = append(a, stringify(child))
	}
	return strings.Join(a, ";\n")
}

func (n ExpressionStatement) String() string {
	return stringify(n.Expression)
}

func (n CallExpression) String() string {
	var a []string
	for _, child := range n.Arguments {
		a = append(a, stringify(child))
	}

	c := stringify(n.Callee)

	// call expressions are handled
	// differently if the callee is
	// a function expression
	switch n.Callee.(type) {
	case Identifier, MemberExpression:
		return c + "(" + strings.Join(a, ", ") + ")"
	case FunctionExpression:
		return "(" + c + ")(" + strings.Join(a, ", ") + ")"
	default:
		panic(fmt.Errorf("Assembler/CallExpression: unhandled call expression type %s", reflect.TypeOf(n.Callee)))
	}
}

func (n FunctionExpression) String() string {
	var a []string
	for _, child := range n.Params {
		a = append(a, stringify(child))
	}

	body := stringify(n.Body)

	return "function(" + strings.Join(a, ", ") + ") {\n" + body + "\n}"
}

func (n FunctionDeclaration) String() string {
	var a []string
	for _, child := range n.Params {
		a = append(a, stringify(child))
	}

	body := stringify(n.Body)

	name := ""
	if n.ID != nil {
		n := stringify(*n.ID)
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

	fn := async + "function" + generator + name + " (" + strings.Join(a, ", ") + ") {\n" + body + "\n}"

	return fn
}

func (n FunctionBody) String() string {

	var a []string
	for _, child := range n.Body {
		a = append(a, stringify(child))
	}

	return strings.Join(a, ";\n") + ";"
}

func (n MemberExpression) String() string {
	obj := stringify(n.Object)
	prop := stringify(n.Property)

	// e.g. hi[world]
	if n.Computed {
		return obj + "[" + prop + "]"
	}

	// hi.world
	return obj + "." + prop
}

func (n Identifier) String() string {
	return n.Name
}

// TODO: Move this into the syntax itself
func (n Literal) String() string {
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

func (n RegExpLiteral) String() string {
	panic("RegExpLiteral stringer not implemented yet")
}

func (n VariableDeclaration) String() string {
	var a []string
	for _, child := range n.Declarations {
		a = append(a, stringify(child))
	}
	return n.Kind + " " + strings.Join(a, ", ")
}

func (n VariableDeclarator) String() string {
	v := stringify(n.ID)
	if n.Init == nil {
		return v
	}

	x := stringify(n.Init)
	return v + " = " + x
}

func (n ReturnStatement) String() string {
	r := stringify(n.Argument)
	return "return " + r
}

func (n ArrayExpression) String() string {
	var a []string
	for _, child := range n.Elements {
		a = append(a, stringify(child))
	}
	return "[" + strings.Join(a, ", ") + "]"
}

func (n BinaryExpression) String() string {
	l := stringify(n.Left)
	o := stringify(n.Operator)
	r := stringify(n.Right)
	return l + " " + o + " " + r
}

func (n BinaryOperator) String() string {
	return string(n)
}

func (n EmptyStatement) String() string {
	return ""
}

func (n ObjectExpression) String() string {
	var props []string
	for _, prop := range n.Properties {
		k := stringify(prop.Key)
		v := stringify(prop.Value)
		props = append(props, "  "+k+": "+v)
	}
	return "{\n" + strings.Join(props, ",\n") + "\n}"
}

func (n IfStatement) String() string {
	t := stringify(n.Test)
	c := stringify(n.Consequent)

	if n.Alternate == nil {
		return "if (" + t + ") " + c + ""
	}
	a := stringify(n.Alternate)

	return "if (" + t + ") " + c + " else " + a + ""
}

func (n BlockStatement) String() string {
	var a []string
	for _, child := range n.Body {
		a = append(a, stringify(child))
	}
	return "{\n" + strings.Join(a, "\n") + "\n}"
}

func (n LogicalExpression) String() string {
	l := stringify(n.Left)
	r := stringify(n.Right)
	return l + " " + string(n.Operator) + " " + r
}

func (n ForStatement) String() string {
	var cond []string

	cond = append(cond, stringify(n.Init))
	cond = append(cond, stringify(n.Test))
	cond = append(cond, stringify(n.Update))

	// for body
	body := stringify(n.Body)
	return "for (" + strings.Join(cond, "; ") + ") " + body
}

func (n UpdateExpression) String() string {
	x := stringify(n.Argument)
	return x + string(n.Operator)
}

func (n AssignmentExpression) String() string {
	l := stringify(n.Left)
	r := stringify(n.Right)
	return l + " " + string(n.Operator) + " " + r
}

func (n ThisExpression) String() string {
	return "this"
}

func (n NewExpression) String() string {
	c := stringify(n.Callee)

	var a []string
	for _, child := range n.Arguments {
		a = append(a, stringify(child))
	}

	return "new " + c + "(" + strings.Join(a, ", ") + ")"
}

func (n BreakStatement) String() string {
	return "break"
}

func (n SequenceExpression) String() string {
	var exprs []string
	for _, expr := range n.Expressions {
		exprs = append(exprs, stringify(expr))
	}
	return strings.Join(exprs, ", ")
}

func (n AwaitExpression) String() string {
	return "await " + stringify(n.Argument)
}

// func (n DebuggerStatement) String() string {
// 	return "DebuggerStatement", nil
// }

// func (n WithStatement) String() string {
// 	return "WithStatement", nil
// }

// func (n LabeledStatement) String() string {
// 	return "LabeledStatement", nil
// }

// func (n ContinueStatement) String() string {
// 	return "ContinueStatement", nil
// }

// func (n IfStatement) String() string {
// 	return "IfStatement", nil
// }

// func (n SwitchStatement) String() string {
// 	return "SwitchStatement", nil
// }

// func (n ThrowStatement) String() string {
// 	return "ThrowStatement", nil
// }

// func (n TryStatement) String() string {
// 	return "TryStatement", nil
// }

// func (n WhileStatement) String() string {
// 	return "WhileStatement", nil
// }

// func (n DoWhileStatement) String() string {
// 	return "DoWhileStatement", nil
// }

// func (n ForStatement) String() string {
// 	return "ForStatement", nil
// }

// func (n ForInStatement) String() string {
// 	return "ForInStatement", nil
// }

// func (n Declaration) String() string {
// 	return "Declaration", nil
// }
