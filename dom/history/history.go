package history

import "github.com/matthewmueller/joy/macro"

// History struct
// js:"History,omit"
type History struct {
}

// Back fn
// js:"back"
func (*History) Back(distance *interface{}) {
	macro.Rewrite("$_.back($1)", distance)
}

// Forward fn
// js:"forward"
func (*History) Forward(distance *interface{}) {
	macro.Rewrite("$_.forward($1)", distance)
}

// Go fn
// js:"go"
func (*History) Go(delta *interface{}) {
	macro.Rewrite("$_.go($1)", delta)
}

// PushState fn
// js:"pushState"
func (*History) PushState(statedata interface{}, title *string, url *string) {
	macro.Rewrite("$_.pushState($1, $2, $3)", statedata, title, url)
}

// ReplaceState fn
// js:"replaceState"
func (*History) ReplaceState(statedata interface{}, title *string, url *string) {
	macro.Rewrite("$_.replaceState($1, $2, $3)", statedata, title, url)
}

// Length prop
// js:"length"
func (*History) Length() (length int) {
	macro.Rewrite("$_.length")
	return length
}

// State prop
// js:"state"
func (*History) State() (state interface{}) {
	macro.Rewrite("$_.state")
	return state
}
