package window

import (
	"github.com/matthewmueller/golly/dom/mediakeysystemconfiguration"
	"github.com/matthewmueller/golly/js"
)

// MediaKeySystemAccess struct
// js:"MediaKeySystemAccess,omit"
type MediaKeySystemAccess struct {
}

// CreateMediaKeys fn
func (*MediaKeySystemAccess) CreateMediaKeys() (m *MediaKeys) {
	js.Rewrite("await $<.createMediaKeys()")
	return m
}

// GetConfiguration fn
func (*MediaKeySystemAccess) GetConfiguration() (m *mediakeysystemconfiguration.MediaKeySystemConfiguration) {
	js.Rewrite("$<.getConfiguration()")
	return m
}

// KeySystem prop
func (*MediaKeySystemAccess) KeySystem() (keySystem string) {
	js.Rewrite("$<.keySystem")
	return keySystem
}
