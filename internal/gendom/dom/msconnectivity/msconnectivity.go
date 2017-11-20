package msconnectivity

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/msicetype"
	"github.com/matthewmueller/golly/internal/gendom/dom/msicewarningflags"
	"github.com/matthewmueller/golly/internal/gendom/dom/msrelayaddress"
)

type MSConnectivity struct {
	iceType         *msicetype.MSIceType
	iceWarningFlags *msicewarningflags.MSIceWarningFlags
	relayAddress    *msrelayaddress.MSRelayAddress
}
