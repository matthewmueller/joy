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

// CreateFunctionDeclaration fn
func CreateFunctionDeclaration(id *Identifier, params []IPattern, body FunctionBody) FunctionDeclaration {
	return FunctionDeclaration{
		Type:   "FunctionDeclaration",
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
		args = append(args, argument.(interface{}))
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

// CreateVariableDeclaration fn
func CreateVariableDeclaration(kind string, decls ...VariableDeclarator) VariableDeclaration {
	return VariableDeclaration{
		Type:         "VariableDeclaration",
		Declarations: decls,
		Kind:         kind,
	}
}

// CreateVariableDeclarator fn
func CreateVariableDeclarator(id IPattern, init IExpression) VariableDeclarator {
	return VariableDeclarator{
		Type: "VariableDeclarator",
		ID:   id,
		Init: init,
	}
}

// CreateReturnStatement fn
func CreateReturnStatement(argument IExpression) ReturnStatement {
	return ReturnStatement{
		Type:     "ReturnStatement",
		Argument: argument,
	}
}

// CreateArrayExpression fn
func CreateArrayExpression(elements ...IExpression) ArrayExpression {
	return ArrayExpression{
		Type:     "ArrayExpression",
		Elements: elements,
	}
}

// CreateBinaryExpression fn
func CreateBinaryExpression(l IExpression, op BinaryOperator, r IExpression) BinaryExpression {
	return BinaryExpression{
		Type:     "BinaryExpression",
		Left:     l,
		Operator: op,
		Right:    r,
	}
}

// CreateProperty fn
func CreateProperty(key interface{}, value IExpression, kind string) Property {
	return Property{
		Type:  "Property",
		Key:   key,
		Value: value,
		Kind:  kind,
	}
}

// CreateObjectExpression fn
func CreateObjectExpression(properties []Property) ObjectExpression {
	return ObjectExpression{
		Type:       "ObjectExpression",
		Properties: properties,
	}
}

// CreateIfStatement fn
func CreateIfStatement(test IExpression, consequent IStatement, alternate IStatement) IfStatement {
	return IfStatement{
		Type:       "IfStatement",
		Test:       test,
		Consequent: consequent,
		Alternate:  alternate,
	}
}

// CreateLogicalExpression fn
func CreateLogicalExpression(left IExpression, operator LogicalOperator, right IExpression) LogicalExpression {
	return LogicalExpression{
		Type:     "LogicalExpression",
		Left:     left,
		Operator: operator,
		Right:    right,
	}
}
