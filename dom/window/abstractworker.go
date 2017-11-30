package window

// AbstractWorker interface
// js:"AbstractWorker"
type AbstractWorker interface {

	// Onerror prop
	// js:"onerror"
	Onerror() (onerror func(Event))

	// SetOnerror prop
	// js:"onerror"
	SetOnerror(onerror func(Event))
}
