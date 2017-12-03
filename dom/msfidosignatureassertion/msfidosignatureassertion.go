package msfidosignatureassertion

import (
	"github.com/matthewmueller/joy/dom/msassertion"
	"github.com/matthewmueller/joy/dom/mscredentialtype"
	"github.com/matthewmueller/joy/dom/msfidosignature"
	"github.com/matthewmueller/joy/macro"
)

var _ msassertion.MSAssertion = (*MSFIDOSignatureAssertion)(nil)

// MSFIDOSignatureAssertion struct
// js:"MSFIDOSignatureAssertion,omit"
type MSFIDOSignatureAssertion struct {
}

// Signature prop
// js:"signature"
func (*MSFIDOSignatureAssertion) Signature() (signature *msfidosignature.MSFIDOSignature) {
	macro.Rewrite("$_.signature")
	return signature
}

// ID prop
// js:"id"
func (*MSFIDOSignatureAssertion) ID() (id string) {
	macro.Rewrite("$_.id")
	return id
}

// Type prop
// js:"type"
func (*MSFIDOSignatureAssertion) Type() (kind *mscredentialtype.MSCredentialType) {
	macro.Rewrite("$_.type")
	return kind
}
