package audionode

import (
	"github.com/matthewmueller/golly/dom2/audioprocessingevent"
	"github.com/matthewmueller/golly/js"
)

// ScriptProcessorNode struct
// js:"ScriptProcessorNode,omit"
type ScriptProcessorNode struct {
	AudioNode
}

// BufferSize
func (*ScriptProcessorNode) BufferSize() (bufferSize int) {
	js.Rewrite("$<.BufferSize")
	return bufferSize
}

// Onaudioprocess
func (*ScriptProcessorNode) Onaudioprocess() (onaudioprocess func(*audioprocessingevent.AudioProcessingEvent)) {
	js.Rewrite("$<.Onaudioprocess")
	return onaudioprocess
}

// Onaudioprocess
func (*ScriptProcessorNode) SetOnaudioprocess(onaudioprocess func(*audioprocessingevent.AudioProcessingEvent)) {
	js.Rewrite("$<.Onaudioprocess = $1", onaudioprocess)
}
