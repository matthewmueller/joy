package webauthnassertion

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/scopedcredential"
	"github.com/matthewmueller/golly/js"
)

type WebAuthnAssertion struct {
}

func (*WebAuthnAssertion) GetAuthenticatorData() (authenticatorData []byte) {
	js.Rewrite("$<.authenticatorData")
	return authenticatorData
}

func (*WebAuthnAssertion) GetClientData() (clientData []byte) {
	js.Rewrite("$<.clientData")
	return clientData
}

func (*WebAuthnAssertion) GetCredential() (credential *scopedcredential.ScopedCredential) {
	js.Rewrite("$<.credential")
	return credential
}

func (*WebAuthnAssertion) GetSignature() (signature []byte) {
	js.Rewrite("$<.signature")
	return signature
}
