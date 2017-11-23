package webauthnassertion

import (
	"github.com/matthewmueller/golly/dom/scopedcredential"
	"github.com/matthewmueller/golly/js"
)

// WebAuthnAssertion struct
// js:"WebAuthnAssertion,omit"
type WebAuthnAssertion struct {
}

// AuthenticatorData prop
func (*WebAuthnAssertion) AuthenticatorData() (authenticatorData []byte) {
	js.Rewrite("$<.authenticatorData")
	return authenticatorData
}

// ClientData prop
func (*WebAuthnAssertion) ClientData() (clientData []byte) {
	js.Rewrite("$<.clientData")
	return clientData
}

// Credential prop
func (*WebAuthnAssertion) Credential() (credential *scopedcredential.ScopedCredential) {
	js.Rewrite("$<.credential")
	return credential
}

// Signature prop
func (*WebAuthnAssertion) Signature() (signature []byte) {
	js.Rewrite("$<.signature")
	return signature
}
