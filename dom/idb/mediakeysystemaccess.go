package idb

import (
	"github.com/matthewmueller/golly/dom2/mediakeysystemconfiguration"
	"github.com/matthewmueller/golly/js"
)

// MediaKeySystemAccess struct
// js:"MediaKeySystemAccess,omit"
type MediaKeySystemAccess struct {
}

// CreateMediaKeys
func (*MediaKeySystemAccess) CreateMediaKeys() (m *MediaKeys) {
	js.Rewrite("await $<.CreateMediaKeys()")
	return m
}

// GetConfiguration
func (*MediaKeySystemAccess) GetConfiguration() (m *mediakeysystemconfiguration.MediaKeySystemConfiguration) {
	js.Rewrite("$<.GetConfiguration()")
	return m
}

// KeySystem
func (*MediaKeySystemAccess) KeySystem() (keySystem string) {
	js.Rewrite("$<.KeySystem")
	return keySystem
}
