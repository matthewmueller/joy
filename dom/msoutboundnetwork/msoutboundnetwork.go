package msoutboundnetwork

import "github.com/matthewmueller/joy/dom/msnetwork"

type MSOutboundNetwork struct {
	*msnetwork.MSNetwork

	appliedBandwidthLimit *uint
}
