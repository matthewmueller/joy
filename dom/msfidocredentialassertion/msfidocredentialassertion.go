package msfidocredentialassertion

import (
	"github.com/matthewmueller/golly/dom2/msassertion"
	"github.com/matthewmueller/golly/dom2/mstransporttype"
	"github.com/matthewmueller/golly/js"
)

// MSFIDOCredentialAssertion struct
// js:"MSFIDOCredentialAssertion,omit"
type MSFIDOCredentialAssertion struct {
	msassertion.MSAssertion
}

// Algorithm prop
func (*MSFIDOCredentialAssertion) Algorithm() (algorithm interface{}) {
	js.Rewrite("$<.algorithm")
	return algorithm
}

// Attestation prop
func (*MSFIDOCredentialAssertion) Attestation() (attestation bool) {
	js.Rewrite("$<.attestation")
	return attestation
}

// PublicKey prop
func (*MSFIDOCredentialAssertion) PublicKey() (publicKey string) {
	js.Rewrite("$<.publicKey")
	return publicKey
}

// TransportHints prop
func (*MSFIDOCredentialAssertion) TransportHints() (transportHints []*mstransporttype.MSTransportType) {
	js.Rewrite("$<.transportHints")
	return transportHints
}
