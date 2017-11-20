package msmediakeys

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/msmediakeysession"
	"github.com/matthewmueller/golly/js"
)

type MSMediaKeys struct {
}

func (*MSMediaKeys) CreateSession(kind string, initData []uint8, cdmData []uint8) (m *msmediakeysession.MSMediaKeySession) {
	js.Rewrite("$<.createSession($1, $2, $3)", kind, initData, cdmData)
	return m
}

func (*MSMediaKeys) IsTypeSupported(keySystem string, kind string) (b bool) {
	js.Rewrite("$<.isTypeSupported($1, $2)", keySystem, kind)
	return b
}

func (*MSMediaKeys) IsTypeSupportedWithFeatures(keySystem string, kind string) (s string) {
	js.Rewrite("$<.isTypeSupportedWithFeatures($1, $2)", keySystem, kind)
	return s
}

func (*MSMediaKeys) GetKeySystem() (keySystem string) {
	js.Rewrite("$<.keySystem")
	return keySystem
}
