package cryptokeypair

import (
	"github.com/matthewmueller/joy/dom/cryptokey"
	"github.com/matthewmueller/joy/macro"
)

// CryptoKeyPair struct
// js:"CryptoKeyPair,omit"
type CryptoKeyPair struct {
}

// PrivateKey prop
// js:"privateKey"
func (*CryptoKeyPair) PrivateKey() (privateKey *cryptokey.CryptoKey) {
	macro.Rewrite("$_.privateKey")
	return privateKey
}

// SetPrivateKey prop
// js:"privateKey"
func (*CryptoKeyPair) SetPrivateKey(privateKey *cryptokey.CryptoKey) {
	macro.Rewrite("$_.privateKey = $1", privateKey)
}

// PublicKey prop
// js:"publicKey"
func (*CryptoKeyPair) PublicKey() (publicKey *cryptokey.CryptoKey) {
	macro.Rewrite("$_.publicKey")
	return publicKey
}

// SetPublicKey prop
// js:"publicKey"
func (*CryptoKeyPair) SetPublicKey(publicKey *cryptokey.CryptoKey) {
	macro.Rewrite("$_.publicKey = $1", publicKey)
}
