package mspayloadbase

import "github.com/matthewmueller/golly/dom/rtcstats"

type MSPayloadBase struct {
	*rtcstats.RTCStats

	payloadDescription *string
}
