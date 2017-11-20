package svgpreserveaspectratio

import "github.com/matthewmueller/golly/js"

type SVGPreserveAspectRatio struct {
}

func (*SVGPreserveAspectRatio) GetAlign() (align uint8) {
	js.Rewrite("$<.align")
	return align
}

func (*SVGPreserveAspectRatio) SetAlign(align uint8) {
	js.Rewrite("$<.align = $1", align)
}

func (*SVGPreserveAspectRatio) GetMeetOrSlice() (meetOrSlice uint8) {
	js.Rewrite("$<.meetOrSlice")
	return meetOrSlice
}

func (*SVGPreserveAspectRatio) SetMeetOrSlice(meetOrSlice uint8) {
	js.Rewrite("$<.meetOrSlice = $1", meetOrSlice)
}
