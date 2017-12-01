package msgesture

import (
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// New fn
func New() *MSGesture {
	js.Rewrite("MSGesture")
	return &MSGesture{}
}

// MSGesture struct
// js:"MSGesture,omit"
type MSGesture struct {
}

// AddPointer fn
// js:"addPointer"
func (*MSGesture) AddPointer(pointerId int) {
	js.Rewrite("$_.addPointer($1)", pointerId)
}

// Stop fn
// js:"stop"
func (*MSGesture) Stop() {
	js.Rewrite("$_.stop()")
}

// Target prop
// js:"target"
func (*MSGesture) Target() (target window.Element) {
	js.Rewrite("$_.target")
	return target
}

// SetTarget prop
// js:"target"
func (*MSGesture) SetTarget(target window.Element) {
	js.Rewrite("$_.target = $1", target)
}
