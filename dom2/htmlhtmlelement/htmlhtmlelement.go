package htmlhtmlelement

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"HTMLHTMLElement,omit"
type HTMLHTMLElement struct {
	window.HTMLElement
}

// Version Sets or retrieves the DTD version that governs the current document.
func (*HTMLHtmlElement) Version() (version string) {
	js.Rewrite("$<.Version")
	return version
}

// Version Sets or retrieves the DTD version that governs the current document.
func (*HTMLHtmlElement) SetVersion(version string) {
	js.Rewrite("$<.Version = $1", version)
}
