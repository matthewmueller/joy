package cryptokeypair

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/cryptokey"
	"github.com/matthewmueller/golly/js"
)

type CryptoKeyPair struct {
}

func (*CryptoKeyPair) GetPrivateKey() (privateKey *cryptokey.CryptoKey) {
	js.Rewrite("$<.privateKey")
	return privateKey
}

func (*CryptoKeyPair) SetPrivateKey(privateKey *cryptokey.CryptoKey) {
	js.Rewrite("$<.privateKey = $1", privateKey)
}

func (*CryptoKeyPair) GetPublicKey() (publicKey *cryptokey.CryptoKey) {
	js.Rewrite("$<.publicKey")
	return publicKey
}

func (*CryptoKeyPair) SetPublicKey(publicKey *cryptokey.CryptoKey) {
	js.Rewrite("$<.publicKey = $1", publicKey)
}
