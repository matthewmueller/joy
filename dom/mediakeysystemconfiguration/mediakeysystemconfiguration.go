package mediakeysystemconfiguration

import (
	"github.com/matthewmueller/golly/dom2/mediakeysrequirement"
	"github.com/matthewmueller/golly/dom2/mediakeysystemmediacapability"
)

type MediaKeySystemConfiguration struct {
	audioCapabilities     *[]*mediakeysystemmediacapability.MediaKeySystemMediaCapability
	distinctiveIdentifier *mediakeysrequirement.MediaKeysRequirement
	initDataTypes         *[]string
	persistentState       *mediakeysrequirement.MediaKeysRequirement
	videoCapabilities     *[]*mediakeysystemmediacapability.MediaKeySystemMediaCapability
}
