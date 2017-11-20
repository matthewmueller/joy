package msmediakeysession

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/eventtarget"
	"github.com/matthewmueller/golly/internal/gendom/dom/msmediakeyerror"
	"github.com/matthewmueller/golly/js"
)

type MSMediaKeySession struct {
	*eventtarget.EventTarget
}

func (*MSMediaKeySession) Close() {
	js.Rewrite("$<.close()")
}

func (*MSMediaKeySession) Update(key []uint8) {
	js.Rewrite("$<.update($1)", key)
}

func (*MSMediaKeySession) GetError() (err *msmediakeyerror.MSMediaKeyError) {
	js.Rewrite("$<.error")
	return err
}

func (*MSMediaKeySession) GetKeySystem() (keySystem string) {
	js.Rewrite("$<.keySystem")
	return keySystem
}

func (*MSMediaKeySession) GetSessionID() (sessionId string) {
	js.Rewrite("$<.sessionId")
	return sessionId
}
