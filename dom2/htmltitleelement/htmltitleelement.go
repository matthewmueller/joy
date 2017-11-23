package htmltitleelement

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLTitleElement struct
// js:"HTMLTitleElement,omit"
type HTMLTitleElement struct {
	window.HTMLElement
}

// Text Retrieves or sets the text of the object as a string.
func (*HTMLTitleElement) Text() (text string) {
	js.Rewrite("$<.Text")
	return text
}

// Text Retrieves or sets the text of the object as a string.
func (*HTMLTitleElement) SetText(text string) {
	js.Rewrite("$<.Text = $1", text)
}
