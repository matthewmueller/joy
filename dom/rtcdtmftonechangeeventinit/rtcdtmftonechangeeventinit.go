package rtcdtmftonechangeeventinit

import "github.com/matthewmueller/golly/dom2/eventinit"

type RTCDTMFToneChangeEventInit struct {
	*eventinit.EventInit

	tone *string
}
