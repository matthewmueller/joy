package mediakeysystemconfiguration

import (
	"github.com/matthewmueller/joy/dom/mediakeysrequirement"
	"github.com/matthewmueller/joy/dom/mediakeysystemmediacapability"
)

type MediaKeySystemConfiguration struct {
	audioCapabilities     *[]*mediakeysystemmediacapability.MediaKeySystemMediaCapability
	distinctiveIdentifier *mediakeysrequirement.MediaKeysRequirement
	initDataTypes         *[]string
	persistentState       *mediakeysrequirement.MediaKeysRequirement
	videoCapabilities     *[]*mediakeysystemmediacapability.MediaKeySystemMediaCapability
}
