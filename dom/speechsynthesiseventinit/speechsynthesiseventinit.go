package speechsynthesiseventinit

import (
	"github.com/matthewmueller/golly/dom2/eventinit"
	"github.com/matthewmueller/golly/dom2/window"
)

type SpeechSynthesisEventInit struct {
	*eventinit.EventInit

	charIndex   *uint
	elapsedTime *float32
	name        *string
	utterance   *window.SpeechSynthesisUtterance
}
