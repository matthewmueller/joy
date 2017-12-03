package cssimportrule

import (
	"github.com/matthewmueller/joy/dom/medialist"
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/macro"
)

var _ window.CSSRule = (*CSSImportRule)(nil)

// CSSImportRule struct
// js:"CSSImportRule,omit"
type CSSImportRule struct {
}

// Href prop
// js:"href"
func (*CSSImportRule) Href() (href string) {
	macro.Rewrite("$_.href")
	return href
}

// Media prop
// js:"media"
func (*CSSImportRule) Media() (media *medialist.MediaList) {
	macro.Rewrite("$_.media")
	return media
}

// StyleSheet prop
// js:"styleSheet"
func (*CSSImportRule) StyleSheet() (styleSheet *window.CSSStyleSheet) {
	macro.Rewrite("$_.styleSheet")
	return styleSheet
}

// CSSText prop
// js:"cssText"
func (*CSSImportRule) CSSText() (cssText string) {
	macro.Rewrite("$_.cssText")
	return cssText
}

// SetCSSText prop
// js:"cssText"
func (*CSSImportRule) SetCSSText(cssText string) {
	macro.Rewrite("$_.cssText = $1", cssText)
}

// ParentRule prop
// js:"parentRule"
func (*CSSImportRule) ParentRule() (parentRule window.CSSRule) {
	macro.Rewrite("$_.parentRule")
	return parentRule
}

// ParentStyleSheet prop
// js:"parentStyleSheet"
func (*CSSImportRule) ParentStyleSheet() (parentStyleSheet *window.CSSStyleSheet) {
	macro.Rewrite("$_.parentStyleSheet")
	return parentStyleSheet
}

// Type prop
// js:"type"
func (*CSSImportRule) Type() (kind uint8) {
	macro.Rewrite("$_.type")
	return kind
}
