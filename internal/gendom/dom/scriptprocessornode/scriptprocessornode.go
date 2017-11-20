package scriptprocessornode

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/audionode"
	"github.com/matthewmueller/golly/internal/gendom/dom/audioprocessingevent"
	"github.com/matthewmueller/golly/js"
)

type ScriptProcessorNode struct {
	*audionode.AudioNode
}

func (*ScriptProcessorNode) GetBufferSize() (bufferSize int) {
	js.Rewrite("$<.bufferSize")
	return bufferSize
}

func (*ScriptProcessorNode) GetOnaudioprocess() (audioprocess *audioprocessingevent.AudioProcessingEvent) {
	js.Rewrite("$<.onaudioprocess")
	return audioprocess
}

func (*ScriptProcessorNode) SetOnaudioprocess(audioprocess *audioprocessingevent.AudioProcessingEvent) {
	js.Rewrite("$<.onaudioprocess = $1", audioprocess)
}
