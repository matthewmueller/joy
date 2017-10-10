package promise

// Promise struct
// js:"promise,global"
type Promise struct {
	value interface{}
	err   error
}

// Resolve the promise
// js:"resolve"
func Resolve(v interface{}) *Promise {
	return &Promise{value: v}
}

// Reject the promise
// js:"reject"
func Reject(err error) *Promise {
	return &Promise{err: err}
}

// Then fn
// js:"then"
func (p *Promise) Then(func(v interface{}) interface{}) *Promise {
	return p
}

// Catch fn
// js:"catch"
func (p *Promise) Catch(func(err error) interface{}) *Promise {
	return p
}
