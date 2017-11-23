package textmetrics

import "github.com/matthewmueller/golly/js"

// js:"TextMetrics,omit"
type TextMetrics struct {
}

// Width
func (*TextMetrics) Width() (width float32) {
	js.Rewrite("$<.Width")
	return width
}
