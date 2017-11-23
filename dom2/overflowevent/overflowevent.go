package overflowevent

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// OverflowEvent struct
// js:"OverflowEvent,omit"
type OverflowEvent struct {
	window.UIEvent
}

// HorizontalOverflow
func (*OverflowEvent) HorizontalOverflow() (horizontalOverflow bool) {
	js.Rewrite("$<.HorizontalOverflow")
	return horizontalOverflow
}

// Orient
func (*OverflowEvent) Orient() (orient uint) {
	js.Rewrite("$<.Orient")
	return orient
}

// VerticalOverflow
func (*OverflowEvent) VerticalOverflow() (verticalOverflow bool) {
	js.Rewrite("$<.VerticalOverflow")
	return verticalOverflow
}
