package webauthnassertion

import (
	"github.com/matthewmueller/joy/dom/scopedcredential"
	"github.com/matthewmueller/joy/js"
)

// WebAuthnAssertion struct
// js:"WebAuthnAssertion,omit"
type WebAuthnAssertion struct {
}

// AuthenticatorData prop
// js:"authenticatorData"
func (*WebAuthnAssertion) AuthenticatorData() (authenticatorData []byte) {
	js.Rewrite("$_.authenticatorData")
	return authenticatorData
}

// ClientData prop
// js:"clientData"
func (*WebAuthnAssertion) ClientData() (clientData []byte) {
	js.Rewrite("$_.clientData")
	return clientData
}

// Credential prop
// js:"credential"
func (*WebAuthnAssertion) Credential() (credential *scopedcredential.ScopedCredential) {
	js.Rewrite("$_.credential")
	return credential
}

// Signature prop
// js:"signature"
func (*WebAuthnAssertion) Signature() (signature []byte) {
	js.Rewrite("$_.signature")
	return signature
}
