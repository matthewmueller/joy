package js

import (
	"github.com/matthewmueller/golly/jsast"
)

// Statement interface
type Statement interface {
	Statement() (jsast.IStatement, error)
	String() string
}

// Expression interface
type Expression interface {
	String() string
}

// Raw statement
func Raw(src string) {
	jsast.Parse(src)
}

// Promises struct
type Promises struct {
	expr   string
	result interface{}
	err    error
}

// Promise fn
func Promise(expr string) *Promises {
	return &Promises{
		expr: expr,
	}
}

// Then fn
func (p *Promises) Then(result interface{}) *Promises {
	p.result = result
	return p
}

// Catch fn
func (p *Promises) Catch(err error) *Promises {
	p.err = err
	return p
}

// Statement fn
func (p *Promises) Statement() (jsast.IStatement, error) {
	return jsast.CreateExpressionStatement(
		jsast.CreateIdentifier("PROMISE"),
	), nil
	// return nil, nil
}

// String fn
func (p *Promises) String() string {
	return ""
}
