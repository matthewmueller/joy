package msmediakeys

import (
	"github.com/matthewmueller/golly/dom2/msmediakeysession"
	"github.com/matthewmueller/golly/js"
)

// js:"MSMediaKeys,omit"
type MSMediaKeys struct {
}

// CreateSession
func (*MSMediaKeys) CreateSession(kind string, initData []uint8, cdmData *[]uint8) (m *msmediakeysession.MSMediaKeySession) {
	js.Rewrite("$<.CreateSession($1, $2, $3)", kind, initData, cdmData)
	return m
}

// IsTypeSupported
func (*MSMediaKeys) IsTypeSupported(keySystem string, kind *string) (b bool) {
	js.Rewrite("$<.IsTypeSupported($1, $2)", keySystem, kind)
	return b
}

// IsTypeSupportedWithFeatures
func (*MSMediaKeys) IsTypeSupportedWithFeatures(keySystem string, kind *string) (s string) {
	js.Rewrite("$<.IsTypeSupportedWithFeatures($1, $2)", keySystem, kind)
	return s
}

// KeySystem
func (*MSMediaKeys) KeySystem() (keySystem string) {
	js.Rewrite("$<.KeySystem")
	return keySystem
}
