package js

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/kr/pretty"
)

// Stringer interface
type Stringer interface {
	String() (string, error)
}

var _ Stringer = (*CallExpression)(nil)

// Assemble JS from the AST
func Assemble(node interface{}) (string, error) {
	pretty.Println(node)

	switch t := node.(type) {
	case Program:
		return t.String()
	}

	return "", errors.New("ast must start with a program")
}

func stringify(node interface{}) (string, error) {
	if node == nil {
		return "", nil
	}
	// cast to stringer
	stringer, ok := node.(Stringer)
	if !ok {
		n := node.(INode)
		return "", fmt.Errorf("%s || %s does not implement stringer", n.Node().Type, reflect.TypeOf(node))
	}

	return stringer.String()
}

func stringifyEach(nodes ...interface{}) ([]string, error) {
	var out []string

	for _, node := range nodes {
		s, e := stringify(node)
		if e != nil {
			return nil, e
		}
		out = append(out, s)
	}

	return out, nil
}

func (p *Program) String() (string, error) {
	results := []string{}
	body := p.Body
	l := len(body)

	for i := 0; i < l; i++ {
		child := body[i]
		switch child.(type) {
		// statements || directives
		case IStatement, Directive:
		default:
			return "", errors.New("a program's body must contain either directives or statements")
		}

		// string
		s, e := stringify(child)
		if e != nil {
			return "", e
		}
		results = append(results, s)
	}

	return strings.Join(results, ";\n"), nil
}

func (n ExpressionStatement) String() (string, error) {
	return stringify(n.Expression)
}

func (n CallExpression) String() (string, error) {
	var args []string
	for _, arg := range n.Arguments {
		v, e := stringify(arg)
		if e != nil {
			return "", e
		}
		args = append(args, v)
	}

	c, e := stringify(n.Callee)
	if e != nil {
		return "", e
	}

	// call expressions are handled
	// differently if the callee is
	// a function expression
	switch n.Callee.(type) {
	case Identifier, MemberExpression:
		return c + "(" + strings.Join(args, ", ") + ")", nil
	case FunctionExpression:
		return "(" + c + ")(" + strings.Join(args, ", ") + ")", nil
	default:
		return "", fmt.Errorf("Assembler/CallExpression: unhandled call expression type %s", reflect.TypeOf(n.Callee))
	}
}

func (n FunctionExpression) String() (string, error) {
	var params []string
	for _, param := range n.Params {
		s, e := stringify(param)
		if e != nil {
			return "", e
		}
		params = append(params, s)
	}

	body, e := stringify(n.Body)
	if e != nil {
		return "", e
	}

	fn := "function(" + strings.Join(params, ", ") + ") {\n" + body + "\n}"
	return fn, nil
}

func (n FunctionDeclaration) String() (string, error) {
	var params []string
	for _, param := range n.Params {
		s, e := stringify(param)
		if e != nil {
			return "", e
		}
		params = append(params, s)
	}

	body, e := stringify(n.Body)
	if e != nil {
		return "", e
	}

	name := ""
	if n.ID != nil {
		n, e := stringify(*n.ID)
		if e != nil {
			return "", e
		}
		name = " " + n
	}

	fn := "function" + name + " (" + strings.Join(params, ", ") + ") { \n" + body + "\n}"
	return fn, nil
}

func (n FunctionBody) String() (string, error) {
	var stmts []string

	// TODO: statement | directive
	for _, statement := range n.Body {
		s, e := stringify(statement)
		if e != nil {
			return "", e
		}
		stmts = append(stmts, s)
	}

	return strings.Join(stmts, ";\n") + ";", nil
}

func (n MemberExpression) String() (string, error) {
	obj, e := stringify(n.Object)
	if e != nil {
		return "", e
	}

	prop, e := stringify(n.Property)
	if e != nil {
		return "", e
	}

	// e.g. hi[world]
	if n.Computed {
		return obj + "[" + prop + "]", nil
	}

	// hi.world
	return obj + "." + prop, nil
}

func (n Identifier) String() (string, error) {
	return n.Name, nil
}

// TODO: Move this into the syntax itself
func (n Literal) String() (string, error) {
	switch t := n.Value.(type) {
	case string:
		return t, nil
	case bool:
		return strconv.FormatBool(t), nil
	case int:
		return strconv.Itoa(t), nil
	case float32:
		return strconv.FormatFloat(float64(t), 'f', -1, 32), nil
	case nil:
		return "null", nil
	default:
		return "", fmt.Errorf("literal needs to handle %t", reflect.TypeOf(t))
	}
}

func (n VariableDeclaration) String() (string, error) {
	var decls []string
	for _, decl := range n.Declarations {
		d, e := stringify(decl)
		if e != nil {
			return "", e
		}
		decls = append(decls, d)
	}

	return n.Kind + " " + strings.Join(decls, ", "), nil
}

func (n VariableDeclarator) String() (string, error) {
	v, e := stringify(n.ID)
	if e != nil {
		return "", e
	}
	x, e := stringify(n.Init)
	if e != nil {
		return "", e
	}
	return v + " = " + x, nil
}

func (n ReturnStatement) String() (string, error) {
	r, e := stringify(n.Argument)
	if e != nil {
		return "", e
	}

	return "return " + r, nil
}

func (n ArrayExpression) String() (string, error) {
	var decls []string
	for _, decl := range n.Elements {
		d, e := stringify(decl)
		if e != nil {
			return "", e
		}
		decls = append(decls, d)
	}
	return "[" + strings.Join(decls, ", ") + "]", nil
}

func (n BinaryExpression) String() (string, error) {
	l, e := stringify(n.Left)
	if e != nil {
		return "", e
	}
	o, e := stringify(n.Operator)
	if e != nil {
		return "", e
	}
	r, e := stringify(n.Right)
	if e != nil {
		return "", e
	}

	return l + " " + o + " " + r, nil
}

func (n BinaryOperator) String() (string, error) {
	return string(n), nil
}

// func (n BlockStatement) String() (string, error) {
// 	return "BlockStatement", nil
// }

func (n EmptyStatement) String() (string, error) {
	return "", nil
}

func (n ObjectExpression) String() (string, error) {
	var props []string

	for _, prop := range n.Properties {
		k, e := stringify(prop.Key)
		if e != nil {
			return "", e
		}

		v, e := stringify(prop.Value)
		if e != nil {
			return "", e
		}

		props = append(props, "  "+k+": "+v)
	}

	return "{\n" + strings.Join(props, ",\n") + "\n}", nil
}

func (n IfStatement) String() (string, error) {
	t, e := stringify(n.Test)
	if e != nil {
		return "", e
	}

	c, e := stringify(n.Consequent)
	if e != nil {
		return "", e
	}

	if n.Alternate == nil {
		return "if (" + t + ") " + c + "", nil
	}

	a, e := stringify(n.Alternate)
	if e != nil {
		return "", e
	}

	return "if (" + t + ") " + c + " else " + a + "", nil
}

func (n BlockStatement) String() (string, error) {
	var stmts []string

	for _, stmt := range n.Body {
		switch t := stmt.(type) {
		case IStatement:
			s, e := stringify(t)
			if e != nil {
				return "", e
			}
			stmts = append(stmts, s)
		default:
			return "", fmt.Errorf("block statements can only contain statements, but we got %s", reflect.TypeOf(stmt))
		}
	}

	return "{\n" + strings.Join(stmts, "\n") + "\n}", nil
}

func (n LogicalExpression) String() (string, error) {
	l, e := stringify(n.Left)
	if e != nil {
		return "", e
	}

	r, e := stringify(n.Right)
	if e != nil {
		return "", e
	}

	switch string(n.Operator) {
	case "&&", "||":
	default:
		return "", errors.New("LogicalExpression: logical expression must be either && or ||")
	}

	return l + " " + string(n.Operator) + " " + r, nil
}

func (n ForStatement) String() (string, error) {
	var cond []string
	var e error

	// since this is a union type, we'll need to typecheck
	var init string
	switch n.Init.(type) {
	case VariableDeclaration, IExpression:
		init, e = stringify(n.Init)
	case nil:
		init = ""
	default:
		return "", fmt.Errorf("ForStatement: init can only be VariableDeclaration | Expression | null, but got %s", reflect.TypeOf(n.Init))
	}
	if e != nil {
		return "", e
	}
	cond = append(cond, init)

	// expression
	test, e := stringify(n.Test)
	if e != nil {
		return "", e
	}
	cond = append(cond, test)

	// expression
	update, e := stringify(n.Update)
	if e != nil {
		return "", e
	}
	cond = append(cond, update)

	// for body
	b, e := stringify(n.Body)
	if e != nil {
		return "", e
	}

	return "for (" + strings.Join(cond, "; ") + ") {\n" + b + "\n}", nil
}

func (n UpdateExpression) String() (string, error) {
	x, e := stringify(n.Argument)
	if e != nil {
		return "", e
	}

	return x + string(n.Operator), nil
}

// func (n DebuggerStatement) String() (string, error) {
// 	return "DebuggerStatement", nil
// }

// func (n WithStatement) String() (string, error) {
// 	return "WithStatement", nil
// }

// func (n LabeledStatement) String() (string, error) {
// 	return "LabeledStatement", nil
// }

// func (n BreakStatement) String() (string, error) {
// 	return "BreakStatement", nil
// }

// func (n ContinueStatement) String() (string, error) {
// 	return "ContinueStatement", nil
// }

// func (n IfStatement) String() (string, error) {
// 	return "IfStatement", nil
// }

// func (n SwitchStatement) String() (string, error) {
// 	return "SwitchStatement", nil
// }

// func (n ThrowStatement) String() (string, error) {
// 	return "ThrowStatement", nil
// }

// func (n TryStatement) String() (string, error) {
// 	return "TryStatement", nil
// }

// func (n WhileStatement) String() (string, error) {
// 	return "WhileStatement", nil
// }

// func (n DoWhileStatement) String() (string, error) {
// 	return "DoWhileStatement", nil
// }

// func (n ForStatement) String() (string, error) {
// 	return "ForStatement", nil
// }

// func (n ForInStatement) String() (string, error) {
// 	return "ForInStatement", nil
// }

// func (n Declaration) String() (string, error) {
// 	return "Declaration", nil
// }
