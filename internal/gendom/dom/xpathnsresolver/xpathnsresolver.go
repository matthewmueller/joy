package xpathnsresolver

import "github.com/matthewmueller/golly/js"

type XPathNSResolver struct {
}

func (*XPathNSResolver) LookupNamespaceURI(prefix string) (s string) {
	js.Rewrite("$<.lookupNamespaceURI($1)", prefix)
	return s
}
