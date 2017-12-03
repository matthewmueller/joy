package barprop

import "github.com/matthewmueller/joy/macro"

// BarProp struct
// js:"BarProp,omit"
type BarProp struct {
}

// Visible prop
// js:"visible"
func (*BarProp) Visible() (visible bool) {
	macro.Rewrite("$_.visible")
	return visible
}
