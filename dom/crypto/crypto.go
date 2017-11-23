package crypto

import (
	"github.com/matthewmueller/golly/dom2/subtlecrypto"
	"github.com/matthewmueller/golly/js"
)

// Crypto struct
// js:"Crypto,omit"
type Crypto struct {
}

// GetRandomValues fn
func (*Crypto) GetRandomValues(array []byte) (b []byte) {
	js.Rewrite("$<.getRandomValues($1)", array)
	return b
}

// Subtle prop
func (*Crypto) Subtle() (subtle *subtlecrypto.SubtleCrypto) {
	js.Rewrite("$<.subtle")
	return subtle
}
