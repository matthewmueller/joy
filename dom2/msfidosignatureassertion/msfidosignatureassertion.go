package msfidosignatureassertion

import (
	"github.com/matthewmueller/golly/dom2/msassertion"
	"github.com/matthewmueller/golly/dom2/msfidosignature"
	"github.com/matthewmueller/golly/js"
)

// js:"MSFIDOSignatureAssertion,omit"
type MSFIDOSignatureAssertion struct {
	msassertion.MSAssertion
}

// Signature
func (*MSFIDOSignatureAssertion) Signature() (signature *msfidosignature.MSFIDOSignature) {
	js.Rewrite("$<.Signature")
	return signature
}
