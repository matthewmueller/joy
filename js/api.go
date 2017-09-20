package js

import (
	"strings"
)

// CreateIdentifier fn
func CreateIdentifier(name string) Identifier {
	return Identifier{
		Type: "Identifier",
		Name: name,
	}
}

// CreateString fn
func CreateString(value string) Literal {
	return Literal{
		Type:  "Literal",
		Value: value,
	}
}

// CreateBoolean fn
func CreateBoolean(value bool) Literal {
	return Literal{
		Type:  "Literal",
		Value: value,
	}
}

// CreateNull fn
func CreateNull() Literal {
	return Literal{
		Type:  "Literal",
		Value: nil,
	}
}

// CreateInt fn
func CreateInt(n int) Literal {
	return Literal{
		Type:  "Literal",
		Value: n,
	}
}

// CreateFloat fn
func CreateFloat(n float32) Literal {
	return Literal{
		Type:  "Literal",
		Value: n,
	}
}

// CreateRegex fn
func CreateRegex(pattern string, flags []string) RegExpLiteral {
	return RegExpLiteral{
		Type:  "RegExpLiteral",
		Value: pattern,
		Regex: struct {
			Pattern string `json:"pattern,omitempty"`
			Flags   string `json:"flags,omitempty"`
		}{
			Pattern: pattern,
			Flags:   strings.Join(flags, ""),
		},
	}
}

// CreateProgram fn
func CreateProgram(body ...interface{}) Program {
	// ensure the body contains either statements or directives
	l := len(body)

	for i := 0; i < l; i++ {
		child := body[i]
		switch child.(type) {
		// statements || directives
		case IStatement, Directive:
		default:
			panic("a program's body must contain either directives or statements")
		}
	}

	return Program{
		Type: "Program",
		Body: body,
	}
}

// CreateFunction fn
func CreateFunction(id *Identifier, params []IPattern, body FunctionBody) Function {
	return Function{
		ID:     id,
		Params: params,
		Body:   body,
	}
}

// CreateEmptyStatement fn
func CreateEmptyStatement() EmptyStatement {
	return EmptyStatement{
		Type: "EmptyStatement",
	}
}

// CreateCallExpression fn
func CreateCallExpression(callee IExpression, arguments []IExpression) CallExpression {
	// var args []Expression
	// for _, argument := range arguments {
	// 	args = append(args, argument.Expression())
	// }

	return CallExpression{
		Type:      "CallExpression",
		Callee:    callee,
		Arguments: arguments,
	}
}

// CreateFunctionExpression fn
func CreateFunctionExpression(id *Identifier, params []IPattern, body FunctionBody) FunctionExpression {
	return FunctionExpression{
		Type:   "FunctionExpression",
		Params: params,
		Body:   body,
	}
}

// CreateFunctionBody fn
func CreateFunctionBody(body ...interface{}) FunctionBody {
	// ensure the body contains either statements or directives
	l := len(body)

	for i := 0; i < l; i++ {
		child := body[i]
		switch child.(type) {
		// statements || directives
		case IStatement, Directive:
		default:
			panic("a function's body must contain either directives or statements")
		}
	}

	return FunctionBody{
		Type: "FunctionBody",
		Body: body,
	}
}

// CreateExpressionStatement fn
func CreateExpressionStatement(expression IExpression) ExpressionStatement {
	return ExpressionStatement{
		Type:       "ExpressionStatement",
		Expression: expression,
	}
}

// CreateBlockStatement fn
func CreateBlockStatement(body ...IStatement) BlockStatement {
	var args []interface{}
	for _, argument := range body {
		args = append(args, argument.Statement())
	}

	return BlockStatement{
		Type: "BlockStatement",
		Body: args,
	}
}

// CreateMemberExpression fn
func CreateMemberExpression(object IExpression, property IExpression, computed bool) MemberExpression {
	return MemberExpression{
		Type:     "MemberExpression",
		Object:   object,
		Property: property,
		Computed: computed,
	}
}

