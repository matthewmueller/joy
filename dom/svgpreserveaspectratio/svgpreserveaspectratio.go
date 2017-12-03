package svgpreserveaspectratio

import "github.com/matthewmueller/joy/macro"

// SVGPreserveAspectRatio struct
// js:"SVGPreserveAspectRatio,omit"
type SVGPreserveAspectRatio struct {
}

// Align prop
// js:"align"
func (*SVGPreserveAspectRatio) Align() (align uint8) {
	macro.Rewrite("$_.align")
	return align
}

// SetAlign prop
// js:"align"
func (*SVGPreserveAspectRatio) SetAlign(align uint8) {
	macro.Rewrite("$_.align = $1", align)
}

// MeetOrSlice prop
// js:"meetOrSlice"
func (*SVGPreserveAspectRatio) MeetOrSlice() (meetOrSlice uint8) {
	macro.Rewrite("$_.meetOrSlice")
	return meetOrSlice
}

// SetMeetOrSlice prop
// js:"meetOrSlice"
func (*SVGPreserveAspectRatio) SetMeetOrSlice(meetOrSlice uint8) {
	macro.Rewrite("$_.meetOrSlice = $1", meetOrSlice)
}
