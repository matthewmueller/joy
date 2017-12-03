package mediakeymessageeventinit

import (
	"github.com/matthewmueller/joy/dom/eventinit"
	"github.com/matthewmueller/joy/dom/mediakeymessagetype"
)

type MediaKeyMessageEventInit struct {
	*eventinit.EventInit

	message     *[]byte
	messageType *mediakeymessagetype.MediaKeyMessageType
}
