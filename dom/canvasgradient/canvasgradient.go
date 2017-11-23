package canvasgradient

import "github.com/matthewmueller/golly/js"

// CanvasGradient struct
// js:"CanvasGradient,omit"
type CanvasGradient struct {
}

// AddColorStop fn
func (*CanvasGradient) AddColorStop(offset float32, color string) {
	js.Rewrite("$<.addColorStop($1, $2)", offset, color)
}
