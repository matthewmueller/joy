package msfidosignatureassertion

import (
	"github.com/matthewmueller/golly/dom/msassertion"
	"github.com/matthewmueller/golly/dom/msfidosignature"
	"github.com/matthewmueller/golly/js"
)

// MSFIDOSignatureAssertion struct
// js:"MSFIDOSignatureAssertion,omit"
type MSFIDOSignatureAssertion struct {
	msassertion.MSAssertion
}

// Signature prop
func (*MSFIDOSignatureAssertion) Signature() (signature *msfidosignature.MSFIDOSignature) {
	js.Rewrite("$<.signature")
	return signature
}
