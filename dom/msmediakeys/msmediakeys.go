package msmediakeys

import (
	"github.com/matthewmueller/joy/dom/msmediakeysession"
	"github.com/matthewmueller/joy/macro"
)

// New fn
func New(keysystem string) *MSMediaKeys {
	macro.Rewrite("MSMediaKeys")
	return &MSMediaKeys{}
}

// MSMediaKeys struct
// js:"MSMediaKeys,omit"
type MSMediaKeys struct {
}

// CreateSession fn
// js:"createSession"
func (*MSMediaKeys) CreateSession(kind string, initData []uint8, cdmData *[]uint8) (m *msmediakeysession.MSMediaKeySession) {
	macro.Rewrite("$_.createSession($1, $2, $3)", kind, initData, cdmData)
	return m
}

// IsTypeSupported fn
// js:"isTypeSupported"
func (*MSMediaKeys) IsTypeSupported(keySystem string, kind *string) (b bool) {
	macro.Rewrite("$_.isTypeSupported($1, $2)", keySystem, kind)
	return b
}

// IsTypeSupportedWithFeatures fn
// js:"isTypeSupportedWithFeatures"
func (*MSMediaKeys) IsTypeSupportedWithFeatures(keySystem string, kind *string) (s string) {
	macro.Rewrite("$_.isTypeSupportedWithFeatures($1, $2)", keySystem, kind)
	return s
}

// KeySystem prop
// js:"keySystem"
func (*MSMediaKeys) KeySystem() (keySystem string) {
	macro.Rewrite("$_.keySystem")
	return keySystem
}
