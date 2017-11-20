package textmetrics

import "github.com/matthewmueller/golly/js"

type TextMetrics struct {
}

func (*TextMetrics) GetWidth() (width float32) {
	js.Rewrite("$<.width")
	return width
}
