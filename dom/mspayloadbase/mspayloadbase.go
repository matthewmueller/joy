package mspayloadbase

import "github.com/matthewmueller/golly/dom2/rtcstats"

type MSPayloadBase struct {
	*rtcstats.RTCStats

	payloadDescription *string
}
