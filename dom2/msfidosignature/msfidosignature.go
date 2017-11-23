package msfidosignature

import "github.com/matthewmueller/golly/js"

// js:"MSFIDOSignature,omit"
type MSFIDOSignature struct {
}

// AuthnrData
func (*MSFIDOSignature) AuthnrData() (authnrData string) {
	js.Rewrite("$<.AuthnrData")
	return authnrData
}

// ClientData
func (*MSFIDOSignature) ClientData() (clientData string) {
	js.Rewrite("$<.ClientData")
	return clientData
}

// Signature
func (*MSFIDOSignature) Signature() (signature string) {
	js.Rewrite("$<.Signature")
	return signature
}
