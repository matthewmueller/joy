package msconnectivity

import (
	"github.com/matthewmueller/joy/dom/msicetype"
	"github.com/matthewmueller/joy/dom/msicewarningflags"
	"github.com/matthewmueller/joy/dom/msrelayaddress"
)

type MSConnectivity struct {
	iceType         *msicetype.MSIceType
	iceWarningFlags *msicewarningflags.MSIceWarningFlags
	relayAddress    *msrelayaddress.MSRelayAddress
}
