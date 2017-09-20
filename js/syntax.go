package js

// interface inheritance assertions
var _ IExpression = (*Identifier)(nil)
var _ IPattern = (*Identifier)(nil)
var _ IExpression = (*Literal)(nil)
var _ ILiteral = (*RegExpLiteral)(nil)
var _ INode = (*Program)(nil)
var _ INode = (*Function)(nil)
var _ INode = (*Statement)(nil)
var _ IStatement = (*ExpressionStatement)(nil)
var _ IExpressionStatement = (*Directive)(nil)
var _ IStatement = (*BlockStatement)(nil)
var _ IBlockStatement = (*FunctionBody)(nil)
var _ IStatement = (*EmptyStatement)(nil)
var _ IStatement = (*DebuggerStatement)(nil)
var _ IStatement = (*WithStatement)(nil)
var _ IStatement = (*ReturnStatement)(nil)
var _ IStatement = (*LabeledStatement)(nil)
var _ IStatement = (*BreakStatement)(nil)
var _ IStatement = (*ContinueStatement)(nil)
var _ IStatement = (*IfStatement)(nil)
var _ IStatement = (*SwitchStatement)(nil)
var _ INode = (*SwitchCase)(nil)
var _ IStatement = (*ThrowStatement)(nil)
var _ IStatement = (*TryStatement)(nil)
var _ INode = (*CatchClause)(nil)
var _ IStatement = (*WhileStatement)(nil)
var _ IStatement = (*ForStatement)(nil)
var _ IStatement = (*ForInStatement)(nil)
var _ IStatement = (*Declaration)(nil)
var _ IFunction = (*FunctionDeclaration)(nil)
var _ IDeclaration = (*FunctionDeclaration)(nil)
var _ IDeclaration = (*VariableDeclaration)(nil)
var _ INode = (*VariableDeclarator)(nil)
var _ INode = (*Expression)(nil)
var _ IExpression = (*ThisExpression)(nil)
var _ IExpression = (*ArrayExpression)(nil)
var _ IExpression = (*ObjectExpression)(nil)
var _ INode = (*Property)(nil)
var _ IFunction = (*FunctionExpression)(nil)
var _ IExpression = (*FunctionExpression)(nil)
var _ IExpression = (*UnaryExpression)(nil)
var _ IExpression = (*UpdateExpression)(nil)
var _ IExpression = (*BinaryExpression)(nil)
var _ IExpression = (*AssignmentExpression)(nil)
var _ IExpression = (*LogicalExpression)(nil)
var _ IExpression = (*MemberExpression)(nil)
var _ IPattern = (*MemberExpression)(nil)
var _ IExpression = (*ConditionalExpression)(nil)
var _ IExpression = (*CallExpression)(nil)
var _ IExpression = (*NewExpression)(nil)
var _ IExpression = (*SequenceExpression)(nil)
var _ INode = (*Pattern)(nil)

// INode interface
type INode interface {
	Node() Node
}

// IExpression interface
type IExpression interface {
	Expression() Expression
}

// IStatement interface
type IStatement interface {
	Statement() Statement
}

// IPattern interface
type IPattern interface {
	Pattern() Pattern
}

// ILiteral interface
type ILiteral interface {
	Literal() Literal
}

// IExpressionStatement interface
type IExpressionStatement interface {
	ExpressionStatement() ExpressionStatement
}

// IBlockStatement interface
type IBlockStatement interface {
	BlockStatement() BlockStatement
}

// IDeclaration interface
type IDeclaration interface {
	Declaration() Declaration
}

// IFunction interface
type IFunction interface {
	Function() Function
}

// Position struct
type Position struct {
	Line   uint `json:"line,omitempty"`
	Column uint `json:"column,omitempty"`
}

// SourceLocation struct
type SourceLocation struct {
	Source *string  `json:"source,omitempty"`
	Start  Position `json:"start,omitempty"`
	End    Position `json:"end,omitempty"`
}

// Node struct
type Node struct {
	Type string          `json:"type,omitempty"`
	Loc  *SourceLocation `json:"loc,omitempty"`
}

// Identifier struct
type Identifier struct {
	Type string `json:"type,omitempty"`
	Name string `json:"name,omitempty"`
}

// Expression fn
func (n Identifier) Expression() Expression {
	return Expression{
		Type: n.Type,
	}
}

// Pattern fn
func (n Identifier) Pattern() Pattern {
	return Pattern{
		Type: n.Type,
	}
}

// Literal struct
type Literal struct {
	Type  string      `json:"type,omitempty"`
	Value interface{} `json:"value,omitempty"` // string | boolean | null | number | RegExp;
}

// Expression fn
func (n Literal) Expression() Expression {
	return Expression{
		Type: n.Type,
	}
}

// RegExpLiteral struct
type RegExpLiteral struct {
	Type  string      `json:"type,omitempty"`
	Value interface{} `json:"value,omitempty"` // string | boolean | null | number | RegExp;
	Regex struct {
		Pattern string `json:"pattern,omitempty"`
		Flags   string `json:"flags,omitempty"`
	} `json:"regex,omitempty"`
}

// Literal fn
func (n RegExpLiteral) Literal() Literal {
	return Literal{
		Type:  n.Type,
		Value: n.Value,
	}
}

// Program struct
type Program struct {
	Type string `json:"type,omitempty"`

	// [ Directive | Statement ]
	Body []interface{} `json:"body,omitempty"`
}

// Node fn
func (n Program) Node() Node {
	return Node{
		Type: n.Type,
	}
}

// Function struct
type Function struct {
	Type   string
	ID     *Identifier  `json:"id,omitempty"`
	Params []Pattern    `json:"params,omitempty"`
	Body   FunctionBody `json:"body,omitempty"`
}

// Node fn
func (n Function) Node() Node {
	return Node{
		Type: n.Type,
	}
}

// Statement struct
type Statement struct {
	Type string
}

// Node fn
func (n Statement) Node() Node {
	return Node{
		Type: n.Type,
	}
}

// ExpressionStatement struct
type ExpressionStatement struct {
	Type       string     `json:"type,omitempty"`
	Expression Expression `json:"expression,omitempty"`
}

// Statement fn
func (n ExpressionStatement) Statement() Statement {
	return Statement{
		Type: n.Type,
	}
}

// Directive struct
type Directive struct {
	Type       string     `json:"type,omitempty"`
	Expression Expression `json:"expression,omitempty"`
	Directive  string     `json:"directive,omitempty"`
}

// ExpressionStatement fn
func (n Directive) ExpressionStatement() ExpressionStatement {
	return ExpressionStatement{
		Type:       n.Type,
		Expression: n.Expression,
	}
}

// BlockStatement struct
type BlockStatement struct {
	Type string `json:"type,omitempty"`

	// [ Directive | Statement ]
	//
	// TODO: this is non-standard, but it doesn't
	// make sense that a FunctionBody is a
	// BlockStatement but the FunctionBody
	// can contain a DirectiveType while
	// the BlockStatement cannot
	Body []interface{} `json:"body,omitempty"`
}

// Statement fn
func (n BlockStatement) Statement() Statement {
	return Statement{
		Type: n.Type,
	}
}

// FunctionBody struct
type FunctionBody struct {
	Type string `json:"type,omitempty"`

	// [ Directive | Statement ]
	Body []interface{} `json:"body,omitempty"`
}

// BlockStatement fn
func (n FunctionBody) BlockStatement() BlockStatement {
	return BlockStatement{
		Type: n.Type,
		Body: n.Body,
	}
}

// EmptyStatement struct
type EmptyStatement struct {
	Type string `json:"type,omitempty"`
}

// Statement fn
func (n EmptyStatement) Statement() Statement {
	return Statement{
		Type: n.Type,
	}
}

// DebuggerStatement struct
type DebuggerStatement struct {
	Type string `json:"type,omitempty"`
}

// Statement fn
func (n DebuggerStatement) Statement() Statement {
	return Statement{
		Type: n.Type,
	}
}

// WithStatement struct
type WithStatement struct {
	Type   string     `json:"type,omitempty"`
	Object Expression `json:"object,omitempty"`
	Body   Statement  `json:"body,omitempty"`
}

// Statement fn
func (n WithStatement) Statement() Statement {
	return Statement{
		Type: n.Type,
	}
}

// ReturnStatement struct
type ReturnStatement struct {
	Type     string      `json:"type,omitempty"`
	Argument *Expression `json:"argument,omitempty"`
}

// Statement fn
func (n ReturnStatement) Statement() Statement {
	return Statement{
		Type: n.Type,
	}
}

// LabeledStatement struct
type LabeledStatement struct {
	Type  string     `json:"type,omitempty"`
	Label Identifier `json:"label,omitempty"`
	Body  Statement  `json:"body,omitempty"`
}

// Statement fn
func (n LabeledStatement) Statement() Statement {
	return Statement{
		Type: n.Type,
	}
}

// BreakStatement struct
type BreakStatement struct {
	Type  string      `json:"type,omitempty"`
	Label *Identifier `json:"label,omitempty"`
}

// Statement fn
func (n BreakStatement) Statement() Statement {
	return Statement{
		Type: n.Type,
	}
}

// ContinueStatement struct
type ContinueStatement struct {
	Type  string      `json:"type,omitempty"`
	Label *Identifier `json:"label,omitempty"`
}

// Statement fn
func (n ContinueStatement) Statement() Statement {
	return Statement{
		Type: n.Type,
	}
}

// IfStatement struct
type IfStatement struct {
	Type       string     `json:"type,omitempty"`
	Test       Expression `json:"test,omitempty"`
	Consequent Statement  `json:"consequent,omitempty"`
	Alternate  *Statement `json:"alternate,omitempty"`
}

// Statement fn
func (n IfStatement) Statement() Statement {
	return Statement{
		Type: n.Type,
	}
}

// SwitchStatement struct
type SwitchStatement struct {
	Type         string       `json:"type,omitempty"`
	Discriminant Expression   `json:"discriminant,omitempty"`
	Cases        []SwitchCase `json:"cases,omitempty"`
}

// Statement fn
func (n SwitchStatement) Statement() Statement {
	return Statement{
		Type: n.Type,
	}
}

// SwitchCase struct
type SwitchCase struct {
	Type       string      `json:"type,omitempty"`
	Test       *Expression `json:"test,omitempty"`
	Consequent []Statement `json:"consequent,omitempty"`
}

// Node fn
func (n SwitchCase) Node() Node {
	return Node{
		Type: n.Type,
	}
}

// ThrowStatement struct
type ThrowStatement struct {
	Type     string     `json:"type,omitempty"`
	Argument Expression `json:"argument,omitempty"`
}

// Statement fn
func (n ThrowStatement) Statement() Statement {
	return Statement{
		Type: n.Type,
	}
}

// TryStatement struct
type TryStatement struct {
	Type      string          `json:"type,omitempty"`
	Block     BlockStatement  `json:"block,omitempty"`
	Handler   *CatchClause    `json:"handler,omitempty"`
	Finalizer *BlockStatement `json:"finalizer,omitempty"`
}

// Statement fn
func (n TryStatement) Statement() Statement {
	return Statement{
		Type: n.Type,
	}
}

// CatchClause struct
type CatchClause struct {
	Type  string         `json:"type,omitempty"`
	Param Pattern        `json:"param,omitempty"`
	Body  BlockStatement `json:"body,omitempty"`
}

// Node fn
func (n CatchClause) Node() Node {
	return Node{
		Type: n.Type,
	}
}

// WhileStatement struct
type WhileStatement struct {
	Type string     `json:"type,omitempty"`
	Test Expression `json:"test,omitempty"`
	Body Statement  `json:"body,omitempty"`
}

// Statement fn
func (n WhileStatement) Statement() Statement {
	return Statement{
		Type: n.Type,
	}
}

// DoWhileStatement struct
type DoWhileStatement struct {
	Type string     `json:"type,omitempty"`
	Body Statement  `json:"body,omitempty"`
	Test Expression `json:"test,omitempty"`
}

// Statement fn
func (n DoWhileStatement) Statement() Statement {
	return Statement{
		Type: n.Type,
	}
}

// ForStatement struct
type ForStatement struct {
	Type   string      `json:"type,omitempty"`
	Init   interface{} `json:"init,omitempty"` //  VariableDeclaration | Expression | null
	Test   *Expression `json:"test,omitempty"`
	Update *Expression `json:"update,omitempty"`
	Body   Statement   `json:"body,omitempty"`
}

// Statement fn
func (n ForStatement) Statement() Statement {
	return Statement{
		Type: n.Type,
	}
}

// ForInStatement struct
type ForInStatement struct {
	Type  string      `json:"type,omitempty"`
	Left  interface{} `json:"left,omitempty"` // VariableDeclaration |  Pattern
	Right Expression  `json:"right,omitempty"`
	Body  Statement   `json:"body,omitempty"`
}

// Statement fn
func (n ForInStatement) Statement() Statement {
	return Statement{
		Type: n.Type,
	}
}

// Declaration struct
type Declaration struct {
	Type string `json:"type,omitempty"`
}

// Statement fn
func (n Declaration) Statement() Statement {
	return Statement{
		Type: n.Type,
	}
}

// FunctionDeclaration struct
type FunctionDeclaration struct {
	Type   string      `json:"type,omitempty"`
	ID     *Identifier `json:"id,omitempty"`
	Params []Pattern
	Body   FunctionBody
}

// Declaration fn
func (n FunctionDeclaration) Declaration() Declaration {
	return Declaration{
		Type: n.Type,
	}
}

// Function fn
func (n FunctionDeclaration) Function() Function {
	return Function{
		Type:   n.Type,
		ID:     n.ID,
		Params: n.Params,
		Body:   n.Body,
	}
}

// VariableDeclaration struct
type VariableDeclaration struct {
	Type         string               `json:"type,omitempty"`
	Declarations []VariableDeclarator `json:"declarations,omitempty"`
	Kind         string               `json:"kind,omitempty"` // "var"
}

// Declaration fn
func (n VariableDeclaration) Declaration() Declaration {
	return Declaration{
		Type: n.Type,
	}
}

// VariableDeclarator struct
type VariableDeclarator struct {
	Type string      `json:"type,omitempty"`
	ID   Pattern     `json:"id,omitempty"`
	Init *Expression `json:"init,omitempty"`
}

// Node fn
func (n VariableDeclarator) Node() Node {
	return Node{
		Type: n.Type,
	}
}

// Expression struct
type Expression struct {
	Type string
}

// Node fn
func (n Expression) Node() Node {
	return Node{
		Type: n.Type,
	}
}

// ThisExpression struct
type ThisExpression struct {
	Type string `json:"type,omitempty"`
}

// Expression fn
func (n ThisExpression) Expression() Expression {
	return Expression{
		Type: n.Type,
	}
}

// ArrayExpression struct
type ArrayExpression struct {
	Type     string        `json:"type,omitempty"`
	Elements []*Expression `json:"elements,omitempty"`
}

// Expression fn
func (n ArrayExpression) Expression() Expression {
	return Expression{
		Type: n.Type,
	}
}

// ObjectExpression struct
type ObjectExpression struct {
	Type       string     `json:"type,omitempty"`
	Properties []Property `json:"properties,omitempty"`
}

// Expression fn
func (n ObjectExpression) Expression() Expression {
	return Expression{
		Type: n.Type,
	}
}

// Property struct
type Property struct {
	Type  string      `json:"type,omitempty"`
	Key   interface{} `json:"key,omitempty"` // Literal | Identifier
	Value Expression  `json:"value,omitempty"`
	Kind  string      `json:"kind,omitempty"` // "init" | "get" | "set"
}

// Node fn
func (n Property) Node() Node {
	return Node{
		Type: n.Type,
	}
}

// FunctionExpression struct
type FunctionExpression struct {
	Type   string `json:"type,omitempty"`
	ID     *Identifier
	Params []Pattern
	Body   FunctionBody
}

// Function fn
func (n FunctionExpression) Function() Function {
	return Function{
		Type:   n.Type,
		ID:     n.ID,
		Params: n.Params,
		Body:   n.Body,
	}
}

// Expression fn
func (n FunctionExpression) Expression() Expression {
	return Expression{
		Type: n.Type,
	}
}

// UnaryExpression struct
type UnaryExpression struct {
	Type     string        `json:"type,omitempty"`
	Operator UnaryOperator `json:"operator,omitempty"`
	Prefix   bool          `json:"prefix,omitempty"`
	Argument Expression    `json:"argument,omitempty"`
}

// Expression fn
func (n UnaryExpression) Expression() Expression {
	return Expression{
		Type: n.Type,
	}
}

// UnaryOperator enum
// "-" | "+" | "!" | "~" | "typeof" | "void" | "delete"
type UnaryOperator string

// UpdateExpression struct
type UpdateExpression struct {
	Type     string         `json:"type,omitempty"`
	Operator UpdateOperator `json:"operator,omitempty"`
	Argument Expression     `json:"argument,omitempty"`
	Prefix   bool           `json:"prefix,omitempty"`
}

// Expression fn
func (n UpdateExpression) Expression() Expression {
	return Expression{
		Type: n.Type,
	}
}

// UpdateOperator enum
// "++" | "--"
type UpdateOperator string

// BinaryExpression struct
type BinaryExpression struct {
	Type     string         `json:"type,omitempty"`
	Operator BinaryOperator `json:"operator,omitempty"`
	Left     Expression     `json:"left,omitempty"`
	Right    Expression     `json:"right,omitempty"`
}

// Expression fn
func (n BinaryExpression) Expression() Expression {
	return Expression{
		Type: n.Type,
	}
}

// BinaryOperator enum
// "==" | "!=" | "===" | "!=="
// | "<" | "<=" | ">" | ">="
// | "<<" | ">>" | ">>>"
// | "+" | "-" | "*" | "/" | "%"
// | "|" | "^" | "&" | "in"
// | "instanceof"
type BinaryOperator string

// AssignmentExpression struct
type AssignmentExpression struct {
	Type     string             `json:"type,omitempty"`
	Operator AssignmentOperator `json:"operator,omitempty"`
	Left     interface{}        `json:"left,omitempty"` // Pattern | Expression
	Right    Expression         `json:"right,omitempty"`
}

// Expression fn
func (n AssignmentExpression) Expression() Expression {
	return Expression{
		Type: n.Type,
	}
}

// AssignmentOperator enum
// 	"=" | "+=" | "-=" | "*=" | "/=" | "%="
// | "<<=" | ">>=" | ">>>="
// | "|=" | "^=" | "&="
type AssignmentOperator string

// LogicalExpression struct
type LogicalExpression struct {
	Type     string          `json:"type,omitempty"`
	Operator LogicalOperator `json:"operator,omitempty"`
	Left     Expression      `json:"left,omitempty"`
	Right    Expression      `json:"right,omitempty"`
}

// Expression fn
func (n LogicalExpression) Expression() Expression {
	return Expression{
		Type: n.Type,
	}
}

// LogicalOperator struct
// "||" | "&&"
type LogicalOperator string

// MemberExpression struct
//
// A member expression. If computed is true, the node corresponds to
// a computed (a[b]) member expression and property is an Expression.
// If computed is false, the node corresponds to a static (a.b) member
// expression and property is an Identifier.
type MemberExpression struct {
	Type     string     `json:"type,omitempty"`
	Object   Expression `json:"object,omitempty"`
	Property Expression `json:"property,omitempty"`
	Computed bool       `json:"computed,omitempty"`
}

// Expression fn
func (n MemberExpression) Expression() Expression {
	return Expression{
		Type: n.Type,
	}
}

// Pattern fn
func (n MemberExpression) Pattern() Pattern {
	return Pattern{
		Type: n.Type,
	}
}

// ConditionalExpression struct
//
// A conditional expression, i.e., a ternary ?/: expression.
type ConditionalExpression struct {
	Type       string     `json:"type,omitempty"`
	Test       Expression `json:"test,omitempty"`
	Alternate  Expression `json:"alternate,omitempty"`
	Consequent Expression `json:"consequent,omitempty"`
}

// Expression fn
func (n ConditionalExpression) Expression() Expression {
	return Expression{
		Type: n.Type,
	}
}

// CallExpression struct
//
// A function or method call expression.
type CallExpression struct {
	Type      string       `json:"type,omitempty"`
	Callee    Expression   `json:"callee,omitempty"`
	Arguments []Expression `json:"arguments,omitempty"`
}

// Expression fn
func (n CallExpression) Expression() Expression {
	return Expression{
		Type: n.Type,
	}
}

// NewExpression struct
//
// A `new` expression.
type NewExpression struct {
	Type      string       `json:"type,omitempty"`
	Callee    Expression   `json:"callee,omitempty"`
	Arguments []Expression `json:"arguments,omitempty"`
}

// Expression fn
func (n NewExpression) Expression() Expression {
	return Expression{
		Type: n.Type,
	}
}

// SequenceExpression struct
//
// A sequence expression, i.e., a comma-separated sequence of expressions.
type SequenceExpression struct {
	Type        string       `json:"type,omitempty"`
	Expressions []Expression `json:"expressions,omitempty"`
}

// Expression fn
func (n SequenceExpression) Expression() Expression {
	return Expression{
		Type: n.Type,
	}
}

// Pattern struct
//
// Destructuring binding and assignment are not part of ES5,
// but all binding positions accept Pattern to allow for destructuring
// in ES6. Nevertheless, for ES5, the only Pattern subtype is Identifier.
type Pattern struct {
	Type string
}

// Node fn
func (n Pattern) Node() Node {
	return Node{
		Type: n.Type,
	}
}
