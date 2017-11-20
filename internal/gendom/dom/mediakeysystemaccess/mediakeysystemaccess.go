package mediakeysystemaccess

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/mediakeys"
	"github.com/matthewmueller/golly/internal/gendom/dom/mediakeysystemconfiguration"
	"github.com/matthewmueller/golly/js"
)

type MediaKeySystemAccess struct {
}

func (*MediaKeySystemAccess) CreateMediaKeys() (m *mediakeys.MediaKeys) {
	js.Rewrite("await $<.createMediaKeys()")
	return m
}

func (*MediaKeySystemAccess) GetConfiguration() (m *mediakeysystemconfiguration.MediaKeySystemConfiguration) {
	js.Rewrite("$<.getConfiguration()")
	return m
}

func (*MediaKeySystemAccess) GetKeySystem() (keySystem string) {
	js.Rewrite("$<.keySystem")
	return keySystem
}
