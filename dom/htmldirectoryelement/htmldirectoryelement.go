package htmldirectoryelement

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLDirectoryElement struct
// js:"HTMLDirectoryElement,omit"
type HTMLDirectoryElement struct {
	window.HTMLElement
}

// Compact prop
func (*HTMLDirectoryElement) Compact() (compact bool) {
	js.Rewrite("$<.compact")
	return compact
}

// Compact prop
func (*HTMLDirectoryElement) SetCompact(compact bool) {
	js.Rewrite("$<.compact = $1", compact)
}
