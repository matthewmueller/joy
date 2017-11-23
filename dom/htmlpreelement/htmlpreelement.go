package htmlpreelement

import (
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLPreElement struct
// js:"HTMLPreElement,omit"
type HTMLPreElement struct {
	window.HTMLElement
}

// Width prop Sets or gets a value that you can use to implement your own width functionality for the object.
func (*HTMLPreElement) Width() (width int) {
	js.Rewrite("$<.width")
	return width
}

// Width prop Sets or gets a value that you can use to implement your own width functionality for the object.
func (*HTMLPreElement) SetWidth(width int) {
	js.Rewrite("$<.width = $1", width)
}
