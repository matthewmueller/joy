package cssnamespacerule

import (
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// CSSNamespaceRule struct
// js:"CSSNamespaceRule,omit"
type CSSNamespaceRule struct {
	window.CSSRule
}

// NamespaceURI prop
func (*CSSNamespaceRule) NamespaceURI() (namespaceURI string) {
	js.Rewrite("$<.namespaceURI")
	return namespaceURI
}

// Prefix prop
func (*CSSNamespaceRule) Prefix() (prefix string) {
	js.Rewrite("$<.prefix")
	return prefix
}
