package svgzoomandpan

import "github.com/matthewmueller/joy/js"

// SVGZoomAndPan struct
// js:"SVGZoomAndPan,omit"
type SVGZoomAndPan struct {
}

// ZoomAndPan prop
// js:"zoomAndPan"
func (*SVGZoomAndPan) ZoomAndPan() (zoomAndPan uint8) {
	js.Rewrite("$_.zoomAndPan")
	return zoomAndPan
}
