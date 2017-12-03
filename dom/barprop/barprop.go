package barprop

import "github.com/matthewmueller/joy/js"

// BarProp struct
// js:"BarProp,omit"
type BarProp struct {
}

// Visible prop
// js:"visible"
func (*BarProp) Visible() (visible bool) {
	js.Rewrite("$_.visible")
	return visible
}
