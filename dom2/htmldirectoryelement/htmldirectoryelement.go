package htmldirectoryelement

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"HTMLDirectoryElement,omit"
type HTMLDirectoryElement struct {
	window.HTMLElement
}

// Compact
func (*HTMLDirectoryElement) Compact() (compact bool) {
	js.Rewrite("$<.Compact")
	return compact
}

// Compact
func (*HTMLDirectoryElement) SetCompact(compact bool) {
	js.Rewrite("$<.Compact = $1", compact)
}
