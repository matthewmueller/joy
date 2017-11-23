package msconnectivity

import (
	"github.com/matthewmueller/golly/dom/msicetype"
	"github.com/matthewmueller/golly/dom/msicewarningflags"
	"github.com/matthewmueller/golly/dom/msrelayaddress"
)

type MSConnectivity struct {
	iceType         *msicetype.MSIceType
	iceWarningFlags *msicewarningflags.MSIceWarningFlags
	relayAddress    *msrelayaddress.MSRelayAddress
}
