package msgesture

import (
	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/macro"
)

// New fn
func New() *MSGesture {
	macro.Rewrite("new MSGesture()")
	return &MSGesture{}
}

// MSGesture struct
// js:"MSGesture,omit"
type MSGesture struct {
}

// AddPointer fn
// js:"addPointer"
func (*MSGesture) AddPointer(pointerId int) {
	macro.Rewrite("$_.addPointer($1)", pointerId)
}

// Stop fn
// js:"stop"
func (*MSGesture) Stop() {
	macro.Rewrite("$_.stop()")
}

// Target prop
// js:"target"
func (*MSGesture) Target() (target window.Element) {
	macro.Rewrite("$_.target")
	return target
}

// SetTarget prop
// js:"target"
func (*MSGesture) SetTarget(target window.Element) {
	macro.Rewrite("$_.target = $1", target)
}
