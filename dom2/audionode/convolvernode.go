package audionode

import (
	"github.com/matthewmueller/golly/dom2/audiobuffer"
	"github.com/matthewmueller/golly/js"
)

// js:"ConvolverNode,omit"
type ConvolverNode struct {
	AudioNode
}

// Buffer
func (*ConvolverNode) Buffer() (buffer *audiobuffer.AudioBuffer) {
	js.Rewrite("$<.Buffer")
	return buffer
}

// Buffer
func (*ConvolverNode) SetBuffer(buffer *audiobuffer.AudioBuffer) {
	js.Rewrite("$<.Buffer = $1", buffer)
}

// Normalize
func (*ConvolverNode) Normalize() (normalize bool) {
	js.Rewrite("$<.Normalize")
	return normalize
}

// Normalize
func (*ConvolverNode) SetNormalize(normalize bool) {
	js.Rewrite("$<.Normalize = $1", normalize)
}
