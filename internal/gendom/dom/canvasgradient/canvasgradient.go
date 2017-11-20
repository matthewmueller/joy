package canvasgradient

import "github.com/matthewmueller/golly/js"

type CanvasGradient struct {
}

func (*CanvasGradient) AddColorStop(offset float32, color string) {
	js.Rewrite("$<.addColorStop($1, $2)", offset, color)
}
