package windowconsole

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/console"
	"github.com/matthewmueller/golly/js"
)

type WindowConsole struct {
}

func (*WindowConsole) GetConsole() (console *console.Console) {
	js.Rewrite("$<.console")
	return console
}
