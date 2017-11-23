package cryptokey

import (
	"github.com/matthewmueller/golly/dom2/keyalgorithm"
	"github.com/matthewmueller/golly/js"
)

// CryptoKey struct
// js:"CryptoKey,omit"
type CryptoKey struct {
}

// Algorithm prop
func (*CryptoKey) Algorithm() (algorithm *keyalgorithm.KeyAlgorithm) {
	js.Rewrite("$<.algorithm")
	return algorithm
}

// Extractable prop
func (*CryptoKey) Extractable() (extractable bool) {
	js.Rewrite("$<.extractable")
	return extractable
}

// Type prop
func (*CryptoKey) Type() (kind string) {
	js.Rewrite("$<.type")
	return kind
}

// Usages prop
func (*CryptoKey) Usages() (usages []string) {
	js.Rewrite("$<.usages")
	return usages
}
