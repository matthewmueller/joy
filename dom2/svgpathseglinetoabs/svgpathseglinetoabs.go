package svgpathseglinetoabs

import "github.com/matthewmueller/golly/js"

// js:"SVGPathSegLinetoAbs,omit"
type SVGPathSegLinetoAbs struct {
	svgpathseg.SVGPathSeg
}

// X
func (*SVGPathSegLinetoAbs) X() (x float32) {
	js.Rewrite("$<.X")
	return x
}

// X
func (*SVGPathSegLinetoAbs) SetX(x float32) {
	js.Rewrite("$<.X = $1", x)
}

// Y
func (*SVGPathSegLinetoAbs) Y() (y float32) {
	js.Rewrite("$<.Y")
	return y
}

// Y
func (*SVGPathSegLinetoAbs) SetY(y float32) {
	js.Rewrite("$<.Y = $1", y)
}
