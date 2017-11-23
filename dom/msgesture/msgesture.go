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
func (*MSGesture) AddPointer(pointerId int) {
	js.Rewrite("$<.addPointer($1)", pointerId)
}

// Stop fn
func (*MSGesture) Stop() {
	js.Rewrite("$<.stop()")
}

// Target prop
func (*MSGesture) Target() (target window.Element) {
	js.Rewrite("$<.target")
	return target
}

// Target prop
func (*MSGesture) SetTarget(target window.Element) {
	js.Rewrite("$<.target = $1", target)
}
