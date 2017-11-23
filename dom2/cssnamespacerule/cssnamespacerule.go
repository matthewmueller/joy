package cssnamespacerule

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"CSSNamespaceRule,omit"
type CSSNamespaceRule struct {
	window.CSSRule
}

// NamespaceURI
func (*CSSNamespaceRule) NamespaceURI() (namespaceURI string) {
	js.Rewrite("$<.NamespaceURI")
	return namespaceURI
}

// Prefix
func (*CSSNamespaceRule) Prefix() (prefix string) {
	js.Rewrite("$<.Prefix")
	return prefix
}
