package windowtimersextension

import "github.com/matthewmueller/golly/js"

// WindowTimersExtension struct
// js:"WindowTimersExtension,omit"
type WindowTimersExtension struct {
}

// ClearImmediate
func (*WindowTimersExtension) ClearImmediate(handle int) {
	js.Rewrite("$<.ClearImmediate($1)", handle)
}

// SetImmediate
func (*WindowTimersExtension) SetImmediate(expression interface{}, args interface{}) (i int) {
	js.Rewrite("$<.SetImmediate($1, $2)", expression, args)
	return i
}
