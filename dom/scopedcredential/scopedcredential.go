package scopedcredential

import (
	"github.com/matthewmueller/golly/dom2/scopedcredentialtype"
	"github.com/matthewmueller/golly/js"
)

// ScopedCredential struct
// js:"ScopedCredential,omit"
type ScopedCredential struct {
}

// ID prop
func (*ScopedCredential) ID() (id []byte) {
	js.Rewrite("$<.id")
	return id
}

// Type prop
func (*ScopedCredential) Type() (kind *scopedcredentialtype.ScopedCredentialType) {
	js.Rewrite("$<.type")
	return kind
}
