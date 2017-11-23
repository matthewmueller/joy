package mediakeymessageeventinit

import (
	"github.com/matthewmueller/golly/dom2/eventinit"
	"github.com/matthewmueller/golly/dom2/mediakeymessagetype"
)

type MediaKeyMessageEventInit struct {
	*eventinit.EventInit

	message     *[]byte
	messageType *mediakeymessagetype.MediaKeyMessageType
}
