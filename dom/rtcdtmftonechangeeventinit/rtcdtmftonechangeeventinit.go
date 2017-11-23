package rtcdtmftonechangeeventinit

import "github.com/matthewmueller/golly/dom/eventinit"

type RTCDTMFToneChangeEventInit struct {
	*eventinit.EventInit

	tone *string
}
