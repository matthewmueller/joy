package svgpathsegcurvetoquadraticsmoothrel

import (
	"github.com/matthewmueller/joy/dom/svgpathseg"
	"github.com/matthewmueller/joy/macro"
)

var _ svgpathseg.SVGPathSeg = (*SVGPathSegCurvetoQuadraticSmoothRel)(nil)

// SVGPathSegCurvetoQuadraticSmoothRel struct
// js:"SVGPathSegCurvetoQuadraticSmoothRel,omit"
type SVGPathSegCurvetoQuadraticSmoothRel struct {
}

// X prop
// js:"x"
func (*SVGPathSegCurvetoQuadraticSmoothRel) X() (x float32) {
	macro.Rewrite("$_.x")
	return x
}

// SetX prop
// js:"x"
func (*SVGPathSegCurvetoQuadraticSmoothRel) SetX(x float32) {
	macro.Rewrite("$_.x = $1", x)
}

// Y prop
// js:"y"
func (*SVGPathSegCurvetoQuadraticSmoothRel) Y() (y float32) {
	macro.Rewrite("$_.y")
	return y
}

// SetY prop
// js:"y"
func (*SVGPathSegCurvetoQuadraticSmoothRel) SetY(y float32) {
	macro.Rewrite("$_.y = $1", y)
}

// PathSegType prop
// js:"pathSegType"
func (*SVGPathSegCurvetoQuadraticSmoothRel) PathSegType() (pathSegType uint8) {
	macro.Rewrite("$_.pathSegType")
	return pathSegType
}

// PathSegTypeAsLetter prop
// js:"pathSegTypeAsLetter"
func (*SVGPathSegCurvetoQuadraticSmoothRel) PathSegTypeAsLetter() (pathSegTypeAsLetter string) {
	macro.Rewrite("$_.pathSegTypeAsLetter")
	return pathSegTypeAsLetter
}
