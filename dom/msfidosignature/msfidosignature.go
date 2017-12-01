package msfidosignature

import "github.com/matthewmueller/golly/js"

// MSFIDOSignature struct
// js:"MSFIDOSignature,omit"
type MSFIDOSignature struct {
}

// AuthnrData prop
// js:"authnrData"
func (*MSFIDOSignature) AuthnrData() (authnrData string) {
	js.Rewrite("$_.authnrData")
	return authnrData
}

// ClientData prop
// js:"clientData"
func (*MSFIDOSignature) ClientData() (clientData string) {
	js.Rewrite("$_.clientData")
	return clientData
}

// Signature prop
// js:"signature"
func (*MSFIDOSignature) Signature() (signature string) {
	js.Rewrite("$_.signature")
	return signature
}
