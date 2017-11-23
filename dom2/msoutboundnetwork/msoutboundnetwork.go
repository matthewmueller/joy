package msoutboundnetwork

type MSOutboundNetwork struct {
	*msnetwork.MSNetwork

	appliedBandwidthLimit *uint
}
