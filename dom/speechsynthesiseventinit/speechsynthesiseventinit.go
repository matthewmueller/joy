package speechsynthesiseventinit

import (
	"github.com/matthewmueller/joy/dom/eventinit"
	"github.com/matthewmueller/joy/dom/window"
)

type SpeechSynthesisEventInit struct {
	*eventinit.EventInit

	charIndex   *uint
	elapsedTime *float32
	name        *string
	utterance   *window.SpeechSynthesisUtterance
}
