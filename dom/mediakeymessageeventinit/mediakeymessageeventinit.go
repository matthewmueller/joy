package mediakeymessageeventinit

import (
	"github.com/matthewmueller/golly/dom/eventinit"
	"github.com/matthewmueller/golly/dom/mediakeymessagetype"
)

type MediaKeyMessageEventInit struct {
	*eventinit.EventInit

	message     *[]byte
	messageType *mediakeymessagetype.MediaKeyMessageType
}
