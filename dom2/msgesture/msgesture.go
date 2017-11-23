package msgesture

import (
	"github.com/matthewmueller/golly/dom2/window"
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

// AddPointer
func (*MSGesture) AddPointer(pointerId int) {
	js.Rewrite("$<.AddPointer($1)", pointerId)
}

// Stop
func (*MSGesture) Stop() {
	js.Rewrite("$<.Stop()")
}

// Target
func (*MSGesture) Target() (target window.Element) {
	js.Rewrite("$<.Target")
	return target
}

// Target
func (*MSGesture) SetTarget(target window.Element) {
	js.Rewrite("$<.Target = $1", target)
}
