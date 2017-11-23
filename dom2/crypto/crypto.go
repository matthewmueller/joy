package crypto

import (
	"github.com/matthewmueller/golly/dom2/subtlecrypto"
	"github.com/matthewmueller/golly/js"
)

// Crypto struct
// js:"Crypto,omit"
type Crypto struct {
}

// GetRandomValues
func (*Crypto) GetRandomValues(array []byte) (b []byte) {
	js.Rewrite("$<.GetRandomValues($1)", array)
	return b
}

// Subtle
func (*Crypto) Subtle() (subtle *subtlecrypto.SubtleCrypto) {
	js.Rewrite("$<.Subtle")
	return subtle
}
