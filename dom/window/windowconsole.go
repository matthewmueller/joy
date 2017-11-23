package window

import "github.com/matthewmueller/golly/js"

// WindowConsole struct
// js:"WindowConsole,omit"
type WindowConsole struct {
}

// Console prop
func (*WindowConsole) Console() (console *Console) {
	js.Rewrite("$<.console")
	return console
}
