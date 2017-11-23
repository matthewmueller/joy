package msconnectivity

import (
	"github.com/matthewmueller/golly/dom2/msicetype"
	"github.com/matthewmueller/golly/dom2/msicewarningflags"
	"github.com/matthewmueller/golly/dom2/msrelayaddress"
)

type MSConnectivity struct {
	iceType         *msicetype.MSIceType
	iceWarningFlags *msicewarningflags.MSIceWarningFlags
	relayAddress    *msrelayaddress.MSRelayAddress
}
