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
func (*CryptoKeyPair) PrivateKey() (privateKey *cryptokey.CryptoKey) {
	js.Rewrite("$<.privateKey")
	return privateKey
}

// PrivateKey prop
func (*CryptoKeyPair) SetPrivateKey(privateKey *cryptokey.CryptoKey) {
	js.Rewrite("$<.privateKey = $1", privateKey)
}

// PublicKey prop
func (*CryptoKeyPair) PublicKey() (publicKey *cryptokey.CryptoKey) {
	js.Rewrite("$<.publicKey")
	return publicKey
}

// PublicKey prop
func (*CryptoKeyPair) SetPublicKey(publicKey *cryptokey.CryptoKey) {
	js.Rewrite("$<.publicKey = $1", publicKey)
}
