package barprop

import "github.com/matthewmueller/golly/js"

// BarProp struct
// js:"BarProp,omit"
type BarProp struct {
}

// Visible
func (*BarProp) Visible() (visible bool) {
	js.Rewrite("$<.Visible")
	return visible
}
