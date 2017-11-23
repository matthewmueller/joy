package xpathnsresolver

import "github.com/matthewmueller/golly/js"

// XPathNSResolver struct
// js:"XPathNSResolver,omit"
type XPathNSResolver struct {
}

// LookupNamespaceURI fn
func (*XPathNSResolver) LookupNamespaceURI(prefix string) (s string) {
	js.Rewrite("$<.lookupNamespaceURI($1)", prefix)
	return s
}
