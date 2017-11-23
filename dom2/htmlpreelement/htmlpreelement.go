package htmlpreelement

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"HTMLPreElement,omit"
type HTMLPreElement struct {
	window.HTMLElement
}

// Width Sets or gets a value that you can use to implement your own width functionality for the object.
func (*HTMLPreElement) Width() (width int) {
	js.Rewrite("$<.Width")
	return width
}

// Width Sets or gets a value that you can use to implement your own width functionality for the object.
func (*HTMLPreElement) SetWidth(width int) {
	js.Rewrite("$<.Width = $1", width)
}
