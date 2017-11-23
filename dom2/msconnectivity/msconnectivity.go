package msconnectivity

import "github.com/matthewmueller/golly/dom2/msicetype"

type MSConnectivity struct {
	iceType         *msicetype.MSIceType
	iceWarningFlags *msicewarningflags.MSIceWarningFlags
	relayAddress    *msrelayaddress.MSRelayAddress
}
