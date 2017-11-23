package window

// js:"AbstractWorker,omit"
type AbstractWorker interface {

	// Onerror
	Onerror() (onerror func(Event))

	// Onerror
	SetOnerror(onerror func(Event))
}
