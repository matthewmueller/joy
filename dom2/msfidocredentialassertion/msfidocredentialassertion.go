package msfidocredentialassertion

import "github.com/matthewmueller/golly/js"

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
