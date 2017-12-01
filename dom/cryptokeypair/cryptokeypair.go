package cryptokeypair

import (
	"github.com/matthewmueller/golly/dom/cryptokey"
	"github.com/matthewmueller/golly/js"
)

// CryptoKeyPair struct
// js:"CryptoKeyPair,omit"
type CryptoKeyPair struct {
}

// PrivateKey prop
// js:"privateKey"
func (*CryptoKeyPair) PrivateKey() (privateKey *cryptokey.CryptoKey) {
	js.Rewrite("$_.privateKey")
	return privateKey
}

// SetPrivateKey prop
// js:"privateKey"
func (*CryptoKeyPair) SetPrivateKey(privateKey *cryptokey.CryptoKey) {
	js.Rewrite("$_.privateKey = $1", privateKey)
}

// PublicKey prop
// js:"publicKey"
func (*CryptoKeyPair) PublicKey() (publicKey *cryptokey.CryptoKey) {
	js.Rewrite("$_.publicKey")
	return publicKey
}

// SetPublicKey prop
// js:"publicKey"
func (*CryptoKeyPair) SetPublicKey(publicKey *cryptokey.CryptoKey) {
	js.Rewrite("$_.publicKey = $1", publicKey)
}
