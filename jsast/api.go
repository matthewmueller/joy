package jsast

import (
	"strings"
)

// Singletons.
var (
	EmptyString = CreateString("")
	True        = CreateBoolean(true)
	False       = CreateBoolean(false)
	Zero        = CreateInt(0)
	Null        = CreateNull()
)

// CreateIdentifier fn
func CreateIdentifier(name string) Identifier {
	return Identifier{
		Type: "Identifier",
		Name: name,
	}
}

// CreateLiteral fn
func CreateLiteral(value string) Literal {
	return Literal{
		Type:  "Identifier",
		Value: value,
	}
}

// CreateString fn
func CreateString(value string) Literal {
	return Literal{
		Type:  "Literal",
		Value: `"` + value + `"`,
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
func CreateFunction(id *Identifier, params []IPattern, body FunctionBody) FunctionDeclaration {
	return FunctionDeclaration{
		Type:   "FunctionDeclaration",
		ID:     id,
		Params: params,
		Body:   body,
	}
}

// CreateGeneratorFunction fn
func CreateGeneratorFunction(id *Identifier, params []IPattern, body FunctionBody) FunctionDeclaration {
	return FunctionDeclaration{
		Type:      "FunctionDeclaration",
		ID:        id,
		Params:    params,
		Body:      body,
		Generator: true,
	}
}

// CreateAsyncFunction fn
func CreateAsyncFunction(id *Identifier, params []IPattern, body FunctionBody) FunctionDeclaration {
	return FunctionDeclaration{
		Type:   "FunctionDeclaration",
		ID:     id,
		Params: params,
		Body:   body,
		Async:  true,
	}
}

// CreateAwaitExpression fn
func CreateAwaitExpression(argument IExpression) AwaitExpression {
	return AwaitExpression{
		Type:     "AwaitExpression",
		Argument: argument,
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

// CreateAsyncFunctionExpression fn
func CreateAsyncFunctionExpression(id *Identifier, params []IPattern, body FunctionBody) FunctionExpression {
	return FunctionExpression{
		Type:   "FunctionExpression",
		ID:     id,
		Params: params,
		Body:   body,
		Async:  true,
	}
}

// CreateGeneratorFunctionExpression fn
func CreateGeneratorFunctionExpression(id *Identifier, params []IPattern, body FunctionBody) FunctionExpression {
	return FunctionExpression{
		Type:      "FunctionExpression",
		ID:        id,
		Params:    params,
		Body:      body,
		Generator: true,
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

// CreateForStatement fn
func CreateForStatement(init interface{}, test IExpression, update IExpression, body IStatement) ForStatement {
	return ForStatement{
		Type:   "ForStatement",
		Init:   init,
		Test:   test,
		Update: update,
		Body:   body,
	}
}

// CreateUpdateExpression fn
func CreateUpdateExpression(argument IExpression, operator UpdateOperator, prefix bool) UpdateExpression {
	return UpdateExpression{
		Type:     "UpdateExpression",
		Argument: argument,
		Operator: operator,
		Prefix:   prefix,
	}
}

// CreateAssignmentExpression fn
func CreateAssignmentExpression(left interface{}, operator AssignmentOperator, right IExpression) AssignmentExpression {
	return AssignmentExpression{
		Type:     "AssignmentExpression",
		Left:     left,
		Operator: operator,
		Right:    right,
	}
}

// CreateThisExpression fn
func CreateThisExpression() ThisExpression {
	return ThisExpression{
		Type: "ThisExpression",
	}
}

// CreateNewExpression fn
func CreateNewExpression(callee IExpression, arguments []IExpression) NewExpression {
	return NewExpression{
		Type:      "NewExpression",
		Callee:    callee,
		Arguments: arguments,
	}
}

// CreateBreakStatement fn
func CreateBreakStatement(label *Identifier) BreakStatement {
	return BreakStatement{
		Type:  "BreakStatement",
		Label: label,
	}
}

// CreateSequenceExpression fn
func CreateSequenceExpression(expressions ...IExpression) SequenceExpression {
	return SequenceExpression{
		Type:        "SequenceExpression",
		Expressions: expressions,
	}
}

// CreateThrowStatement fn
func CreateThrowStatement(argument IExpression) ThrowStatement {
	return ThrowStatement{
		Type:     "ThrowStatement",
		Argument: argument,
	}
}

// CreateRaw fn
func CreateRaw(source string) Raw {
	return Raw{
		Type:   "Raw",
		Source: source,
	}
}

// CreateMultiStatement fn
func CreateMultiStatement(statements ...IStatement) IStatement {
	return MultiStatement{
		Type:       "MultiStatement",
		Statements: statements,
	}
}
