package cssimportrule

import (
	"github.com/matthewmueller/golly/dom/medialist"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// CSSImportRule struct
// js:"CSSImportRule,omit"
type CSSImportRule struct {
	window.CSSRule
}

// Href prop
func (*CSSImportRule) Href() (href string) {
	js.Rewrite("$<.href")
	return href
}

// Media prop
func (*CSSImportRule) Media() (media *medialist.MediaList) {
	js.Rewrite("$<.media")
	return media
}

// StyleSheet prop
func (*CSSImportRule) StyleSheet() (styleSheet *window.CSSStyleSheet) {
	js.Rewrite("$<.styleSheet")
	return styleSheet
}
