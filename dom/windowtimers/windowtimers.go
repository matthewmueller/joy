package windowtimers

import "github.com/matthewmueller/golly/js"

// WindowTimers struct
// js:"WindowTimers,omit"
type WindowTimers struct {
}

// ClearImmediate fn
func (*WindowTimers) ClearImmediate(handle int) {
	js.Rewrite("$<.clearImmediate($1)", handle)
}

// SetImmediate fn
func (*WindowTimers) SetImmediate(expression interface{}, args interface{}) (i int) {
	js.Rewrite("$<.setImmediate($1, $2)", expression, args)
	return i
}

// ClearInterval fn
func (*WindowTimers) ClearInterval(handle int) {
	js.Rewrite("$<.clearInterval($1)", handle)
}

// ClearTimeout fn
func (*WindowTimers) ClearTimeout(handle int) {
	js.Rewrite("$<.clearTimeout($1)", handle)
}

// SetInterval fn
func (*WindowTimers) SetInterval(handler interface{}, timeout *interface{}, args interface{}) (i int) {
	js.Rewrite("$<.setInterval($1, $2, $3)", handler, timeout, args)
	return i
}

// SetTimeout fn
func (*WindowTimers) SetTimeout(handler interface{}, timeout *interface{}, args interface{}) (i int) {
	js.Rewrite("$<.setTimeout($1, $2, $3)", handler, timeout, args)
	return i
}
