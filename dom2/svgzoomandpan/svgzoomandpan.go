package svgzoomandpan

import "github.com/matthewmueller/golly/js"

// SVGZoomAndPan struct
// js:"SVGZoomAndPan,omit"
type SVGZoomAndPan struct {
}

// ZoomAndPan
func (*SVGZoomAndPan) ZoomAndPan() (zoomAndPan uint8) {
	js.Rewrite("$<.ZoomAndPan")
	return zoomAndPan
}
