package msmediakeys

import (
	"github.com/matthewmueller/golly/dom/msmediakeysession"
	"github.com/matthewmueller/golly/js"
)

// MSMediaKeys struct
// js:"MSMediaKeys,omit"
type MSMediaKeys struct {
}

// CreateSession fn
// js:"createSession"
func (*MSMediaKeys) CreateSession(kind string, initData []uint8, cdmData *[]uint8) (m *msmediakeysession.MSMediaKeySession) {
	js.Rewrite("$_.createSession($1, $2, $3)", kind, initData, cdmData)
	return m
}

// IsTypeSupported fn
// js:"isTypeSupported"
func (*MSMediaKeys) IsTypeSupported(keySystem string, kind *string) (b bool) {
	js.Rewrite("$_.isTypeSupported($1, $2)", keySystem, kind)
	return b
}

// IsTypeSupportedWithFeatures fn
// js:"isTypeSupportedWithFeatures"
func (*MSMediaKeys) IsTypeSupportedWithFeatures(keySystem string, kind *string) (s string) {
	js.Rewrite("$_.isTypeSupportedWithFeatures($1, $2)", keySystem, kind)
	return s
}

// KeySystem prop
// js:"keySystem"
func (*MSMediaKeys) KeySystem() (keySystem string) {
	js.Rewrite("$_.keySystem")
	return keySystem
}
