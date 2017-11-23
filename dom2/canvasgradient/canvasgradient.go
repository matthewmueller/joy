package canvasgradient

import "github.com/matthewmueller/golly/js"

// js:"CanvasGradient,omit"
type CanvasGradient struct {
}

// AddColorStop
func (*CanvasGradient) AddColorStop(offset float32, color string) {
	js.Rewrite("$<.AddColorStop($1, $2)", offset, color)
}
