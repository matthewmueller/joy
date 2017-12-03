package xpathnsresolver

import "github.com/matthewmueller/joy/macro"

// XPathNSResolver struct
// js:"XPathNSResolver,omit"
type XPathNSResolver struct {
}

// LookupNamespaceURI fn
// js:"lookupNamespaceURI"
func (*XPathNSResolver) LookupNamespaceURI(prefix string) (s string) {
	macro.Rewrite("$_.lookupNamespaceURI($1)", prefix)
	return s
}
