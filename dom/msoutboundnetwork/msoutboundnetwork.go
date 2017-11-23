package msoutboundnetwork

import "github.com/matthewmueller/golly/dom2/msnetwork"

type MSOutboundNetwork struct {
	*msnetwork.MSNetwork

	appliedBandwidthLimit *uint
}
