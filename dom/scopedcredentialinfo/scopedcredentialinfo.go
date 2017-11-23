package scopedcredentialinfo

import (
	"github.com/matthewmueller/golly/dom2/cryptokey"
	"github.com/matthewmueller/golly/dom2/scopedcredential"
	"github.com/matthewmueller/golly/js"
)

// ScopedCredentialInfo struct
// js:"ScopedCredentialInfo,omit"
type ScopedCredentialInfo struct {
}

// Credential prop
func (*ScopedCredentialInfo) Credential() (credential *scopedcredential.ScopedCredential) {
	js.Rewrite("$<.credential")
	return credential
}

// PublicKey prop
func (*ScopedCredentialInfo) PublicKey() (publicKey *cryptokey.CryptoKey) {
	js.Rewrite("$<.publicKey")
	return publicKey
}
