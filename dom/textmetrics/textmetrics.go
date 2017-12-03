package textmetrics

import "github.com/matthewmueller/joy/macro"

// TextMetrics struct
// js:"TextMetrics,omit"
type TextMetrics struct {
}

// Width prop
// js:"width"
func (*TextMetrics) Width() (width float32) {
	macro.Rewrite("$_.width")
	return width
}
