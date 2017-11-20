package windowtimersextension

import "github.com/matthewmueller/golly/js"

type WindowTimersExtension struct {
}

func (*WindowTimersExtension) ClearImmediate(handle int) {
	js.Rewrite("$<.clearImmediate($1)", handle)
}

func (*WindowTimersExtension) SetImmediate(expression interface{}, args interface{}) (i int) {
	js.Rewrite("$<.setImmediate($1, $2)", expression, args)
	return i
}
