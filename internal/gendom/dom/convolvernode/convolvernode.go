package convolvernode

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/audiobuffer"
	"github.com/matthewmueller/golly/internal/gendom/dom/audionode"
	"github.com/matthewmueller/golly/js"
)

type ConvolverNode struct {
	*audionode.AudioNode
}

func (*ConvolverNode) GetBuffer() (buffer *audiobuffer.AudioBuffer) {
	js.Rewrite("$<.buffer")
	return buffer
}

func (*ConvolverNode) SetBuffer(buffer *audiobuffer.AudioBuffer) {
	js.Rewrite("$<.buffer = $1", buffer)
}

func (*ConvolverNode) GetNormalize() (normalize bool) {
	js.Rewrite("$<.normalize")
	return normalize
}

func (*ConvolverNode) SetNormalize(normalize bool) {
	js.Rewrite("$<.normalize = $1", normalize)
}
