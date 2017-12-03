package crypto

import (
	"github.com/matthewmueller/joy/dom/subtlecrypto"
	"github.com/matthewmueller/joy/js"
)

// Crypto struct
// js:"Crypto,omit"
type Crypto struct {
}

// GetRandomValues fn
// js:"getRandomValues"
func (*Crypto) GetRandomValues(array []byte) (b []byte) {
	js.Rewrite("$_.getRandomValues($1)", array)
	return b
}

// Subtle prop
// js:"subtle"
func (*Crypto) Subtle() (subtle *subtlecrypto.SubtleCrypto) {
	js.Rewrite("$_.subtle")
	return subtle
}
