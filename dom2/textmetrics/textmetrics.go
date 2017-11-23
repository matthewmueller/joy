package textmetrics

import "github.com/matthewmueller/golly/js"

// TextMetrics struct
// js:"TextMetrics,omit"
type TextMetrics struct {
}

// Width
func (*TextMetrics) Width() (width float32) {
	js.Rewrite("$<.Width")
	return width
}
