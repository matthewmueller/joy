package scopedcredentialinfo

import (
	"github.com/matthewmueller/golly/dom/cryptokey"
	"github.com/matthewmueller/golly/dom/scopedcredential"
	"github.com/matthewmueller/golly/js"
)

// ScopedCredentialInfo struct
// js:"ScopedCredentialInfo,omit"
type ScopedCredentialInfo struct {
}

// Credential prop
// js:"credential"
func (*ScopedCredentialInfo) Credential() (credential *scopedcredential.ScopedCredential) {
	js.Rewrite("$_.credential")
	return credential
}

// PublicKey prop
// js:"publicKey"
func (*ScopedCredentialInfo) PublicKey() (publicKey *cryptokey.CryptoKey) {
	js.Rewrite("$_.publicKey")
	return publicKey
}
