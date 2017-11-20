package scopedcredentialinfo

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/cryptokey"
	"github.com/matthewmueller/golly/internal/gendom/dom/scopedcredential"
	"github.com/matthewmueller/golly/js"
)

type ScopedCredentialInfo struct {
}

func (*ScopedCredentialInfo) GetCredential() (credential *scopedcredential.ScopedCredential) {
	js.Rewrite("$<.credential")
	return credential
}

func (*ScopedCredentialInfo) GetPublicKey() (publicKey *cryptokey.CryptoKey) {
	js.Rewrite("$<.publicKey")
	return publicKey
}
