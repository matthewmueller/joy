package window

import (
	"github.com/matthewmueller/joy/dom/mediakeysessiontype"
	"github.com/matthewmueller/joy/macro"
)

// MediaKeys struct
// js:"MediaKeys,omit"
type MediaKeys struct {
}

// CreateSession fn
// js:"createSession"
func (*MediaKeys) CreateSession(sessionType *mediakeysessiontype.MediaKeySessionType) (m *MediaKeySession) {
	macro.Rewrite("$_.createSession($1)", sessionType)
	return m
}

// SetServerCertificate fn
// js:"setServerCertificate"
func (*MediaKeys) SetServerCertificate(serverCertificate []byte) {
	macro.Rewrite("await $_.setServerCertificate($1)", serverCertificate)
}
