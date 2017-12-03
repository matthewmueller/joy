package rtcdtmftonechangeeventinit

import "github.com/matthewmueller/joy/dom/eventinit"

type RTCDTMFToneChangeEventInit struct {
	*eventinit.EventInit

	tone *string
}
