package cryptokeypair

import (
	"github.com/matthewmueller/golly/dom2/cryptokey"
	"github.com/matthewmueller/golly/js"
)

// js:"CryptoKeyPair,omit"
type CryptoKeyPair struct {
}

// PrivateKey
func (*CryptoKeyPair) PrivateKey() (privateKey *cryptokey.CryptoKey) {
	js.Rewrite("$<.PrivateKey")
	return privateKey
}

// PrivateKey
func (*CryptoKeyPair) SetPrivateKey(privateKey *cryptokey.CryptoKey) {
	js.Rewrite("$<.PrivateKey = $1", privateKey)
}

// PublicKey
func (*CryptoKeyPair) PublicKey() (publicKey *cryptokey.CryptoKey) {
	js.Rewrite("$<.PublicKey")
	return publicKey
}

// PublicKey
func (*CryptoKeyPair) SetPublicKey(publicKey *cryptokey.CryptoKey) {
	js.Rewrite("$<.PublicKey = $1", publicKey)
}
