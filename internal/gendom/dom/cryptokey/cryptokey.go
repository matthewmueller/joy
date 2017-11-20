package cryptokey

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/keyalgorithm"
	"github.com/matthewmueller/golly/js"
)

type CryptoKey struct {
}

func (*CryptoKey) GetAlgorithm() (algorithm *keyalgorithm.KeyAlgorithm) {
	js.Rewrite("$<.algorithm")
	return algorithm
}

func (*CryptoKey) GetExtractable() (extractable bool) {
	js.Rewrite("$<.extractable")
	return extractable
}

func (*CryptoKey) GetType() (kind string) {
	js.Rewrite("$<.type")
	return kind
}

func (*CryptoKey) GetUsages() (usages []string) {
	js.Rewrite("$<.usages")
	return usages
}
