package mediakeysystemconfiguration

type MediaKeySystemConfiguration struct {
	audioCapabilities     *[]*mediakeysystemmediacapability.MediaKeySystemMediaCapability
	distinctiveIdentifier *mediakeysrequirement.MediaKeysRequirement
	initDataTypes         *[]string
	persistentState       *mediakeysrequirement.MediaKeysRequirement
	videoCapabilities     *[]*mediakeysystemmediacapability.MediaKeySystemMediaCapability
}
