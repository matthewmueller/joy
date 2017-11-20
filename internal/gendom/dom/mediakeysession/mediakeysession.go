package mediakeysession

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/eventtarget"
	"github.com/matthewmueller/golly/internal/gendom/dom/mediakeystatusmap"
	"github.com/matthewmueller/golly/js"
)

type MediaKeySession struct {
	*eventtarget.EventTarget
}

func (*MediaKeySession) Close() {
	js.Rewrite("await $<.close()")
}

func (*MediaKeySession) GenerateRequest(initDataType string, initData []byte) {
	js.Rewrite("await $<.generateRequest($1, $2)", initDataType, initData)
}

func (*MediaKeySession) Load(sessionId string) (b bool) {
	js.Rewrite("await $<.load($1)", sessionId)
	return b
}

func (*MediaKeySession) Remove() {
	js.Rewrite("await $<.remove()")
}

func (*MediaKeySession) Update(response []byte) {
	js.Rewrite("await $<.update($1)", response)
}

func (*MediaKeySession) GetClosed() {
	js.Rewrite("await $<.closed")
}

func (*MediaKeySession) GetExpiration() (expiration float32) {
	js.Rewrite("$<.expiration")
	return expiration
}

func (*MediaKeySession) GetKeyStatuses() (keyStatuses *mediakeystatusmap.MediaKeyStatusMap) {
	js.Rewrite("$<.keyStatuses")
	return keyStatuses
}

func (*MediaKeySession) GetSessionID() (sessionId string) {
	js.Rewrite("$<.sessionId")
	return sessionId
}
