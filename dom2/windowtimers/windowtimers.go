package windowtimers

import "github.com/matthewmueller/golly/js"

// WindowTimers struct
// js:"WindowTimers,omit"
type WindowTimers struct {
}

// ClearImmediate
func (*WindowTimers) ClearImmediate(handle int) {
	js.Rewrite("$<.ClearImmediate($1)", handle)
}

// SetImmediate
func (*WindowTimers) SetImmediate(expression interface{}, args interface{}) (i int) {
	js.Rewrite("$<.SetImmediate($1, $2)", expression, args)
	return i
}

// ClearInterval
func (*WindowTimers) ClearInterval(handle int) {
	js.Rewrite("$<.ClearInterval($1)", handle)
}

// ClearTimeout
func (*WindowTimers) ClearTimeout(handle int) {
	js.Rewrite("$<.ClearTimeout($1)", handle)
}

// SetInterval
func (*WindowTimers) SetInterval(handler interface{}, timeout *interface{}, args interface{}) (i int) {
	js.Rewrite("$<.SetInterval($1, $2, $3)", handler, timeout, args)
	return i
}

// SetTimeout
func (*WindowTimers) SetTimeout(handler interface{}, timeout *interface{}, args interface{}) (i int) {
	js.Rewrite("$<.SetTimeout($1, $2, $3)", handler, timeout, args)
	return i
}
