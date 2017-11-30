package mouseevent

import (
	"github.com/matthewmueller/golly/dom/event"
	"github.com/matthewmueller/golly/dom/uievent"
	"github.com/matthewmueller/golly/js"
)

var _ uievent.UIEvent = (*MouseEvent)(nil)
var _ event.Event = (*MouseEvent)(nil)

// MouseEvent struct
// js:"MouseEvent,omit"
type MouseEvent struct {
}

// PreventDefault fn
// js:"preventDefault"
func (*MouseEvent) PreventDefault() {
	js.Rewrite("$_.preventDefault()")
}

// StopImmediatePropagation fn
// js:"stopImmediatePropagation"
func (*MouseEvent) StopImmediatePropagation() {
	js.Rewrite("$_.stopImmediatePropagation()")
}

// StopPropagation fn
// js:"stopPropagation"
func (*MouseEvent) StopPropagation() {
	js.Rewrite("$_.stopPropagation()")
}

// X prop
// js:"x"
func (*MouseEvent) X() (x int) {
	js.Rewrite("$_.x")
	return x
}

// Y prop
// js:"y"
func (*MouseEvent) Y() (y int) {
	js.Rewrite("$_.y")
	return y
}

// Type prop
// js:"type"
func (*MouseEvent) Type() (kind string) {
	js.Rewrite("$_.type")
	return kind
}
