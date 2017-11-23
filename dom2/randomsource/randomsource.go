package randomsource

import "github.com/matthewmueller/golly/js"

// RandomSource struct
// js:"RandomSource,omit"
type RandomSource struct {
}

// GetRandomValues
func (*RandomSource) GetRandomValues(array []byte) (b []byte) {
	js.Rewrite("$<.GetRandomValues($1)", array)
	return b
}
