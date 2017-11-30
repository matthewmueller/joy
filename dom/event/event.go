package event

// Event interface
// js:"Event"
type Event interface {

	// PreventDefault
	// js:"preventDefault"
	PreventDefault()

	// StopImmediatePropagation
	// js:"stopImmediatePropagation"
	StopImmediatePropagation()

	// StopPropagation
	// js:"stopPropagation"
	StopPropagation()

	// Type prop
	// js:"type"
	Type() (kind string)
}
