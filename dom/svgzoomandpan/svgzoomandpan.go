package svgzoomandpan

import "github.com/matthewmueller/joy/macro"

// SVGZoomAndPan struct
// js:"SVGZoomAndPan,omit"
type SVGZoomAndPan struct {
}

// ZoomAndPan prop
// js:"zoomAndPan"
func (*SVGZoomAndPan) ZoomAndPan() (zoomAndPan uint8) {
	macro.Rewrite("$_.zoomAndPan")
	return zoomAndPan
}
