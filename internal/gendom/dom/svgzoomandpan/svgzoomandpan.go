package svgzoomandpan

import "github.com/matthewmueller/golly/js"

type SVGZoomAndPan struct {
}

func (*SVGZoomAndPan) GetZoomAndPan() (zoomAndPan uint8) {
	js.Rewrite("$<.zoomAndPan")
	return zoomAndPan
}
