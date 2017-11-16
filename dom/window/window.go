package window

import "github.com/matthewmueller/golly/js"

type Event struct{}

// Methods
func (e *Event) PreventDefault() {

}

// Properties
func (e *Event) Bubbles() bool {
	js.Rewrite("$<.bubbles")
	return false
}

func addEventListener(kind string, listener func(ev Event)) {

}
