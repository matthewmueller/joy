package js

// CreateNode fn
func CreateNode(kind string) Node {
	return Node{
		Type: kind,
	}
}

// CreateProgram fn
func CreateProgram(body ...interface{}) Program {
	return Program{
		Type: "Program",
		Body: body,
	}
}

// CreateStatement fn
// func CreateStatement() Statement {
// 	return Statement{
// 		Node: CreateNode("Statement"),
// 	}
// }

// // CreateEmptyStatement fn
// func CreateEmptyStatement() EmptyStatement {
// 	return EmptyStatement{
// 		Statement: CreateNode("EmptyStatement"),
// 	}
// }

// CreateExpressionStatement fn
func CreateExpressionStatement() ExpressionStatement {
	return ExpressionStatement{}
}

// CreateCallExpression fn
func CreateCallExpression() CallExpression {
	return CallExpression{}
}

// CreateFunctionExpression fn
func CreateFunctionExpression() FunctionExpression {
	return FunctionExpression{}
}

// CreateBlockStatement fn
func CreateBlockStatement() BlockStatement {
	return BlockStatement{}
}

// CreateMemberExpression fn
func CreateMemberExpression() MemberExpression {
	return MemberExpression{}
}

// CreateIdentifier fn
func CreateIdentifier() Identifier {
	return Identifier{}
}

// CreateLiteral fn
func CreateLiteral() Literal {
	return Literal{}
}
