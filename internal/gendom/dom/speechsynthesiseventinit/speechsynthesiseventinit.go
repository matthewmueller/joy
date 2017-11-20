package speechsynthesiseventinit

import "github.com/matthewmueller/golly/internal/gendom/dom/speechsynthesisutterance"

type SpeechSynthesisEventInit struct {
	*EventInit

	charIndex   *uint
	elapsedTime *float32
	name        *string
	utterance   *speechsynthesisutterance.SpeechSynthesisUtterance
}
