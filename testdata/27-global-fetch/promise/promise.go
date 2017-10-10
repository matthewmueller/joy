package promise

// Promise struct
type Promise struct {
	value interface{}
	err   error
}

// Resolve the promise
func Resolve(v interface{}) *Promise {
	return &Promise{value: v}
}

// Reject the promise
func Reject(err error) *Promise {
	return &Promise{err: err}
}

// Then fn
func (p *Promise) Then(func(v interface{}) interface{}) *Promise {
	return p
}

// Catch fn
func (p *Promise) Catch(func(err error) interface{}) *Promise {
	return p
}
