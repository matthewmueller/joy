package htmlhtmlelement

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLHTMLElement struct
// js:"HTMLHTMLElement,omit"
type HTMLHTMLElement struct {
	window.HTMLElement
}

// Version prop Sets or retrieves the DTD version that governs the current document.
func (*HTMLHtmlElement) Version() (version string) {
	js.Rewrite("$<.version")
	return version
}

// Version prop Sets or retrieves the DTD version that governs the current document.
func (*HTMLHtmlElement) SetVersion(version string) {
	js.Rewrite("$<.version = $1", version)
}
