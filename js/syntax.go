package js

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
	Expression `json:"expression,omitempty"`
	Pattern    `json:"pattern,omitempty"`
	Type       string `json:"type,omitempty"`
	Name       string `json:"name,omitempty"`
}

// Literal struct
type Literal struct {
	Expression `json:"expression,omitempty"`
	Type       string      `json:"type,omitempty"`
	Value      interface{} `json:"value,omitempty"` // string | boolean | null | number | RegExp;
}

// RegExpLiteral struct
type RegExpLiteral struct {
	Literal `json:"literal,omitempty"`
	Regex   struct {
		Pattern string `json:"pattern,omitempty"`
		Flags   string `json:"flags,omitempty"`
	} `json:"regex,omitempty"`
}

// Program struct
type Program struct {
	Node `json:"node,omitempty"`
	Type string `json:"type,omitempty"`

	// [ Directive | Statement ]
	Body []interface{} `json:"body,omitempty"`
}

// Function struct
type Function struct {
	Node   `json:"node,omitempty"`
	ID     *Identifier  `json:"id,omitempty"`
	Params []Pattern    `json:"params,omitempty"`
	Body   FunctionBody `json:"body,omitempty"`
}

// Statement struct
type Statement struct {
	Node `json:"node,omitempty"`
}

// ExpressionStatement struct
type ExpressionStatement struct {
	Statement  `json:"statement,omitempty"`
	Type       string     `json:"type,omitempty"`
	Expression Expression `json:"expression,omitempty"`
}

// Directive struct
type Directive struct {
	ExpressionStatement `json:"expression_statement,omitempty"`
	Directive           string `json:"directive,omitempty"`
}

// BlockStatement struct
type BlockStatement struct {
	Statement `json:"statement,omitempty"`
	Type      string      `json:"type,omitempty"`
	Body      []Statement `json:"body,omitempty"`
}

// FunctionBody struct
type FunctionBody struct {
	BlockStatement `json:"block_statement,omitempty"`

	// [ Directive | Statement ]
	Body []interface{} `json:"body,omitempty"`
}

// EmptyStatement struct
type EmptyStatement struct {
	Statement `json:"statement,omitempty"`
	Type      string `json:"type,omitempty"`
}

// DebuggerStatement struct
type DebuggerStatement struct {
	Statement `json:"statement,omitempty"`
	Type      string `json:"type,omitempty"`
}

// With struct
type With struct {
	Statement `json:"statement,omitempty"`
	Object    Expression `json:"object,omitempty"`
	Body      Statement  `json:"body,omitempty"`
}

// ReturnStatement struct
type ReturnStatement struct {
	Statement `json:"statement,omitempty"`
	Type      string      `json:"type,omitempty"`
	Argument  *Expression `json:"argument,omitempty"`
}

// LabeledStatement struct
type LabeledStatement struct {
	Statement `json:"statement,omitempty"`
	Type      string     `json:"type,omitempty"`
	Label     Identifier `json:"label,omitempty"`
	Body      Statement  `json:"body,omitempty"`
}

// BreakStatement struct
type BreakStatement struct {
	Statement `json:"statement,omitempty"`
	Type      string      `json:"type,omitempty"`
	Label     *Identifier `json:"label,omitempty"`
}

// ContinueStatement struct
type ContinueStatement struct {
	Statement `json:"statement,omitempty"`
	Type      string      `json:"type,omitempty"`
	Label     *Identifier `json:"label,omitempty"`
}

// IfStatement struct
type IfStatement struct {
	Statement  `json:"statement,omitempty"`
	Test       Expression `json:"test,omitempty"`
	Consequent Statement  `json:"consequent,omitempty"`
	Alternate  *Statement `json:"alternate,omitempty"`
}

// SwitchStatement struct
type SwitchStatement struct {
	Statement    `json:"statement,omitempty"`
	Type         string       `json:"type,omitempty"`
	Discriminant Expression   `json:"discriminant,omitempty"`
	Cases        []SwitchCase `json:"cases,omitempty"`
}

// SwitchCase struct
type SwitchCase struct {
	Node       `json:"node,omitempty"`
	Type       string      `json:"type,omitempty"`
	Test       *Expression `json:"test,omitempty"`
	Consequent []Statement `json:"consequent,omitempty"`
}

// ThrowStatement struct
type ThrowStatement struct {
	Statement `json:"statement,omitempty"`
	Type      string     `json:"type,omitempty"`
	Argument  Expression `json:"argument,omitempty"`
}

// TryStatement struct
type TryStatement struct {
	Statement `json:"statement,omitempty"`
	Type      string          `json:"type,omitempty"`
	Block     BlockStatement  `json:"block,omitempty"`
	Handler   *CatchClause    `json:"handler,omitempty"`
	Finalizer *BlockStatement `json:"finalizer,omitempty"`
}

// CatchClause struct
type CatchClause struct {
	Node  `json:"node,omitempty"`
	Type  string         `json:"type,omitempty"`
	Param Pattern        `json:"param,omitempty"`
	Body  BlockStatement `json:"body,omitempty"`
}

// WhileStatement struct
type WhileStatement struct {
	Statement `json:"statement,omitempty"`
	Type      string     `json:"type,omitempty"`
	Test      Expression `json:"test,omitempty"`
	Body      Statement  `json:"body,omitempty"`
}

// DoWhileStatement struct
type DoWhileStatement struct {
	Statement `json:"statement,omitempty"`
	Type      string     `json:"type,omitempty"`
	Body      Statement  `json:"body,omitempty"`
	Test      Expression `json:"test,omitempty"`
}

// ForStatement struct
type ForStatement struct {
	Statement `json:"statement,omitempty"`
	Type      string      `json:"type,omitempty"`
	Init      interface{} `json:"init,omitempty"` //  VariableDeclaration | Expression | null
	Test      *Expression `json:"test,omitempty"`
	Update    *Expression `json:"update,omitempty"`
	Body      Statement   `json:"body,omitempty"`
}

// ForInStatement struct
type ForInStatement struct {
	Statement `json:"statement,omitempty"`
	Left      interface{} `json:"left,omitempty"` // VariableDeclaration |  Pattern
	Right     Expression  `json:"right,omitempty"`
	Body      Statement   `json:"body,omitempty"`
}

// Declaration struct
type Declaration struct {
	Statement `json:"statement,omitempty"`
}

// FunctionDeclaration struct
type FunctionDeclaration struct {
	Function    `json:"function,omitempty"`
	Declaration `json:"declaration,omitempty"`
	Type        string     `json:"type,omitempty"`
	ID          Identifier `json:"id,omitempty"`
}

// VariableDeclaration struct
type VariableDeclaration struct {
	Declaration  `json:"declaration,omitempty"`
	Type         string               `json:"type,omitempty"`
	Declarations []VariableDeclarator `json:"declarations,omitempty"`
	Kind         string               `json:"kind,omitempty"` // "var"
}

// VariableDeclarator struct
type VariableDeclarator struct {
	Node `json:"node,omitempty"`
	Type string      `json:"type,omitempty"`
	ID   Pattern     `json:"id,omitempty"`
	Init *Expression `json:"init,omitempty"`
}

// Expression struct
type Expression struct {
	Node `json:"node,omitempty"`
}

// ThisExpression struct
type ThisExpression struct {
	Expression `json:"expression,omitempty"`
	Type       string `json:"type,omitempty"`
}

// ArrayExpression struct
type ArrayExpression struct {
	Expression `json:"expression,omitempty"`
	Type       string        `json:"type,omitempty"`
	Elements   []*Expression `json:"elements,omitempty"`
}

// ObjectExpression struct
type ObjectExpression struct {
	Expression `json:"expression,omitempty"`
	Type       string     `json:"type,omitempty"`
	Properties []Property `json:"properties,omitempty"`
}

// Property struct
type Property struct {
	Node  `json:"node,omitempty"`
	Type  string      `json:"type,omitempty"`
	Key   interface{} `json:"key,omitempty"` // Literal | Identifier
	Value Expression  `json:"value,omitempty"`
	Kind  string      `json:"kind,omitempty"` // "init" | "get" | "set"
}

// FunctionExpression struct
type FunctionExpression struct {
	Function   `json:"function,omitempty"`
	Expression `json:"expression,omitempty"`
	Type       string `json:"type,omitempty"`
}

// UnaryExpression struct
type UnaryExpression struct {
	Expression `json:"expression,omitempty"`
	Type       string        `json:"type,omitempty"`
	Operator   UnaryOperator `json:"operator,omitempty"`
	Prefix     bool          `json:"prefix,omitempty"`
	Argument   Expression    `json:"argument,omitempty"`
}

// UnaryOperator enum
// "-" | "+" | "!" | "~" | "typeof" | "void" | "delete"
type UnaryOperator string

// UpdateExpression struct
type UpdateExpression struct {
	Expression `json:"expression,omitempty"`
	Type       string         `json:"type,omitempty"`
	Operator   UpdateOperator `json:"operator,omitempty"`
	Argument   Expression     `json:"argument,omitempty"`
	Prefix     bool           `json:"prefix,omitempty"`
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
	Expression `json:"expression,omitempty"`
	Type       string             `json:"type,omitempty"`
	Operator   AssignmentOperator `json:"operator,omitempty"`
	Left       interface{}        `json:"left,omitempty"` // Pattern | Expression
	Right      Expression         `json:"right,omitempty"`
}

// AssignmentOperator enum
// 	"=" | "+=" | "-=" | "*=" | "/=" | "%="
// | "<<=" | ">>=" | ">>>="
// | "|=" | "^=" | "&="
type AssignmentOperator string

// LogicalExpression struct
type LogicalExpression struct {
	Expression `json:"expression,omitempty"`
	Type       string          `json:"type,omitempty"`
	Operator   LogicalOperator `json:"operator,omitempty"`
	Left       Expression      `json:"left,omitempty"`
	Right      Expression      `json:"right,omitempty"`
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
	Expression `json:"expression,omitempty"`
	Pattern    `json:"pattern,omitempty"`
	Type       string     `json:"type,omitempty"`
	Object     Expression `json:"object,omitempty"`
	Property   Expression `json:"property,omitempty"`
	Computed   bool       `json:"computed,omitempty"`
}

// ConditionalExpression struct
//
// A conditional expression, i.e., a ternary ?/: expression.
type ConditionalExpression struct {
	Expression `json:"expression,omitempty"`
	Type       string     `json:"type,omitempty"`
	Test       Expression `json:"test,omitempty"`
	Alternate  Expression `json:"alternate,omitempty"`
	Consequent Expression `json:"consequent,omitempty"`
}

// CallExpression struct
//
// A function or method call expression.
type CallExpression struct {
	Expression `json:"expression,omitempty"`
	Type       string       `json:"type,omitempty"`
	Callee     Expression   `json:"callee,omitempty"`
	Arguments  []Expression `json:"arguments,omitempty"`
}

// NewExpression struct
//
// A `new` expression.
type NewExpression struct {
	Expression `json:"expression,omitempty"`
	Type       string       `json:"type,omitempty"`
	Callee     Expression   `json:"callee,omitempty"`
	Arguments  []Expression `json:"arguments,omitempty"`
}

// SequenceExpression struct
//
// A sequence expression, i.e., a comma-separated sequence of expressions.
type SequenceExpression struct {
	Type        string       `json:"type,omitempty"`
	Expressions []Expression `json:"expressions,omitempty"`
}

// Pattern struct
//
// Destructuring binding and assignment are not part of ES5,
// but all binding positions accept Pattern to allow for destructuring
// in ES6. Nevertheless, for ES5, the only Pattern subtype is Identifier.
type Pattern struct {
	Node `json:"node,omitempty"`
}
