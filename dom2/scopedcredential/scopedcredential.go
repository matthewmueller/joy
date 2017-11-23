package scopedcredential

import (
	"github.com/matthewmueller/golly/dom2/scopedcredentialtype"
	"github.com/matthewmueller/golly/js"
)

// js:"ScopedCredential,omit"
type ScopedCredential struct {
}

// ID
func (*ScopedCredential) ID() (id []byte) {
	js.Rewrite("$<.ID")
	return id
}

// Type
func (*ScopedCredential) Type() (kind *scopedcredentialtype.ScopedCredentialType) {
	js.Rewrite("$<.Type")
	return kind
}
