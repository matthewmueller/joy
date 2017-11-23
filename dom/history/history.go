package history

import "github.com/matthewmueller/golly/js"

// History struct
// js:"History,omit"
type History struct {
}

// Back fn
func (*History) Back(distance *interface{}) {
	js.Rewrite("$<.back($1)", distance)
}

// Forward fn
func (*History) Forward(distance *interface{}) {
	js.Rewrite("$<.forward($1)", distance)
}

// Go fn
func (*History) Go(delta *interface{}) {
	js.Rewrite("$<.go($1)", delta)
}

// PushState fn
func (*History) PushState(statedata interface{}, title *string, url *string) {
	js.Rewrite("$<.pushState($1, $2, $3)", statedata, title, url)
}

// ReplaceState fn
func (*History) ReplaceState(statedata interface{}, title *string, url *string) {
	js.Rewrite("$<.replaceState($1, $2, $3)", statedata, title, url)
}

// Length prop
func (*History) Length() (length int) {
	js.Rewrite("$<.length")
	return length
}

// State prop
func (*History) State() (state interface{}) {
	js.Rewrite("$<.state")
	return state
}
