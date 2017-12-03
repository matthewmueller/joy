package history

import "github.com/matthewmueller/joy/js"

// History struct
// js:"History,omit"
type History struct {
}

// Back fn
// js:"back"
func (*History) Back(distance *interface{}) {
	js.Rewrite("$_.back($1)", distance)
}

// Forward fn
// js:"forward"
func (*History) Forward(distance *interface{}) {
	js.Rewrite("$_.forward($1)", distance)
}

// Go fn
// js:"go"
func (*History) Go(delta *interface{}) {
	js.Rewrite("$_.go($1)", delta)
}

// PushState fn
// js:"pushState"
func (*History) PushState(statedata interface{}, title *string, url *string) {
	js.Rewrite("$_.pushState($1, $2, $3)", statedata, title, url)
}

// ReplaceState fn
// js:"replaceState"
func (*History) ReplaceState(statedata interface{}, title *string, url *string) {
	js.Rewrite("$_.replaceState($1, $2, $3)", statedata, title, url)
}

// Length prop
// js:"length"
func (*History) Length() (length int) {
	js.Rewrite("$_.length")
	return length
}

// State prop
// js:"state"
func (*History) State() (state interface{}) {
	js.Rewrite("$_.state")
	return state
}
