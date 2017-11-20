package windowtimers

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/windowtimersextension"
	"github.com/matthewmueller/golly/js"
)

type WindowTimers struct {
	*windowtimersextension.WindowTimersExtension
}

func (*WindowTimers) ClearInterval(handle int) {
	js.Rewrite("$<.clearInterval($1)", handle)
}

func (*WindowTimers) ClearTimeout(handle int) {
	js.Rewrite("$<.clearTimeout($1)", handle)
}

func (*WindowTimers) SetInterval(handler interface{}, timeout interface{}, args interface{}) (i int) {
	js.Rewrite("$<.setInterval($1, $2, $3)", handler, timeout, args)
	return i
}

func (*WindowTimers) SetTimeout(handler interface{}, timeout interface{}, args interface{}) (i int) {
	js.Rewrite("$<.setTimeout($1, $2, $3)", handler, timeout, args)
	return i
}
