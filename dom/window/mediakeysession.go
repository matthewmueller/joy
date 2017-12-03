package window

import (
	"github.com/matthewmueller/joy/dom/mediakeystatusmap"
	"github.com/matthewmueller/joy/macro"
)

var _ EventTarget = (*MediaKeySession)(nil)

// MediaKeySession struct
// js:"MediaKeySession,omit"
type MediaKeySession struct {
}

// Close fn
// js:"close"
func (*MediaKeySession) Close() {
	macro.Rewrite("await $_.close()")
}

// GenerateRequest fn
// js:"generateRequest"
func (*MediaKeySession) GenerateRequest(initDataType string, initData []byte) {
	macro.Rewrite("await $_.generateRequest($1, $2)", initDataType, initData)
}

// Load fn
// js:"load"
func (*MediaKeySession) Load(sessionId string) (b bool) {
	macro.Rewrite("await $_.load($1)", sessionId)
	return b
}

// Remove fn
// js:"remove"
func (*MediaKeySession) Remove() {
	macro.Rewrite("await $_.remove()")
}

// Update fn
// js:"update"
func (*MediaKeySession) Update(response []byte) {
	macro.Rewrite("await $_.update($1)", response)
}

// AddEventListener fn
// js:"addEventListener"
func (*MediaKeySession) AddEventListener(kind string, listener func(evt Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

// DispatchEvent fn
// js:"dispatchEvent"
func (*MediaKeySession) DispatchEvent(evt Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

// RemoveEventListener fn
// js:"removeEventListener"
func (*MediaKeySession) RemoveEventListener(kind string, listener func(evt Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

// Closed prop
// js:"closed"
func (*MediaKeySession) Closed() {
	macro.Rewrite("await $_.closed")
}

// Expiration prop
// js:"expiration"
func (*MediaKeySession) Expiration() (expiration float32) {
	macro.Rewrite("$_.expiration")
	return expiration
}

// KeyStatuses prop
// js:"keyStatuses"
func (*MediaKeySession) KeyStatuses() (keyStatuses *mediakeystatusmap.MediaKeyStatusMap) {
	macro.Rewrite("$_.keyStatuses")
	return keyStatuses
}

// SessionID prop
// js:"sessionId"
func (*MediaKeySession) SessionID() (sessionId string) {
	macro.Rewrite("$_.sessionId")
	return sessionId
}
