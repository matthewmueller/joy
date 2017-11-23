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
func (*MediaKeys) CreateSession(sessionType *mediakeysessiontype.MediaKeySessionType) (m *MediaKeySession) {
	js.Rewrite("$<.createSession($1)", sessionType)
	return m
}

// SetServerCertificate fn
func (*MediaKeys) SetServerCertificate(serverCertificate []byte) {
	js.Rewrite("await $<.setServerCertificate($1)", serverCertificate)
}
