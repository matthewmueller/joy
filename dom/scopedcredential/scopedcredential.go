package scopedcredential

import (
	"github.com/matthewmueller/joy/dom/scopedcredentialtype"
	"github.com/matthewmueller/joy/js"
)

// ScopedCredential struct
// js:"ScopedCredential,omit"
type ScopedCredential struct {
}

// ID prop
// js:"id"
func (*ScopedCredential) ID() (id []byte) {
	js.Rewrite("$_.id")
	return id
}

// Type prop
// js:"type"
func (*ScopedCredential) Type() (kind *scopedcredentialtype.ScopedCredentialType) {
	js.Rewrite("$_.type")
	return kind
}
