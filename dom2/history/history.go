package history

import "github.com/matthewmueller/golly/js"

// js:"History,omit"
type History struct {
}

// Back
func (*History) Back(distance *interface{}) {
	js.Rewrite("$<.Back($1)", distance)
}

// Forward
func (*History) Forward(distance *interface{}) {
	js.Rewrite("$<.Forward($1)", distance)
}

// Go
func (*History) Go(delta *interface{}) {
	js.Rewrite("$<.Go($1)", delta)
}

// PushState
func (*History) PushState(statedata interface{}, title *string, url *string) {
	js.Rewrite("$<.PushState($1, $2, $3)", statedata, title, url)
}

// ReplaceState
func (*History) ReplaceState(statedata interface{}, title *string, url *string) {
	js.Rewrite("$<.ReplaceState($1, $2, $3)", statedata, title, url)
}

// Length
func (*History) Length() (length int) {
	js.Rewrite("$<.Length")
	return length
}

// State
func (*History) State() (state interface{}) {
	js.Rewrite("$<.State")
	return state
}
