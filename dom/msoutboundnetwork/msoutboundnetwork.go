package msoutboundnetwork

import "github.com/matthewmueller/golly/dom/msnetwork"

type MSOutboundNetwork struct {
	*msnetwork.MSNetwork

	appliedBandwidthLimit *uint
}
