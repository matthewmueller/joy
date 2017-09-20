package js

import (
	"errors"
	"strings"
)

// Stringer interface
type Stringer interface {
	String() (string, error)
}

// Generate JS from the AST
func Generate(node interface{}) (string, error) {
	switch t := node.(type) {
	case Program:
		return t.String()
	}

	return "", errors.New("ast must start with a program")
}

func (p *Program) String() (string, error) {
	results := []string{}
	body := p.Body
	l := len(body)

	for i := 0; i < l; i++ {
		child := body[i]
		switch child.(type) {
		// statements
		case ExpressionStatement, BlockStatement, EmptyStatement,
			DebuggerStatement, WithStatement, ReturnStatement,
			LabeledStatement, BreakStatement, ContinueStatement,
			IfStatement, SwitchStatement, ThrowStatement, TryStatement,
			WhileStatement, DoWhileStatement, ForStatement, ForInStatement:
		// directives
		case Directive:
		default:
			return "", errors.New("a program's body must contain either directives or statements")
		}

		// cast to stringer
		stringer, ok := child.(Stringer)
		if !ok {
			return "", errors.New("child does not implement stringer")
		}

		// string
		s, e := stringer.String()
		if e != nil {
			return "", e
		}
		results = append(results, s)
	}

	return strings.Join(results, "\n"), nil
}

func (n ExpressionStatement) String() (string, error) {
	return "ExpressionStatement", nil
}
func (n BlockStatement) String() (string, error) {
	return "BlockStatement", nil
}

func (n EmptyStatement) String() (string, error) {
	return "EmptyStatement", nil
}

func (n DebuggerStatement) String() (string, error) {
	return "DebuggerStatement", nil
}

func (n WithStatement) String() (string, error) {
	return "WithStatement", nil
}

func (n ReturnStatement) String() (string, error) {
	return "ReturnStatement", nil
}

func (n LabeledStatement) String() (string, error) {
	return "LabeledStatement", nil
}

func (n BreakStatement) String() (string, error) {
	return "BreakStatement", nil
}

func (n ContinueStatement) String() (string, error) {
	return "ContinueStatement", nil
}

func (n IfStatement) String() (string, error) {
	return "IfStatement", nil
}

func (n SwitchStatement) String() (string, error) {
	return "SwitchStatement", nil
}

func (n ThrowStatement) String() (string, error) {
	return "ThrowStatement", nil
}

func (n TryStatement) String() (string, error) {
	return "TryStatement", nil
}

func (n WhileStatement) String() (string, error) {
	return "WhileStatement", nil
}

func (n DoWhileStatement) String() (string, error) {
	return "DoWhileStatement", nil
}

func (n ForStatement) String() (string, error) {
	return "ForStatement", nil
}

func (n ForInStatement) String() (string, error) {
	return "ForInStatement", nil
}

func (n Declaration) String() (string, error) {
	return "Declaration", nil
}

// func generateProgram(node Program) (string, error) {
// 	results := []string{}
// 	body := node.Body
// 	l := len(body)

// 	for i := 0; i < l; i++ {
// 		child := body[i]

// 		t, ok := child.Statement
// 		if !ok {
// 			log.Infof("not ok")
// 		} else {
// 			log.Infof("ok!")
// 		}

// 		switch t := child.(type) {
// 		case Directive:
// 			s, e := generateDirective(t)
// 			if e != nil {
// 				return "", e
// 			}
// 			results = append(results, s)
// 		case Statement:
// 			s, e := generateStatement(t)
// 			if e != nil {
// 				return "", e
// 			}
// 			results = append(results, s)
// 		default:
// 			return "", errors.New("a program must contain either directive or statements")
// 		}
// 	}

// 	return strings.Join(results, "\n"), nil
// }

// func generateDirective(node Directive) (string, error) {
// 	return "", nil
// }

// func generateStatement(node Statement) (string, error) {
// 	return "", nil
// }
