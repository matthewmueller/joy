package randomsource

// RandomSource interface
// js:"RandomSource"
type RandomSource interface {

	// GetRandomValues
	// js:"getRandomValues"
	// jsrewrite:"$_.getRandomValues($1)"
	GetRandomValues(array []byte) (b []byte)
}
