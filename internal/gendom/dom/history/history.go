package history

import "github.com/matthewmueller/golly/js"

type History struct {
}

func (*History) Back(distance interface{}) {
	js.Rewrite("$<.back($1)", distance)
}

func (*History) Forward(distance interface{}) {
	js.Rewrite("$<.forward($1)", distance)
}

func (*History) Go(delta interface{}) {
	js.Rewrite("$<.go($1)", delta)
}

func (*History) PushState(statedata interface{}, title string, url string) {
	js.Rewrite("$<.pushState($1, $2, $3)", statedata, title, url)
}

func (*History) ReplaceState(statedata interface{}, title string, url string) {
	js.Rewrite("$<.replaceState($1, $2, $3)", statedata, title, url)
}

func (*History) GetLength() (length int) {
	js.Rewrite("$<.length")
	return length
}

func (*History) GetState() (state interface{}) {
	js.Rewrite("$<.state")
	return state
}
