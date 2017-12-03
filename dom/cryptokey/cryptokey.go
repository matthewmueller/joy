package cryptokey

import (
	"github.com/matthewmueller/joy/dom/keyalgorithm"
	"github.com/matthewmueller/joy/js"
)

// CryptoKey struct
// js:"CryptoKey,omit"
type CryptoKey struct {
}

// Algorithm prop
// js:"algorithm"
func (*CryptoKey) Algorithm() (algorithm *keyalgorithm.KeyAlgorithm) {
	js.Rewrite("$_.algorithm")
	return algorithm
}

// Extractable prop
// js:"extractable"
func (*CryptoKey) Extractable() (extractable bool) {
	js.Rewrite("$_.extractable")
	return extractable
}

// Type prop
// js:"type"
func (*CryptoKey) Type() (kind string) {
	js.Rewrite("$_.type")
	return kind
}

// Usages prop
// js:"usages"
func (*CryptoKey) Usages() (usages []string) {
	js.Rewrite("$_.usages")
	return usages
}
