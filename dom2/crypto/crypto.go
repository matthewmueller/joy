package crypto

import (
	"github.com/matthewmueller/golly/dom2/randomsource"
	"github.com/matthewmueller/golly/dom2/subtlecrypto"
	"github.com/matthewmueller/golly/js"
)

// js:"Crypto,omit"
type Crypto struct {
	randomsource.RandomSource
}

// Subtle
func (*Crypto) Subtle() (subtle *subtlecrypto.SubtleCrypto) {
	js.Rewrite("$<.Subtle")
	return subtle
}
