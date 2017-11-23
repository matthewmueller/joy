package windowtimersextension

import "github.com/matthewmueller/golly/js"

// WindowTimersExtension struct
// js:"WindowTimersExtension,omit"
type WindowTimersExtension struct {
}

// ClearImmediate fn
func (*WindowTimersExtension) ClearImmediate(handle int) {
	js.Rewrite("$<.clearImmediate($1)", handle)
}

// SetImmediate fn
func (*WindowTimersExtension) SetImmediate(expression interface{}, args interface{}) (i int) {
	js.Rewrite("$<.setImmediate($1, $2)", expression, args)
	return i
}
