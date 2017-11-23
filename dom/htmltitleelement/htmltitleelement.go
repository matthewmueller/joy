package htmltitleelement

import (
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLTitleElement struct
// js:"HTMLTitleElement,omit"
type HTMLTitleElement struct {
	window.HTMLElement
}

// Text prop Retrieves or sets the text of the object as a string.
func (*HTMLTitleElement) Text() (text string) {
	js.Rewrite("$<.text")
	return text
}

// Text prop Retrieves or sets the text of the object as a string.
func (*HTMLTitleElement) SetText(text string) {
	js.Rewrite("$<.text = $1", text)
}
