package cssnamespacerule

import (
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/js"
)

var _ window.CSSRule = (*CSSNamespaceRule)(nil)

// CSSNamespaceRule struct
// js:"CSSNamespaceRule,omit"
type CSSNamespaceRule struct {
}

// NamespaceURI prop
// js:"namespaceURI"
func (*CSSNamespaceRule) NamespaceURI() (namespaceURI string) {
	js.Rewrite("$_.namespaceURI")
	return namespaceURI
}

// Prefix prop
// js:"prefix"
func (*CSSNamespaceRule) Prefix() (prefix string) {
	js.Rewrite("$_.prefix")
	return prefix
}

// CSSText prop
// js:"cssText"
func (*CSSNamespaceRule) CSSText() (cssText string) {
	js.Rewrite("$_.cssText")
	return cssText
}

// SetCSSText prop
// js:"cssText"
func (*CSSNamespaceRule) SetCSSText(cssText string) {
	js.Rewrite("$_.cssText = $1", cssText)
}

// ParentRule prop
// js:"parentRule"
func (*CSSNamespaceRule) ParentRule() (parentRule window.CSSRule) {
	js.Rewrite("$_.parentRule")
	return parentRule
}

// ParentStyleSheet prop
// js:"parentStyleSheet"
func (*CSSNamespaceRule) ParentStyleSheet() (parentStyleSheet *window.CSSStyleSheet) {
	js.Rewrite("$_.parentStyleSheet")
	return parentStyleSheet
}

// Type prop
// js:"type"
func (*CSSNamespaceRule) Type() (kind uint8) {
	js.Rewrite("$_.type")
	return kind
}
