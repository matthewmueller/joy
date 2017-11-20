package mediakeysystemconfiguration

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/mediakeysrequirement"
	"github.com/matthewmueller/golly/internal/gendom/dom/mediakeysystemmediacapability"
)

type MediaKeySystemConfiguration struct {
	audioCapabilities     *[]*mediakeysystemmediacapability.MediaKeySystemMediaCapability
	distinctiveIdentifier *mediakeysrequirement.MediaKeysRequirement
	initDataTypes         *[]string
	persistentState       *mediakeysrequirement.MediaKeysRequirement
	videoCapabilities     *[]*mediakeysystemmediacapability.MediaKeySystemMediaCapability
}
