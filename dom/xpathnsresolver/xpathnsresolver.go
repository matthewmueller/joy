package xpathnsresolver

import "github.com/matthewmueller/golly/js"

// XPathNSResolver struct
// js:"XPathNSResolver,omit"
type XPathNSResolver struct {
}

// LookupNamespaceURI fn
// js:"lookupNamespaceURI"
func (*XPathNSResolver) LookupNamespaceURI(prefix string) (s string) {
	js.Rewrite("$_.lookupNamespaceURI($1)", prefix)
	return s
}
