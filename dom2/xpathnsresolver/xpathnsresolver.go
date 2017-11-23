package xpathnsresolver

import "github.com/matthewmueller/golly/js"

// XPathNSResolver struct
// js:"XPathNSResolver,omit"
type XPathNSResolver struct {
}

// LookupNamespaceURI
func (*XPathNSResolver) LookupNamespaceURI(prefix string) (s string) {
	js.Rewrite("$<.LookupNamespaceURI($1)", prefix)
	return s
}
