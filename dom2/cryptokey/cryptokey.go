package cryptokey

import "github.com/matthewmueller/golly/js"

// js:"CryptoKey,omit"
type CryptoKey struct {
}

// Algorithm
func (*CryptoKey) Algorithm() (algorithm *keyalgorithm.KeyAlgorithm) {
	js.Rewrite("$<.Algorithm")
	return algorithm
}

// Extractable
func (*CryptoKey) Extractable() (extractable bool) {
	js.Rewrite("$<.Extractable")
	return extractable
}

// Type
func (*CryptoKey) Type() (kind string) {
	js.Rewrite("$<.Type")
	return kind
}

// Usages
func (*CryptoKey) Usages() (usages []string) {
	js.Rewrite("$<.Usages")
	return usages
}
