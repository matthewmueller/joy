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

// Algorithm
func (*MSFIDOCredentialAssertion) Algorithm() (algorithm interface{}) {
	js.Rewrite("$<.Algorithm")
	return algorithm
}

// Attestation
func (*MSFIDOCredentialAssertion) Attestation() (attestation bool) {
	js.Rewrite("$<.Attestation")
	return attestation
}

// PublicKey
func (*MSFIDOCredentialAssertion) PublicKey() (publicKey string) {
	js.Rewrite("$<.PublicKey")
	return publicKey
}

// TransportHints
func (*MSFIDOCredentialAssertion) TransportHints() (transportHints []*mstransporttype.MSTransportType) {
	js.Rewrite("$<.TransportHints")
	return transportHints
}
