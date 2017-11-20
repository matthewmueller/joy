package barprop

import "github.com/matthewmueller/golly/js"

type BarProp struct {
}

func (*BarProp) GetVisible() (visible bool) {
	js.Rewrite("$<.visible")
	return visible
}
