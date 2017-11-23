package htmlbaseelement

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// HTMLBaseElement struct
// js:"HTMLBaseElement,omit"
type HTMLBaseElement struct {
	window.HTMLElement
}

// Href prop Gets or sets the baseline URL on which relative links are based.
func (*HTMLBaseElement) Href() (href string) {
	js.Rewrite("$<.href")
	return href
}

// Href prop Gets or sets the baseline URL on which relative links are based.
func (*HTMLBaseElement) SetHref(href string) {
	js.Rewrite("$<.href = $1", href)
}

// Target prop Sets or retrieves the window or frame at which to target content.
func (*HTMLBaseElement) Target() (target string) {
	js.Rewrite("$<.target")
	return target
}

// Target prop Sets or retrieves the window or frame at which to target content.
func (*HTMLBaseElement) SetTarget(target string) {
	js.Rewrite("$<.target = $1", target)
}
