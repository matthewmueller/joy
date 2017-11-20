package msfidosignatureassertion

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/msassertion"
	"github.com/matthewmueller/golly/internal/gendom/dom/msfidosignature"
	"github.com/matthewmueller/golly/js"
)

type MSFIDOSignatureAssertion struct {
	*msassertion.MSAssertion
}

func (*MSFIDOSignatureAssertion) GetSignature() (signature *msfidosignature.MSFIDOSignature) {
	js.Rewrite("$<.signature")
	return signature
}
