package mediakeymessageeventinit

import "github.com/matthewmueller/golly/internal/gendom/dom/mediakeymessagetype"

type MediaKeyMessageEventInit struct {
	*EventInit

	message     *[]byte
	messageType *mediakeymessagetype.MediaKeyMessageType
}
