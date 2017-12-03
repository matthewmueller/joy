package cryptokey

import (
	"github.com/matthewmueller/joy/dom/keyalgorithm"
	"github.com/matthewmueller/joy/macro"
)

// CryptoKey struct
// js:"CryptoKey,omit"
type CryptoKey struct {
}

// Algorithm prop
// js:"algorithm"
func (*CryptoKey) Algorithm() (algorithm *keyalgorithm.KeyAlgorithm) {
	macro.Rewrite("$_.algorithm")
	return algorithm
}

// Extractable prop
// js:"extractable"
func (*CryptoKey) Extractable() (extractable bool) {
	macro.Rewrite("$_.extractable")
	return extractable
}

// Type prop
// js:"type"
func (*CryptoKey) Type() (kind string) {
	macro.Rewrite("$_.type")
	return kind
}

// Usages prop
// js:"usages"
func (*CryptoKey) Usages() (usages []string) {
	macro.Rewrite("$_.usages")
	return usages
}
