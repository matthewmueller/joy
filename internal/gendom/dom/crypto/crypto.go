package crypto

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/randomsource"
	"github.com/matthewmueller/golly/internal/gendom/dom/subtlecrypto"
	"github.com/matthewmueller/golly/js"
)

type Crypto struct {
	*randomsource.RandomSource
}

func (*Crypto) GetSubtle() (subtle *subtlecrypto.SubtleCrypto) {
	js.Rewrite("$<.subtle")
	return subtle
}
