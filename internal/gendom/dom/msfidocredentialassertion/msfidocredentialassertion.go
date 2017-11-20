package msfidocredentialassertion

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/msassertion"
	"github.com/matthewmueller/golly/internal/gendom/dom/mstransporttype"
	"github.com/matthewmueller/golly/js"
)

type MSFIDOCredentialAssertion struct {
	*msassertion.MSAssertion
}

func (*MSFIDOCredentialAssertion) GetAlgorithm() (algorithm interface{}) {
	js.Rewrite("$<.algorithm")
	return algorithm
}

func (*MSFIDOCredentialAssertion) GetAttestation() (attestation bool) {
	js.Rewrite("$<.attestation")
	return attestation
}

func (*MSFIDOCredentialAssertion) GetPublicKey() (publicKey string) {
	js.Rewrite("$<.publicKey")
	return publicKey
}

func (*MSFIDOCredentialAssertion) GetTransportHints() (transportHints []*mstransporttype.MSTransportType) {
	js.Rewrite("$<.transportHints")
	return transportHints
}
