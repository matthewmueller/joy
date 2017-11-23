package window

import "github.com/matthewmueller/golly/js"

// WindowConsole struct
// js:"WindowConsole,omit"
type WindowConsole struct {
}

// Console
func (*WindowConsole) Console() (console *Console) {
	js.Rewrite("$<.Console")
	return console
}
