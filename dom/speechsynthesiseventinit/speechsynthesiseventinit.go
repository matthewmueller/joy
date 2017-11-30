package speechsynthesiseventinit

import (
	"github.com/matthewmueller/golly/dom/eventinit"
	"github.com/matthewmueller/golly/dom/window"
)

type SpeechSynthesisEventInit struct {
	*eventinit.EventInit

	charIndex   *uint
	elapsedTime *float32
	name        *string
	utterance   *window.SpeechSynthesisUtterance
}
