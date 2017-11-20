package mediaencryptedevent

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/event"
	"github.com/matthewmueller/golly/js"
)

type MediaEncryptedEvent struct {
	*event.Event
}

func (*MediaEncryptedEvent) GetInitData() (initData []byte) {
	js.Rewrite("$<.initData")
	return initData
}

func (*MediaEncryptedEvent) GetInitDataType() (initDataType string) {
	js.Rewrite("$<.initDataType")
	return initDataType
}
