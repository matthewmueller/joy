package window

import (
	"github.com/matthewmueller/joy/dom/mediakeysystemconfiguration"
	"github.com/matthewmueller/joy/macro"
)

// MediaKeySystemAccess struct
// js:"MediaKeySystemAccess,omit"
type MediaKeySystemAccess struct {
}

// CreateMediaKeys fn
// js:"createMediaKeys"
func (*MediaKeySystemAccess) CreateMediaKeys() (m *MediaKeys) {
	macro.Rewrite("await $_.createMediaKeys()")
	return m
}

// GetConfiguration fn
// js:"getConfiguration"
func (*MediaKeySystemAccess) GetConfiguration() (m *mediakeysystemconfiguration.MediaKeySystemConfiguration) {
	macro.Rewrite("$_.getConfiguration()")
	return m
}

// KeySystem prop
// js:"keySystem"
func (*MediaKeySystemAccess) KeySystem() (keySystem string) {
	macro.Rewrite("$_.keySystem")
	return keySystem
}
