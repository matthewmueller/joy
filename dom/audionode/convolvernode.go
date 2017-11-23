package audionode

import (
	"github.com/matthewmueller/golly/dom/audiobuffer"
	"github.com/matthewmueller/golly/js"
)

// ConvolverNode struct
// js:"ConvolverNode,omit"
type ConvolverNode struct {
	AudioNode
}

// Buffer prop
func (*ConvolverNode) Buffer() (buffer *audiobuffer.AudioBuffer) {
	js.Rewrite("$<.buffer")
	return buffer
}

// Buffer prop
func (*ConvolverNode) SetBuffer(buffer *audiobuffer.AudioBuffer) {
	js.Rewrite("$<.buffer = $1", buffer)
}

// Normalize prop
func (*ConvolverNode) Normalize() (normalize bool) {
	js.Rewrite("$<.normalize")
	return normalize
}

// Normalize prop
func (*ConvolverNode) SetNormalize(normalize bool) {
	js.Rewrite("$<.normalize = $1", normalize)
}
