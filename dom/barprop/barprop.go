package barprop

import "github.com/matthewmueller/golly/js"

// BarProp struct
// js:"BarProp,omit"
type BarProp struct {
}

// Visible prop
func (*BarProp) Visible() (visible bool) {
	js.Rewrite("$<.visible")
	return visible
}
