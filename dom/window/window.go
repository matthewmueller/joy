package window

import (
	"github.com/matthewmueller/golly/dom"
	"github.com/matthewmueller/golly/js"
)

func New() *dom.Window {
	js.Rewrite("window")
	return &dom.Window{}
}
