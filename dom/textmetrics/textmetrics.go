package textmetrics

import "github.com/matthewmueller/joy/js"

// TextMetrics struct
// js:"TextMetrics,omit"
type TextMetrics struct {
}

// Width prop
// js:"width"
func (*TextMetrics) Width() (width float32) {
	js.Rewrite("$_.width")
	return width
}
