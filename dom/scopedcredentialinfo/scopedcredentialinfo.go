package scopedcredentialinfo

import (
	"github.com/matthewmueller/joy/dom/cryptokey"
	"github.com/matthewmueller/joy/dom/scopedcredential"
	"github.com/matthewmueller/joy/macro"
)

// ScopedCredentialInfo struct
// js:"ScopedCredentialInfo,omit"
type ScopedCredentialInfo struct {
}

// Credential prop
// js:"credential"
func (*ScopedCredentialInfo) Credential() (credential *scopedcredential.ScopedCredential) {
	macro.Rewrite("$_.credential")
	return credential
}

// PublicKey prop
// js:"publicKey"
func (*ScopedCredentialInfo) PublicKey() (publicKey *cryptokey.CryptoKey) {
	macro.Rewrite("$_.publicKey")
	return publicKey
}
