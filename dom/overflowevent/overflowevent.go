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

// HorizontalOverflow prop
func (*OverflowEvent) HorizontalOverflow() (horizontalOverflow bool) {
	js.Rewrite("$<.horizontalOverflow")
	return horizontalOverflow
}

// Orient prop
func (*OverflowEvent) Orient() (orient uint) {
	js.Rewrite("$<.orient")
	return orient
}

// VerticalOverflow prop
func (*OverflowEvent) VerticalOverflow() (verticalOverflow bool) {
	js.Rewrite("$<.verticalOverflow")
	return verticalOverflow
}
