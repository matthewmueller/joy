package randomsource

// RandomSource interface
// js:"RandomSource"
type RandomSource interface {

	// GetRandomValues
	// js:"getRandomValues"
	GetRandomValues(array []byte) (b []byte)
}
