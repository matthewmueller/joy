package htmlbaseelement

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"HTMLBaseElement,omit"
type HTMLBaseElement struct {
	window.HTMLElement
}

// Href Gets or sets the baseline URL on which relative links are based.
func (*HTMLBaseElement) Href() (href string) {
	js.Rewrite("$<.Href")
	return href
}

// Href Gets or sets the baseline URL on which relative links are based.
func (*HTMLBaseElement) SetHref(href string) {
	js.Rewrite("$<.Href = $1", href)
}

// Target Sets or retrieves the window or frame at which to target content.
func (*HTMLBaseElement) Target() (target string) {
	js.Rewrite("$<.Target")
	return target
}

// Target Sets or retrieves the window or frame at which to target content.
func (*HTMLBaseElement) SetTarget(target string) {
	js.Rewrite("$<.Target = $1", target)
}
