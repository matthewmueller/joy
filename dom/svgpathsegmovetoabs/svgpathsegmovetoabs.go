package svgpathsegmovetoabs

import (
	"github.com/matthewmueller/joy/dom/svgpathseg"
	"github.com/matthewmueller/joy/js"
)

var _ svgpathseg.SVGPathSeg = (*SVGPathSegMovetoAbs)(nil)

// SVGPathSegMovetoAbs struct
// js:"SVGPathSegMovetoAbs,omit"
type SVGPathSegMovetoAbs struct {
}

// X prop
// js:"x"
func (*SVGPathSegMovetoAbs) X() (x float32) {
	js.Rewrite("$_.x")
	return x
}

// SetX prop
// js:"x"
func (*SVGPathSegMovetoAbs) SetX(x float32) {
	js.Rewrite("$_.x = $1", x)
}

// Y prop
// js:"y"
func (*SVGPathSegMovetoAbs) Y() (y float32) {
	js.Rewrite("$_.y")
	return y
}

// SetY prop
// js:"y"
func (*SVGPathSegMovetoAbs) SetY(y float32) {
	js.Rewrite("$_.y = $1", y)
}

// PathSegType prop
// js:"pathSegType"
func (*SVGPathSegMovetoAbs) PathSegType() (pathSegType uint8) {
	js.Rewrite("$_.pathSegType")
	return pathSegType
}

// PathSegTypeAsLetter prop
// js:"pathSegTypeAsLetter"
func (*SVGPathSegMovetoAbs) PathSegTypeAsLetter() (pathSegTypeAsLetter string) {
	js.Rewrite("$_.pathSegTypeAsLetter")
	return pathSegTypeAsLetter
}
