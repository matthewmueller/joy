package cssimportrule

import "github.com/matthewmueller/golly/js"

// js:"CSSImportRule,omit"
type CSSImportRule struct {
	window.CSSRule
}

// Href
func (*CSSImportRule) Href() (href string) {
	js.Rewrite("$<.Href")
	return href
}

// Media
func (*CSSImportRule) Media() (media *medialist.MediaList) {
	js.Rewrite("$<.Media")
	return media
}

// StyleSheet
func (*CSSImportRule) StyleSheet() (styleSheet *window.CSSStyleSheet) {
	js.Rewrite("$<.StyleSheet")
	return styleSheet
}
