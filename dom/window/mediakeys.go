package window

import (
	"github.com/matthewmueller/golly/dom/mediakeysessiontype"
	"github.com/matthewmueller/golly/js"
)

// MediaKeys struct
// js:"MediaKeys,omit"
type MediaKeys struct {
}

// CreateSession fn
// js:"createSession"
func (*MediaKeys) CreateSession(sessionType *mediakeysessiontype.MediaKeySessionType) (m *MediaKeySession) {
	js.Rewrite("$_.createSession($1)", sessionType)
	return m
}

// SetServerCertificate fn
// js:"setServerCertificate"
func (*MediaKeys) SetServerCertificate(serverCertificate []byte) {
	js.Rewrite("await $_.setServerCertificate($1)", serverCertificate)
}
