package svgpreserveaspectratio

import "github.com/matthewmueller/golly/js"

// SVGPreserveAspectRatio struct
// js:"SVGPreserveAspectRatio,omit"
type SVGPreserveAspectRatio struct {
}

// Align
func (*SVGPreserveAspectRatio) Align() (align uint8) {
	js.Rewrite("$<.Align")
	return align
}

// Align
func (*SVGPreserveAspectRatio) SetAlign(align uint8) {
	js.Rewrite("$<.Align = $1", align)
}

// MeetOrSlice
func (*SVGPreserveAspectRatio) MeetOrSlice() (meetOrSlice uint8) {
	js.Rewrite("$<.MeetOrSlice")
	return meetOrSlice
}

// MeetOrSlice
func (*SVGPreserveAspectRatio) SetMeetOrSlice(meetOrSlice uint8) {
	js.Rewrite("$<.MeetOrSlice = $1", meetOrSlice)
}
