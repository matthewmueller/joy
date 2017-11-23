package randomsource

// js:"RandomSource,omit"
type RandomSource interface {

	// GetRandomValues
	GetRandomValues(array []byte) (b []byte)
}
