package webauthnassertion

import (
	"github.com/matthewmueller/joy/dom/scopedcredential"
	"github.com/matthewmueller/joy/macro"
)

// WebAuthnAssertion struct
// js:"WebAuthnAssertion,omit"
type WebAuthnAssertion struct {
}

// AuthenticatorData prop
// js:"authenticatorData"
func (*WebAuthnAssertion) AuthenticatorData() (authenticatorData []byte) {
	macro.Rewrite("$_.authenticatorData")
	return authenticatorData
}

// ClientData prop
// js:"clientData"
func (*WebAuthnAssertion) ClientData() (clientData []byte) {
	macro.Rewrite("$_.clientData")
	return clientData
}

// Credential prop
// js:"credential"
func (*WebAuthnAssertion) Credential() (credential *scopedcredential.ScopedCredential) {
	macro.Rewrite("$_.credential")
	return credential
}

// Signature prop
// js:"signature"
func (*WebAuthnAssertion) Signature() (signature []byte) {
	macro.Rewrite("$_.signature")
	return signature
}
