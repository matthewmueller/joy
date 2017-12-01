package svgpathsegcurvetocubicsmoothrel

import (
	"github.com/matthewmueller/golly/dom/svgpathseg"
	"github.com/matthewmueller/golly/js"
)

var _ svgpathseg.SVGPathSeg = (*SVGPathSegCurvetoCubicSmoothRel)(nil)

// SVGPathSegCurvetoCubicSmoothRel struct
// js:"SVGPathSegCurvetoCubicSmoothRel,omit"
type SVGPathSegCurvetoCubicSmoothRel struct {
}

// X prop
// js:"x"
func (*SVGPathSegCurvetoCubicSmoothRel) X() (x float32) {
	js.Rewrite("$_.x")
	return x
}

// SetX prop
// js:"x"
func (*SVGPathSegCurvetoCubicSmoothRel) SetX(x float32) {
	js.Rewrite("$_.x = $1", x)
}

// X2 prop
// js:"x2"
func (*SVGPathSegCurvetoCubicSmoothRel) X2() (x2 float32) {
	js.Rewrite("$_.x2")
	return x2
}

// SetX2 prop
// js:"x2"
func (*SVGPathSegCurvetoCubicSmoothRel) SetX2(x2 float32) {
	js.Rewrite("$_.x2 = $1", x2)
}

// Y prop
// js:"y"
func (*SVGPathSegCurvetoCubicSmoothRel) Y() (y float32) {
	js.Rewrite("$_.y")
	return y
}

// SetY prop
// js:"y"
func (*SVGPathSegCurvetoCubicSmoothRel) SetY(y float32) {
	js.Rewrite("$_.y = $1", y)
}

// Y2 prop
// js:"y2"
func (*SVGPathSegCurvetoCubicSmoothRel) Y2() (y2 float32) {
	js.Rewrite("$_.y2")
	return y2
}

// SetY2 prop
// js:"y2"
func (*SVGPathSegCurvetoCubicSmoothRel) SetY2(y2 float32) {
	js.Rewrite("$_.y2 = $1", y2)
}

// PathSegType prop
// js:"pathSegType"
func (*SVGPathSegCurvetoCubicSmoothRel) PathSegType() (pathSegType uint8) {
	js.Rewrite("$_.pathSegType")
	return pathSegType
}

// PathSegTypeAsLetter prop
// js:"pathSegTypeAsLetter"
func (*SVGPathSegCurvetoCubicSmoothRel) PathSegTypeAsLetter() (pathSegTypeAsLetter string) {
	js.Rewrite("$_.pathSegTypeAsLetter")
	return pathSegTypeAsLetter
}
