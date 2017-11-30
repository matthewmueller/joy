package canvasgradient

import "github.com/matthewmueller/golly/js"

// CanvasGradient struct
// js:"CanvasGradient,omit"
type CanvasGradient struct {
}

// AddColorStop fn
// js:"addColorStop"
func (*CanvasGradient) AddColorStop(offset float32, color string) {
	js.Rewrite("$_.addColorStop($1, $2)", offset, color)
}
