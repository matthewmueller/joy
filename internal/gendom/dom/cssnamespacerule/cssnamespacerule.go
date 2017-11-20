package cssnamespacerule

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/cssrule"
	"github.com/matthewmueller/golly/js"
)

type CSSNamespaceRule struct {
	*cssrule.CSSRule
}

func (*CSSNamespaceRule) GetNamespaceURI() (namespaceURI string) {
	js.Rewrite("$<.namespaceURI")
	return namespaceURI
}

func (*CSSNamespaceRule) GetPrefix() (prefix string) {
	js.Rewrite("$<.prefix")
	return prefix
}
