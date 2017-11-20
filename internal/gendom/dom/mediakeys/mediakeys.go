package mediakeys

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/mediakeysession"
	"github.com/matthewmueller/golly/internal/gendom/dom/mediakeysessiontype"
	"github.com/matthewmueller/golly/js"
)

type MediaKeys struct {
}

func (*MediaKeys) CreateSession(sessionType *mediakeysessiontype.MediaKeySessionType) (m *mediakeysession.MediaKeySession) {
	js.Rewrite("$<.createSession($1)", sessionType)
	return m
}

func (*MediaKeys) SetServerCertificate(serverCertificate []byte) {
	js.Rewrite("await $<.setServerCertificate($1)", serverCertificate)
}
