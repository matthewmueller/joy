package cssimportrule

import (
	"github.com/matthewmueller/golly/dom/medialist"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

var _ window.CSSRule = (*CSSImportRule)(nil)

// CSSImportRule struct
// js:"CSSImportRule,omit"
type CSSImportRule struct {
}

// Href prop
// js:"href"
func (*CSSImportRule) Href() (href string) {
	js.Rewrite("$_.href")
	return href
}

// Media prop
// js:"media"
func (*CSSImportRule) Media() (media *medialist.MediaList) {
	js.Rewrite("$_.media")
	return media
}

// StyleSheet prop
// js:"styleSheet"
func (*CSSImportRule) StyleSheet() (styleSheet *window.CSSStyleSheet) {
	js.Rewrite("$_.styleSheet")
	return styleSheet
}

// CSSText prop
// js:"cssText"
func (*CSSImportRule) CSSText() (cssText string) {
	js.Rewrite("$_.cssText")
	return cssText
}

// SetCSSText prop
// js:"cssText"
func (*CSSImportRule) SetCSSText(cssText string) {
	js.Rewrite("$_.cssText = $1", cssText)
}

// ParentRule prop
// js:"parentRule"
func (*CSSImportRule) ParentRule() (parentRule window.CSSRule) {
	js.Rewrite("$_.parentRule")
	return parentRule
}

// ParentStyleSheet prop
// js:"parentStyleSheet"
func (*CSSImportRule) ParentStyleSheet() (parentStyleSheet *window.CSSStyleSheet) {
	js.Rewrite("$_.parentStyleSheet")
	return parentStyleSheet
}

// Type prop
// js:"type"
func (*CSSImportRule) Type() (kind uint8) {
	js.Rewrite("$_.type")
	return kind
}
