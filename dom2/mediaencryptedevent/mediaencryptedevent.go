package mediaencryptedevent

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"MediaEncryptedEvent,omit"
type MediaEncryptedEvent struct {
	window.Event
}

// InitData
func (*MediaEncryptedEvent) InitData() (initData []byte) {
	js.Rewrite("$<.InitData")
	return initData
}

// InitDataType
func (*MediaEncryptedEvent) InitDataType() (initDataType string) {
	js.Rewrite("$<.InitDataType")
	return initDataType
}
