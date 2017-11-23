package msmediakeys

import (
	"github.com/matthewmueller/golly/dom2/msmediakeysession"
	"github.com/matthewmueller/golly/js"
)

// New fn
func New(keysystem string) *MSMediaKeys {
	js.Rewrite("MSMediaKeys")
	return &MSMediaKeys{}
}

// MSMediaKeys struct
// js:"MSMediaKeys,omit"
type MSMediaKeys struct {
}

// CreateSession fn
func (*MSMediaKeys) CreateSession(kind string, initData []uint8, cdmData *[]uint8) (m *msmediakeysession.MSMediaKeySession) {
	js.Rewrite("$<.createSession($1, $2, $3)", kind, initData, cdmData)
	return m
}

// IsTypeSupported fn
func (*MSMediaKeys) IsTypeSupported(keySystem string, kind *string) (b bool) {
	js.Rewrite("$<.isTypeSupported($1, $2)", keySystem, kind)
	return b
}

// IsTypeSupportedWithFeatures fn
func (*MSMediaKeys) IsTypeSupportedWithFeatures(keySystem string, kind *string) (s string) {
	js.Rewrite("$<.isTypeSupportedWithFeatures($1, $2)", keySystem, kind)
	return s
}

// KeySystem prop
func (*MSMediaKeys) KeySystem() (keySystem string) {
	js.Rewrite("$<.keySystem")
	return keySystem
}
