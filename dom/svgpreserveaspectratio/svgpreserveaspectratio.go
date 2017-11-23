package svgpreserveaspectratio

import "github.com/matthewmueller/golly/js"

// SVGPreserveAspectRatio struct
// js:"SVGPreserveAspectRatio,omit"
type SVGPreserveAspectRatio struct {
}

// Align prop
func (*SVGPreserveAspectRatio) Align() (align uint8) {
	js.Rewrite("$<.align")
	return align
}

// Align prop
func (*SVGPreserveAspectRatio) SetAlign(align uint8) {
	js.Rewrite("$<.align = $1", align)
}

// MeetOrSlice prop
func (*SVGPreserveAspectRatio) MeetOrSlice() (meetOrSlice uint8) {
	js.Rewrite("$<.meetOrSlice")
	return meetOrSlice
}

// MeetOrSlice prop
func (*SVGPreserveAspectRatio) SetMeetOrSlice(meetOrSlice uint8) {
	js.Rewrite("$<.meetOrSlice = $1", meetOrSlice)
}
