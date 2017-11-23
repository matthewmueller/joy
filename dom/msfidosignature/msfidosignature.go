package msfidosignature

import "github.com/matthewmueller/golly/js"

// MSFIDOSignature struct
// js:"MSFIDOSignature,omit"
type MSFIDOSignature struct {
}

// AuthnrData prop
func (*MSFIDOSignature) AuthnrData() (authnrData string) {
	js.Rewrite("$<.authnrData")
	return authnrData
}

// ClientData prop
func (*MSFIDOSignature) ClientData() (clientData string) {
	js.Rewrite("$<.clientData")
	return clientData
}

// Signature prop
func (*MSFIDOSignature) Signature() (signature string) {
	js.Rewrite("$<.signature")
	return signature
}
