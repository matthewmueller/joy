package window

// AbstractWorker interface
// js:"AbstractWorker"
type AbstractWorker interface {

	// Onerror prop
	// js:"onerror"
	// jsrewrite:"$_.onerror"
	Onerror() (onerror func(Event))

	// SetOnerror prop
	// js:"onerror"
	// jsrewrite:"$_.onerror = $1"
	SetOnerror(onerror func(Event))
}
