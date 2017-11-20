package msfidosignature

import "github.com/matthewmueller/golly/js"

type MSFIDOSignature struct {
}

func (*MSFIDOSignature) GetAuthnrData() (authnrData string) {
	js.Rewrite("$<.authnrData")
	return authnrData
}

func (*MSFIDOSignature) GetClientData() (clientData string) {
	js.Rewrite("$<.clientData")
	return clientData
}

func (*MSFIDOSignature) GetSignature() (signature string) {
	js.Rewrite("$<.signature")
	return signature
}
