package jsast

// interface inheritance assertions
var _ INode = (*Node)(nil)
var _ INode = (*Identifier)(nil)
var _ IExpression = (*Identifier)(nil)
var _ IPattern = (*Identifier)(nil)
var _ INode = (*Literal)(nil)
var _ ILiteral = (*Literal)(nil)
var _ IExpression = (*Literal)(nil)
var _ INode = (*RegExpLiteral)(nil)
var _ IExpression = (*RegExpLiteral)(nil)
var _ ILiteral = (*RegExpLiteral)(nil)
var _ INode = (*Program)(nil)
var _ INode = (*Function)(nil)
var _ IFunction = (*Function)(nil)
var _ INode = (*Statement)(nil)
var _ IStatement = (*Statement)(nil)
var _ INode = (*ExpressionStatement)(nil)
var _ IStatement = (*ExpressionStatement)(nil)
var _ IExpressionStatement = (*ExpressionStatement)(nil)
var _ INode = (*Directive)(nil)
var _ IStatement = (*Directive)(nil)
var _ IExpressionStatement = (*Directive)(nil)
var _ INode = (*BlockStatement)(nil)
var _ IStatement = (*BlockStatement)(nil)
var _ IBlockStatement = (*BlockStatement)(nil)
var _ INode = (*FunctionBody)(nil)
var _ IStatement = (*FunctionBody)(nil)
var _ IBlockStatement = (*FunctionBody)(nil)
var _ INode = (*EmptyStatement)(nil)
var _ IStatement = (*EmptyStatement)(nil)
var _ INode = (*DebuggerStatement)(nil)
var _ IStatement = (*DebuggerStatement)(nil)
var _ INode = (*WithStatement)(nil)
var _ IStatement = (*WithStatement)(nil)
var _ INode = (*ReturnStatement)(nil)
var _ IStatement = (*ReturnStatement)(nil)
var _ INode = (*LabeledStatement)(nil)
var _ IStatement = (*LabeledStatement)(nil)
var _ INode = (*BreakStatement)(nil)
var _ IStatement = (*BreakStatement)(nil)
var _ INode = (*ContinueStatement)(nil)
var _ IStatement = (*ContinueStatement)(nil)
var _ INode = (*IfStatement)(nil)
var _ IStatement = (*IfStatement)(nil)
var _ INode = (*SwitchStatement)(nil)
var _ IStatement = (*SwitchStatement)(nil)
var _ INode = (*SwitchCase)(nil)
var _ INode = (*ThrowStatement)(nil)
var _ IStatement = (*ThrowStatement)(nil)
var _ INode = (*TryStatement)(nil)
var _ IStatement = (*TryStatement)(nil)
var _ INode = (*CatchClause)(nil)
var _ INode = (*WhileStatement)(nil)
var _ IStatement = (*WhileStatement)(nil)
var _ INode = (*ForStatement)(nil)
var _ IStatement = (*ForStatement)(nil)
var _ INode = (*ForInStatement)(nil)
var _ IStatement = (*ForInStatement)(nil)
var _ INode = (*Declaration)(nil)
var _ IStatement = (*Declaration)(nil)
var _ IDeclaration = (*Declaration)(nil)
var _ INode = (*FunctionDeclaration)(nil)
var _ IStatement = (*FunctionDeclaration)(nil)
var _ IFunction = (*FunctionDeclaration)(nil)
var _ IDeclaration = (*FunctionDeclaration)(nil)
var _ INode = (*VariableDeclaration)(nil)
var _ IStatement = (*VariableDeclaration)(nil)
var _ IDeclaration = (*VariableDeclaration)(nil)
var _ INode = (*VariableDeclarator)(nil)
var _ INode = (*Expression)(nil)
var _ IExpression = (*Expression)(nil)
var _ INode = (*ThisExpression)(nil)
var _ IExpression = (*ThisExpression)(nil)
var _ INode = (*ArrayExpression)(nil)
var _ IExpression = (*ArrayExpression)(nil)
var _ INode = (*ObjectExpression)(nil)
var _ IExpression = (*ObjectExpression)(nil)
var _ INode = (*Property)(nil)
var _ INode = (*FunctionExpression)(nil)
var _ IFunction = (*FunctionExpression)(nil)
var _ IExpression = (*FunctionExpression)(nil)
var _ INode = (*UnaryExpression)(nil)
var _ IExpression = (*UnaryExpression)(nil)
var _ INode = (*UpdateExpression)(nil)
var _ IExpression = (*UpdateExpression)(nil)
var _ INode = (*BinaryExpression)(nil)
var _ IExpression = (*BinaryExpression)(nil)
var _ INode = (*AssignmentExpression)(nil)
var _ IExpression = (*AssignmentExpression)(nil)
var _ INode = (*LogicalExpression)(nil)
var _ IExpression = (*LogicalExpression)(nil)
var _ INode = (*MemberExpression)(nil)
var _ IExpression = (*MemberExpression)(nil)
var _ INode = (*MemberExpression)(nil)
var _ IPattern = (*MemberExpression)(nil)
var _ INode = (*ConditionalExpression)(nil)
var _ IExpression = (*ConditionalExpression)(nil)
var _ INode = (*CallExpression)(nil)
var _ IExpression = (*CallExpression)(nil)
var _ INode = (*NewExpression)(nil)
var _ IExpression = (*NewExpression)(nil)
var _ INode = (*SequenceExpression)(nil)
var _ IExpression = (*SequenceExpression)(nil)
var _ INode = (*AwaitExpression)(nil)
var _ IExpression = (*AwaitExpression)(nil)
var _ INode = (*Pattern)(nil)
var _ IPattern = (*Pattern)(nil)

// INode interface
type INode interface {
	Node() Node
}

// IExpression interface
type IExpression interface {
	INode
	Expression() Expression
}

// IStatement interface
type IStatement interface {
	INode
	Statement() Statement
}

// IPattern interface
type IPattern interface {
	INode
	Pattern() Pattern
}

// ILiteral interface
type ILiteral interface {
	IExpression
	Literal() Literal
}

// IExpressionStatement interface
type IExpressionStatement interface {
	IStatement
	ExpressionStatement() ExpressionStatement
}

// IBlockStatement interface
type IBlockStatement interface {
	IStatement
	BlockStatement() BlockStatement
}

// IDeclaration interface
type IDeclaration interface {
	IStatement
	Declaration() Declaration
}

// IFunction interface
type IFunction interface {
	INode
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

// Node fn
func (n Node) Node() Node {
	return n
}

// Identifier struct
type Identifier struct {
	Type string `json:"type,omitempty"`
	Name string `json:"name,omitempty"`
}

// Node fn
func (n Identifier) Node() Node {
	return Node{
		Type: n.Type,
	}
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

// Node fn
func (n Literal) Node() Node {
	return Node{
		Type: n.Type,
	}
}

// Expression fn
func (n Literal) Expression() Expression {
	return Expression{
		Type: n.Type,
	}
}

// Literal fn
func (n Literal) Literal() Literal {
	return n
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

// Node fn
func (n RegExpLiteral) Node() Node {
	return Node{
		Type: n.Type,
	}
}

// Expression fn
func (n RegExpLiteral) Expression() Expression {
	return Expression{
		Type: n.Type,
	}
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
	Params []IPattern   `json:"params,omitempty"`
	Body   FunctionBody `json:"body,omitempty"`
}

// Node fn
func (n Function) Node() Node {
	return Node{
		Type: n.Type,
	}
}

// Function fn
func (n Function) Function() Function {
	return n
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

// Statement fn
func (n Statement) Statement() Statement {
	return n
}

// ExpressionStatement struct
type ExpressionStatement struct {
	Type       string      `json:"type,omitempty"`
	Expression IExpression `json:"expression,omitempty"`
}

// Node fn
func (n ExpressionStatement) Node() Node {
	return Node{
		Type: n.Type,
	}
}

// Statement fn
func (n ExpressionStatement) Statement() Statement {
	return Statement{
		Type: n.Type,
	}
}

// ExpressionStatement fn
func (n ExpressionStatement) ExpressionStatement() ExpressionStatement {
	return n
}

// Directive struct
type Directive struct {
	Type       string      `json:"type,omitempty"`
	Expression IExpression `json:"expression,omitempty"`
	Directive  string      `json:"directive,omitempty"`
}

// Node fn
func (n Directive) Node() Node {
	return Node{
		Type: n.Type,
	}
}

// Statement fn
func (n Directive) Statement() Statement {
	return Statement{
		Type: n.Type,
	}
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

// Node fn
func (n BlockStatement) Node() Node {
	return Node{
		Type: n.Type,
	}
}

// Statement fn
func (n BlockStatement) Statement() Statement {
	return Statement{
		Type: n.Type,
	}
}

// BlockStatement fn
func (n BlockStatement) BlockStatement() BlockStatement {
	return n
}

// FunctionBody struct
type FunctionBody struct {
	Type string `json:"type,omitempty"`

	// [ Directive | Statement ]
	Body []interface{} `json:"body,omitempty"`
}

// Node fn
func (n FunctionBody) Node() Node {
	return Node{
		Type: n.Type,
	}
}

// Statement fn
func (n FunctionBody) Statement() Statement {
	return Statement{
		Type: n.Type,
	}
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

// Node fn
func (n EmptyStatement) Node() Node {
	return Node{
		Type: n.Type,
	}
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

// Node fn
func (n DebuggerStatement) Node() Node {
	return Node{
		Type: n.Type,
	}
}

// Statement fn
func (n DebuggerStatement) Statement() Statement {
	return Statement{
		Type: n.Type,
	}
}

// WithStatement struct
type WithStatement struct {
	Type   string      `json:"type,omitempty"`
	Object IExpression `json:"object,omitempty"`
	Body   IStatement  `json:"body,omitempty"`
}

// Node fn
func (n WithStatement) Node() Node {
	return Node{
		Type: n.Type,
	}
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
	Argument IExpression `json:"argument,omitempty"`
}

// Node fn
func (n ReturnStatement) Node() Node {
	return Node{
		Type: n.Type,
	}
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
	Body  IStatement `json:"body,omitempty"`
}

// Node fn
func (n LabeledStatement) Node() Node {
	return Node{
		Type: n.Type,
	}
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

// Node fn
func (n BreakStatement) Node() Node {
	return Node{
		Type: n.Type,
	}
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

// Node fn
func (n ContinueStatement) Node() Node {
	return Node{
		Type: n.Type,
	}
}

// Statement fn
func (n ContinueStatement) Statement() Statement {
	return Statement{
		Type: n.Type,
	}
}

// IfStatement struct
type IfStatement struct {
	Type       string      `json:"type,omitempty"`
	Test       IExpression `json:"test,omitempty"`
	Consequent IStatement  `json:"consequent,omitempty"`
	Alternate  IStatement  `json:"alternate,omitempty"`
}

// Node fn
func (n IfStatement) Node() Node {
	return Node{
		Type: n.Type,
	}
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
	Discriminant IExpression  `json:"discriminant,omitempty"`
	Cases        []SwitchCase `json:"cases,omitempty"`
}

// Node fn
func (n SwitchStatement) Node() Node {
	return Node{
		Type: n.Type,
	}
}

// Statement fn
func (n SwitchStatement) Statement() Statement {
	return Statement{
		Type: n.Type,
	}
}

// SwitchCase struct
type SwitchCase struct {
	Type       string       `json:"type,omitempty"`
	Test       IExpression  `json:"test,omitempty"`
	Consequent []IStatement `json:"consequent,omitempty"`
}

// Node fn
func (n SwitchCase) Node() Node {
	return Node{
		Type: n.Type,
	}
}

// ThrowStatement struct
type ThrowStatement struct {
	Type     string      `json:"type,omitempty"`
	Argument IExpression `json:"argument,omitempty"`
}

// Node fn
func (n ThrowStatement) Node() Node {
	return Node{
		Type: n.Type,
	}
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
	Block     IBlockStatement `json:"block,omitempty"`
	Handler   *CatchClause    `json:"handler,omitempty"`
	Finalizer IBlockStatement `json:"finalizer,omitempty"`
}

// Node fn
func (n TryStatement) Node() Node {
	return Node{
		Type: n.Type,
	}
}

// Statement fn
func (n TryStatement) Statement() Statement {
	return Statement{
		Type: n.Type,
	}
}

// CatchClause struct
type CatchClause struct {
	Type  string          `json:"type,omitempty"`
	Param IPattern        `json:"param,omitempty"`
	Body  IBlockStatement `json:"body,omitempty"`
}

// Node fn
func (n CatchClause) Node() Node {
	return Node{
		Type: n.Type,
	}
}

// WhileStatement struct
type WhileStatement struct {
	Type string      `json:"type,omitempty"`
	Test IExpression `json:"test,omitempty"`
	Body IStatement  `json:"body,omitempty"`
}

// Node fn
func (n WhileStatement) Node() Node {
	return Node{
		Type: n.Type,
	}
}

// Statement fn
func (n WhileStatement) Statement() Statement {
	return Statement{
		Type: n.Type,
	}
}

// DoWhileStatement struct
type DoWhileStatement struct {
	Type string      `json:"type,omitempty"`
	Body IStatement  `json:"body,omitempty"`
	Test IExpression `json:"test,omitempty"`
}

// Node fn
func (n DoWhileStatement) Node() Node {
	return Node{
		Type: n.Type,
	}
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
	Test   IExpression `json:"test,omitempty"`
	Update IExpression `json:"update,omitempty"`
	Body   IStatement  `json:"body,omitempty"`
}

// Node fn
func (n ForStatement) Node() Node {
	return Node{
		Type: n.Type,
	}
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
	Right IExpression `json:"right,omitempty"`
	Body  IStatement  `json:"body,omitempty"`
}

// Node fn
func (n ForInStatement) Node() Node {
	return Node{
		Type: n.Type,
	}
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

// Node fn
func (n Declaration) Node() Node {
	return Node{
		Type: n.Type,
	}
}

// Statement fn
func (n Declaration) Statement() Statement {
	return Statement{
		Type: n.Type,
	}
}

// Declaration fn
func (n Declaration) Declaration() Declaration {
	return n
}

// FunctionDeclaration struct
type FunctionDeclaration struct {
	Type      string      `json:"type,omitempty"`
	ID        *Identifier `json:"id,omitempty"`
	Params    []IPattern
	Body      FunctionBody
	Generator bool
	Async     bool
}

// Node fn
func (n FunctionDeclaration) Node() Node {
	return Node{
		Type: n.Type,
	}
}

// Statement fn
func (n FunctionDeclaration) Statement() Statement {
	return Statement{
		Type: n.Type,
	}
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

// Node fn
func (n VariableDeclaration) Node() Node {
	return Node{
		Type: n.Type,
	}
}

// Statement fn
func (n VariableDeclaration) Statement() Statement {
	return Statement{
		Type: n.Type,
	}
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
	ID   IPattern    `json:"id,omitempty"`
	Init IExpression `json:"init,omitempty"`
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

// Expression fn
func (n Expression) Expression() Expression {
	return n
}

// ThisExpression struct
type ThisExpression struct {
	Type string `json:"type,omitempty"`
}

// Node fn
func (n ThisExpression) Node() Node {
	return Node{
		Type: n.Type,
	}
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
	Elements []IExpression `json:"elements,omitempty"`
}

// Node fn
func (n ArrayExpression) Node() Node {
	return Node{
		Type: n.Type,
	}
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

// Node fn
func (n ObjectExpression) Node() Node {
	return Node{
		Type: n.Type,
	}
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
	Value IExpression `json:"value,omitempty"`
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
	Params []IPattern
	Body   FunctionBody
}

// Node fn
func (n FunctionExpression) Node() Node {
	return Node{
		Type: n.Type,
	}
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
	Argument IExpression   `json:"argument,omitempty"`
}

// Node fn
func (n UnaryExpression) Node() Node {
	return Node{
		Type: n.Type,
	}
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
	Argument IExpression    `json:"argument,omitempty"`
	Prefix   bool           `json:"prefix,omitempty"`
}

// Node fn
func (n UpdateExpression) Node() Node {
	return Node{
		Type: n.Type,
	}
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
	Left     IExpression    `json:"left,omitempty"`
	Right    IExpression    `json:"right,omitempty"`
}

// Node fn
func (n BinaryExpression) Node() Node {
	return Node{
		Type: n.Type,
	}
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
	Right    IExpression        `json:"right,omitempty"`
}

// Node fn
func (n AssignmentExpression) Node() Node {
	return Node{
		Type: n.Type,
	}
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
	Left     IExpression     `json:"left,omitempty"`
	Right    IExpression     `json:"right,omitempty"`
}

// Node fn
func (n LogicalExpression) Node() Node {
	return Node{
		Type: n.Type,
	}
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
	Type     string      `json:"type,omitempty"`
	Object   IExpression `json:"object,omitempty"`
	Property IExpression `json:"property,omitempty"`
	Computed bool        `json:"computed,omitempty"`
}

// Node fn
func (n MemberExpression) Node() Node {
	return Node{
		Type: n.Type,
	}
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
	Type       string      `json:"type,omitempty"`
	Test       IExpression `json:"test,omitempty"`
	Alternate  IExpression `json:"alternate,omitempty"`
	Consequent IExpression `json:"consequent,omitempty"`
}

// Node fn
func (n ConditionalExpression) Node() Node {
	return Node{
		Type: n.Type,
	}
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
	Type      string        `json:"type,omitempty"`
	Callee    IExpression   `json:"callee,omitempty"`
	Arguments []IExpression `json:"arguments,omitempty"`
}

// Node fn
func (n CallExpression) Node() Node {
	return Node{
		Type: n.Type,
	}
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
	Type      string        `json:"type,omitempty"`
	Callee    IExpression   `json:"callee,omitempty"`
	Arguments []IExpression `json:"arguments,omitempty"`
}

// Node fn
func (n NewExpression) Node() Node {
	return Node{
		Type: n.Type,
	}
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
	Type        string        `json:"type,omitempty"`
	Expressions []IExpression `json:"expressions,omitempty"`
}

// Node fn
func (n SequenceExpression) Node() Node {
	return Node{
		Type: n.Type,
	}
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

// Pattern fn
func (n Pattern) Pattern() Pattern {
	return n
}

// AwaitExpression struct
type AwaitExpression struct {
	Type     string      `json:"type,omitempty"`
	Argument IExpression `json:"argument,omitempty"`
}

// Node fn
func (n AwaitExpression) Node() Node {
	return Node{
		Type: n.Type,
	}
}

// Expression fn
func (n AwaitExpression) Expression() Expression {
	return Expression{
		Type: n.Type,
	}
}
