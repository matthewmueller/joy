package crypto

import (
	"github.com/matthewmueller/joy/dom/subtlecrypto"
	"github.com/matthewmueller/joy/macro"
)

// Crypto struct
// js:"Crypto,omit"
type Crypto struct {
}

// GetRandomValues fn
// js:"getRandomValues"
func (*Crypto) GetRandomValues(array []byte) (b []byte) {
	macro.Rewrite("$_.getRandomValues($1)", array)
	return b
}

// Subtle prop
// js:"subtle"
func (*Crypto) Subtle() (subtle *subtlecrypto.SubtleCrypto) {
	macro.Rewrite("$_.subtle")
	return subtle
}
