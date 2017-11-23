package webauthnassertion

import "github.com/matthewmueller/golly/js"

// js:"WebAuthnAssertion,omit"
type WebAuthnAssertion struct {
}

// AuthenticatorData
func (*WebAuthnAssertion) AuthenticatorData() (authenticatorData []byte) {
	js.Rewrite("$<.AuthenticatorData")
	return authenticatorData
}

// ClientData
func (*WebAuthnAssertion) ClientData() (clientData []byte) {
	js.Rewrite("$<.ClientData")
	return clientData
}

// Credential
func (*WebAuthnAssertion) Credential() (credential *scopedcredential.ScopedCredential) {
	js.Rewrite("$<.Credential")
	return credential
}

// Signature
func (*WebAuthnAssertion) Signature() (signature []byte) {
	js.Rewrite("$<.Signature")
	return signature
}
