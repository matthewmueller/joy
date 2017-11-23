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

// BufferSize prop
func (*ScriptProcessorNode) BufferSize() (bufferSize int) {
	js.Rewrite("$<.bufferSize")
	return bufferSize
}

// Onaudioprocess prop
func (*ScriptProcessorNode) Onaudioprocess() (onaudioprocess func(*audioprocessingevent.AudioProcessingEvent)) {
	js.Rewrite("$<.onaudioprocess")
	return onaudioprocess
}

// Onaudioprocess prop
func (*ScriptProcessorNode) SetOnaudioprocess(onaudioprocess func(*audioprocessingevent.AudioProcessingEvent)) {
	js.Rewrite("$<.onaudioprocess = $1", onaudioprocess)
}
