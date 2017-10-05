package js

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

// // String fn
// func (p *Promises) String() string {
// 	return ""
// }
